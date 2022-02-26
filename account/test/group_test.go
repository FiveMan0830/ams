package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestCreateGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, createGroupErr := accountManagement.CreateGroup(adminUser, adminPassword, groupName3, leaderUsername3, groupId3)
	deleteGroupErr := accountManagement.DeleteGroup(adminUser, adminPassword, groupName3)

	assert.Nil(t, createGroupErr)
	assert.Nil(t, deleteGroupErr)
	assert.Equal(t, groupName3, group)
}

func TestCreateGroupDuplicateName(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderUsername, groupId1)
	duplicateError := errors.New("team already exist")

	assert.Equal(t, null, group)
	assert.Equal(t, duplicateError, err)
}

func TestSearchGroupLeaderSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	leader, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, groupName)

	assert.Equal(t, leaderId1, leader)
	assert.Equal(t, nil, err)
}

func TestCreateGroupWithNotExistsLeader(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderNotExists, groupId1)
	leaderError := errors.New("user not found")

	assert.Equal(t, null, group)
	assert.Equal(t, leaderError, err)
}

func TestAddMemberToGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)
	result, err = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId2)
	if err != nil {
		t.Errorf("failed to add member to group")
	}

	result2, err2 := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId2, userId1)
	result2, err2 = accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId2, userId2)
	if err2 != nil {
		t.Errorf("failed to add member to group")
	}

	isContainLeader := false
	isContainUser1 := false
	isContainUser2 := false
	for _, user := range result {
		if user.Username == groupLeaderUsername {
			isContainLeader = true
		}
		if user.UserID == userId1 {
			isContainUser1 = true
		}
		if user.UserID == userId2 {
			isContainUser2 = true
		}
	}
	assert.True(t, isContainLeader)
	assert.True(t, isContainUser1)
	assert.True(t, isContainUser2)

	isContainLeader = false
	isContainUser1 = false
	isContainUser2 = false
	for _, user := range result2 {
		if user.Username == groupLeaderUsername2 {
			isContainLeader = true
		}
		if user.UserID == userId1 {
			isContainUser1 = true
		}
		if user.UserID == userId2 {
			isContainUser2 = true
		}
	}
	assert.True(t, isContainLeader)
	assert.True(t, isContainUser1)
	assert.True(t, isContainUser2)
}

// will add a function to add a team to another team in the future
// func TestAddGroupToGroup(t *testing.T) {
// 	defer teardown()
// 	setup()

// 	accountManagement := account.NewLDAPManagement()

// 	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, groupName2)

// 	assert.Contains(t, result, groupName2)
// 	assert.Equal(t, nil, err)
// }

func TestAddMemberToNotExistsGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupNameNotExists, userId1)
	groupNotExistsError := errors.New("team not found")

	assert.Nil(t, result)
	assert.Equal(t, groupNotExistsError, err)
}

func TestAddMemberToGroupWithNotExistsUser(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, usernameNotExists)
	memberNotExistsError := errors.New("user not found")

	assert.Nil(t, result)
	assert.Equal(t, memberNotExistsError, err)
}

func TestAddDuplicateMemberToGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId2)

	result, err := accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)
	memberDuplicateError := errors.New("user already member of the group")

	assert.Nil(t, result)
	assert.Equal(t, memberDuplicateError, err)
}

func TestGetUserBelongedTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId2, userId1)

	result, err := accountManagement.GetUserBelongingTeams(adminUser, adminPassword, username1)

	assert.Equal(t, groupName, result[0].Name)
	assert.Equal(t, groupId1, result[0].Id)
	assert.Equal(t, groupName2, result[1].Name)
	assert.Equal(t, groupId2, result[1].Id)
	assert.Equal(t, err, nil)
}

func TestRemoveMemberFromGroupSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId2)

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId2, userId1)
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId2, userId2)

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupId1, userId1)
	result2, err2 := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupId2, userId2)

	assert.NotContains(t, result, userId1)
	assert.Equal(t, nil, err)

	assert.NotContains(t, result2, username2)
	assert.Equal(t, nil, err2)
}

func TestRemoveMemberFromNotExistsGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupNameNotExists, username1)
	groupNotExistsError := errors.New("team not found")

	assert.Nil(t, result)
	assert.Equal(t, groupNotExistsError, err)
}

func TestRemoveNotExistsMemberFromGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupId1, usernameNotExists)
	userNotExistsError := errors.New("user not found")

	assert.Nil(t, result)
	assert.Equal(t, userNotExistsError, err)
}
