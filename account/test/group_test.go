package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestCreateGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, createGroupError := accountManagement.CreateGroup(adminUser, adminPassword, groupName3, leaderUsername3, groupID3)
	deleteGroupErr := accountManagement.DeleteGroup(adminUser, adminPassword, groupName3)

	assert.Equal(t, nil, createGroupError)
	assert.Equal(t, nil, deleteGroupErr)
	assert.Equal(t, groupName3, group)
}

func TestGetGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	groupList, err := accountManagement.GetGroups(adminUser, adminPassword)

	assert.Contains(t,groupList, groupName)
	assert.Contains(t, groupList, groupName2)
	assert.Equal(t, nil, err)
}

func TestCreateGroupDuplicateName(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderUsername, groupID)
	duplicateError := errors.New("Duplicate Group Name")

	assert.Equal(t, null, group)
	assert.Equal(t, duplicateError, err)
}

func TestSearchGroupLeaderSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	leader, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, groupName)

	assert.Equal(t, groupLeaderUsername, leader)
	assert.Equal(t, nil, err)
}

func TestCreateGroupWithNotExistsLeader(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderNotExists, groupID)
	leaderError := errors.New("User does not exist")

	assert.Equal(t, null, group)
	assert.Equal(t, leaderError, err)
}

func TestAddMemberToGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	result, err = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result2, err2 := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username)
	result2, err2 = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username2)

	assert.Contains(t, result, groupLeaderUsername)
	assert.Contains(t, result, username)
	assert.Contains(t, result, username2)
	assert.Equal(t, nil, err)
	
	assert.Contains(t, result2, groupLeaderUsername2)
	assert.Contains(t, result2, username)
	assert.Contains(t, result2, username2)
	assert.Equal(t, nil, err2,)
}

func TestAddGroupToGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, groupName2)

	assert.Contains(t, result, groupName2)
	assert.Equal(t, nil, err)
}

func TestAddMemberToNotExistsGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, []string([]string(nil)), result)
	assert.Equal(t, groupNotExistsError, err)
}

func TestAddMemberToGroupWithNotExistsUser(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, usernameNotExists)
	memberNotExistsError := errors.New("User does not exist")

	assert.Equal(t, []string([]string(nil)), result)
	assert.Equal(t, memberNotExistsError, err)
}

func TestAddDuplicateMemberToGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	memberDuplicateError := errors.New("User already member of the group")

	assert.Equal(t, []string([]string(nil)), result)
	assert.Equal(t, memberDuplicateError, err)
}

func TestGetGroupMembersSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result, err := accountManagement.GetGroupMembers(adminUser, adminPassword, groupName)

	assert.Contains(t, result, leaderID2)
	assert.Contains(t, result, userID)
	assert.Contains(t, result, userID2)
	assert.Equal(t, nil, err)
}

func TestGetUserBelongedTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username)

	result, err := accountManagement.SearchUserMemberOf(adminUser, adminPassword, username)

	assert.Equal(t, groupName, result[0].Name)
	assert.Equal(t, groupID, result[0].UUID)
	assert.Equal(t, groupName2, result[1].Name)
	assert.Equal(t, groupID2, result[1].UUID)
	assert.Equal(t, err, nil)
}

func TestRemoveMemberFromGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName2, username2)

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
	result2, err2 := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName2, username2)

	assert.NotContains(t, result, username)
	assert.Equal(t, nil, err)

	assert.NotContains(t, result2, username2)
	assert.Equal(t, nil, err2)
}

func TestRemoveMemberFromNotExistsGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupNameNotExists, username)
	groupNotExistsError := errors.New("Group does not exist")

	assert.Equal(t, []string([]string(nil)), result)
	assert.Equal(t, groupNotExistsError, err)
}

func TestRemoveNotExistsMemberFromGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, usernameNotExists)
	userNotExistsError := errors.New("User is not a member of group")

	assert.Equal(t, []string([]string(nil)), result)
	assert.Equal(t, userNotExistsError, err)
}

func TestGetGroupUUID(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchGroupUUID(adminUser, adminPassword, groupName)

	assert.Equal(t, "d23475kl-4862-7456-8473-2c53916648fn", result)
	assert.Equal(t, nil, err)
}