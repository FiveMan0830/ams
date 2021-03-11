package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

// func TestCreateGroupSuccess(t *testing.T) {
// 	accountManagement := account.NewLDAPManagement()

// 	group, createGroupError := accountManagement.CreateGroup(adminUser, adminPassword, groupName3, groupLeaderUsername3)

// 	assert.Equal(t, createGroupError, nil)
	
// 	assert.Equal(t, group, groupName3)
// }

func TestGetGroups(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	groupList, err := accountManagement.GetGroups(adminUser, adminPassword)

	assert.Contains(t, groupList, groupName)
	assert.Contains(t, groupList, groupName2)
	assert.Equal(t, err, nil)
}

func TestGroupNameDuplicate(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderUsername)
	duplicateError := errors.New("Duplicate Group Name")

	assert.Equal(t, group, null)
	assert.Equal(t, err, duplicateError)
}


func TestSearchGroupLeader(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	leader, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, groupName)

	assert.Equal(t, leader, groupLeaderUsername)
	assert.Equal(t, err, nil)
}

func TestGroupLeaderNotExists(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderNotExists)
	leaderError := errors.New("User does not exist")

	assert.Equal(t, group, null)
	assert.Equal(t, err, leaderError)
}

func TestAddMemberToGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	result, err = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result2, err2 := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username)
	result2, err2 = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username2)

	assert.Contains(t, result, groupLeaderUsername)
	assert.Contains(t, result, username)
	assert.Contains(t, result, username2)
	assert.Equal(t, err, nil)
	
	assert.Contains(t, result2, groupLeaderUsername2)
	assert.Contains(t, result2, username)
	assert.Contains(t, result2, username2)
	assert.Equal(t, err2, nil)
}

func TestAddMemberToNotExistsGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, groupNotExistsError)
}

func TestAddNotExistsMemberToGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, usernameNotExists)
	memberNotExistsError := errors.New("User does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, memberNotExistsError)
}

func TestAddDuplicateMemberToGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	memberDuplicateError := errors.New("User already member of the group")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, memberDuplicateError)
}

func TestGetGroupMembers(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.GetGroupMembers(adminUser, adminPassword, groupName)

	assert.Contains(t, result, groupLeaderUsername)
	assert.Contains(t, result, username)
	assert.Contains(t, result, username2)
	assert.Equal(t, err, nil)
}

func TestSearchUserMemberOf(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUserMemberOf(adminUser, adminPassword, username)

	assert.Contains(t, result, groupName)
	assert.Contains(t, result, groupName2)
	assert.Equal(t, err, nil)
}

func TestRemoveMemberFromGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
	result2, err2 := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName2, username2)

	assert.NotContains(t, result, username)
	assert.Equal(t, err, nil)

	assert.NotContains(t, result2, username2)
	assert.Equal(t, err2, nil)
}

func TestRemoveMemberFromNotExistsGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, groupNotExistsError)
}

func TestRemoveNotExistsMemberFromGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, usernameNotExists)
	userNotExistsError := errors.New("User is not a member of group")

	assert.Equal(t, result, []string([]string(nil)))
	assert.Equal(t, err, userNotExistsError)
}