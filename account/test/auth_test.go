package test

import (
	"testing"

	// "errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestUserIsMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	
	defer func() {
		accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
		teardown()
	}()

	setup()
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)

	assert.True(t, accountManagement.IsMember(groupName, userID))
}

func TestUserIsNotMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	
	defer func() {
		accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
		teardown()
	}()

	setup()
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)

	assert.False(t, accountManagement.IsMember(groupName, usernameNotExists))
}

func TestUserIsLeaderOfGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	assert.True(t, accountManagement.IsLeader(groupName, leaderID2))
	assert.True(t, accountManagement.IsMember(groupName, leaderID2))
}

func TestUserIsNotLeaderOfGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	assert.False(t, accountManagement.IsLeader(groupName, usernameNotExists))
}

func TestIsTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	assert.True(t, accountManagement.IsTeam(groupName))
}

func TestIsNotTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	assert.False(t, accountManagement.IsTeam(groupNameNotExists))
}

func TestUserIsProfessor(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Cheng134", "Professor")
		accountManagement.DeleteOu(adminUser, adminPassword, "Professor")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Professor")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Cheng134", "Harry", "Cheng", "Professor", "123", "harry@gmail.com")

	assert.True(t, accountManagement.IsProfessor(userID))
}

func TestUserIsNotProfessor(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Cheng134", "Professor")
		accountManagement.DeleteOu(adminUser, adminPassword, "Professor")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Professor")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Cheng134", "Harry", "Cheng", "Professor", "123", "harry@gmail.com")

	assert.False(t, accountManagement.IsProfessor(usernameNotExists))
}

func TestUserIsStakeholder(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Wang134", "Stakeholder")
		accountManagement.DeleteOu(adminUser, adminPassword, "Stakeholder")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Stakeholder")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Wang134", "Eric", "Wangg", "Stakeholder", "9865", "eric@gmail.com")

	assert.True(t, accountManagement.IsStakeholder(userID))
}

func TestUserIsNotStakeholder(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Wang134", "Stakeholder")
		accountManagement.DeleteOu(adminUser, adminPassword, "Stakeholder")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Stakeholder")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Wang134", "Eric", "Wangg", "Stakeholder", "9865", "eric@gmail.com")

	assert.False(t, accountManagement.IsStakeholder(usernameNotExists))
}
