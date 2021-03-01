package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const adminUser string = "admin"
const adminPassword string = "admin"

// User test data
const userIDNull string = ""
const usernameNotExists string = "usernameNotExists"

const userID string = "c61965be-8176-4419-b289-4d52617728fb"
const username string = "testUser"
const givenName string = "testUser"
const surname string = "testUser"
const userPassword string = "testUser"
const userEmail string = "test@gmail.com"

const userID2 string = "c56515be-1654-7895-1564-1d52513528cf"
const username2 string = "testUser2"
const givenName2 string = "testUser2"
const surname2 string = "testUser2"
const userPassword2 string = "testUser2"
const userEmail2 string = "test2@gmail.com"

const leaderID string = "a34531da-8563-9517-3578-3e38754896dg"
const leaderUsername string = "testLeader"
const leaderGivenName string = "testLeader"
const leaderSurname string = "testLeader"
const leaderPassword string = "testLeader"
const leaderEmail string = "testLeader@gmail.com"

const leaderID2 string = "b96875kl-6842-7539-8549-2c56482648fa"
const leaderUsername2 string = "testLeader2"
const leaderGivenName2 string = "testLeader2"
const leaderSurname2 string = "testLeader2"
const leaderPassword2 string = "testLeader2"
const leaderEmail2 string = "testLeade2r@gmail.com"

// Group test data
const groupNull = ""
const groupNameNotExists = "testGroupNotExists"
const groupName string = "testGroup"
const groupName2 string = "testGroup2"
const groupLeader string = "testLeader"
const groupLeader2 string = "testLeader2"
const groupLeaderNotExists string = "testLeaderNotExists"

// Ou test data
const ouName = "testOu"

func TestCreateUser(t *testing.T) {
	accountManagement := NewLDAPManagement()

	user := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	user2 := accountManagement.CreateUser(adminUser, adminPassword, userID2, username2, givenName2, surname2, userPassword2, userEmail2)

	leader := accountManagement.CreateUser(adminUser, adminPassword, leaderID, leaderUsername, leaderGivenName, leaderSurname, leaderPassword, leaderEmail)
	leader2 := accountManagement.CreateUser(adminUser, adminPassword, leaderID2, leaderUsername2, leaderGivenName2, leaderSurname2, leaderPassword2, leaderEmail2)

	assert.Equal(t, user, nil)
	assert.Equal(t, user2, nil)

	assert.Equal(t, leader, nil)
	assert.Equal(t, leader2, nil)
}

func TestUserDuplicate(t *testing.T) {
	accountManagement := NewLDAPManagement()

	user := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	userError := errors.New("User already exist")

	assert.Equal(t, user, userError)
}

func TestSearchUser(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, username)

	assert.Equal(t, result, userID)
	assert.Equal(t, err, nil)
}

func TestUserNotFound(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, usernameNotExists)
	searchError := errors.New("User not found")

	assert.Equal(t, result, userIDNull)
	assert.Equal(t, err, searchError)
}

func TestGetUUIDByUsername(t *testing.T) {
	accountManagement := NewLDAPManagement()

	uuid, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, username)

	assert.Equal(t, uuid, userID)
	assert.Equal(t, err, nil)
}

func TestCreateGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader)
	group2, err2 := accountManagement.CreateGroup(adminUser, adminPassword, groupName2, groupLeader2)

	assert.Equal(t, group, groupName)
	assert.Equal(t, err, nil)

	assert.Equal(t, group2, groupName2)
	assert.Equal(t, err2, nil)
}

func TestGetGroups(t *testing.T) {
	accountManagement := NewLDAPManagement()

	groupList, err := accountManagement.GetGroups(adminUser, adminPassword)

	assert.Contains(t, groupList, groupName)
	assert.Contains(t, groupList, groupName2)
	assert.Equal(t, err, nil)
}

func TestGroupNameDuplicate(t *testing.T) {
	accountManagement := NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader)
	duplicateError := errors.New("Duplicate Group Name")

	assert.Equal(t, group, groupNull)
	assert.Equal(t, err, duplicateError)
}


func TestSearchGroupLeader(t *testing.T) {
	accountManagement := NewLDAPManagement()

	leader, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, groupName)

	assert.Equal(t, leader, groupLeader)
	assert.Equal(t, err, nil)
}

func TestGroupLeaderNotExists(t *testing.T) {
	accountManagement := NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderNotExists)
	leaderError := errors.New("User does not exist")

	assert.Equal(t, group, groupNull)
	assert.Equal(t, err, leaderError)
}

func TestAddMemberToGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	result, err = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result2, err2 := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username)
	result2, err2 = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username2)

	assert.Contains(t, result, groupLeader)
	assert.Contains(t, result, username)
	assert.Contains(t, result, username2)
	assert.Equal(t, err, nil)
	
	assert.Contains(t, result2, groupLeader2)
	assert.Contains(t, result2, username)
	assert.Contains(t, result2, username2)
	assert.Equal(t, err2, nil)
}

func TestAddMemberToNotExistsGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, groupNotExistsError)
}

func TestAddNotExistsMemberToGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, usernameNotExists)
	memberNotExistsError := errors.New("User does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, memberNotExistsError)
}

func TestAddDuplicateMemberToGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	memberDuplicateError := errors.New("User already member of the group")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, memberDuplicateError)
}

func TestGetGroupMembers(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.GetGroupMembers(adminUser, adminPassword, groupName)

	assert.Contains(t, result, groupLeader)
	assert.Contains(t, result, username)
	assert.Contains(t, result, username2)
	assert.Equal(t, err, nil)
}

func TestSearchUserMemberOf(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.SearchUserMemberOf(adminUser, adminPassword, username)

	assert.Contains(t, result, groupName)
	assert.Contains(t, result, groupName2)
	assert.Equal(t, err, nil)
}

func TestRemoveMemberFromGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
	result2, err2 := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName2, username2)

	assert.NotContains(t, result, username)
	assert.Equal(t, err, nil)

	assert.NotContains(t, result2, username2)
	assert.Equal(t, err2, nil)
}

func TestRemoveMemberFromNotExistsGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, groupNotExistsError)
}

func TestRemoveNotExistsMemberFromGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, usernameNotExists)
	userNotExistsError := errors.New("User is not a member of group")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, userNotExistsError)
}

func TestCreateOU(t *testing.T) {
	accountManagement := NewLDAPManagement()

	ou := accountManagement.CreateOu(adminUser, adminPassword, ouName)

	assert.Equal(t, ou, nil)
}

func TestOUNameDuplicate(t *testing.T) {
	accountManagement := NewLDAPManagement()

	ou := accountManagement.CreateOu(adminUser, adminPassword, ouName)
	duplicateError := errors.New("This Organization Unit already exists")

	assert.Equal(t, ou, duplicateError)
}

func TestTearDown(t *testing.T) {
	accountManagement := NewLDAPManagement()

	deleteUser := accountManagement.DeleteUser(adminUser, adminPassword, username)
	deleteUser2 := accountManagement.DeleteUser(adminUser, adminPassword, username2)

	deleteLeader := accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername)
	deleteLeader2 := accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername2)

	deleteGroup := accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	deleteGroup2 := accountManagement.DeleteGroup(adminUser, adminPassword, groupName2)

	deleteOu := accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	assert.Equal(t, deleteUser, nil)
	assert.Equal(t, deleteUser2, nil)

	assert.Equal(t, deleteLeader, nil)
	assert.Equal(t, deleteLeader2, nil)

	assert.Equal(t, deleteGroup, nil)
	assert.Equal(t, deleteGroup2, nil)

	assert.Equal(t, deleteOu, nil)
}