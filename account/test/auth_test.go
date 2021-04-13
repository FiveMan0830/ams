package test

import (
	"testing"
	// "errors"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestUserIsMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	
	assert.True(t, accountManagement.IsMember(groupName, username))

	accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
}

func TestUserIsNotMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	
	assert.False(t, accountManagement.IsMember(groupName, usernameNotExists))

	accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
}

func TestUserIsLeaderOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	assert.True(t, accountManagement.IsLeader(groupName, leaderUsername2))
	assert.True(t, accountManagement.IsMember(groupName, leaderUsername2))
}

func TestUserIsNotLeaderOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	assert.False(t, accountManagement.IsLeader(groupName, usernameNotExists))
}

func TestUserIsProfessor(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	accountManagement.CreateOu(adminUser, adminPassword, "Professor")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Cheng134", "Harry", "Cheng", "Professor", "123", "harry@gmail.com")
	
	assert.True(t, accountManagement.IsProfessor("Cheng134"))
	
	accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Cheng134", "Professor")
	accountManagement.DeleteOu(adminUser, adminPassword, "Professor")
}



