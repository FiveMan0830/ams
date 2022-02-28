package test

import (
	"testing"

	// "errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func TestUserIsMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	defer func() {
		accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupId1, userId1)
		teardown()
	}()

	setup()
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)

	result, err := accountManagement.IsMember(groupId1, userId1)
	if err != nil {
		t.Errorf("failed to check membership")
	}
	assert.True(t, result)
}

func TestUserIsNotMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	defer func() {
		accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupId1, userId1)
		teardown()
	}()

	setup()
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId1)

	result, err := accountManagement.IsMember(groupId1, userId2)
	if err != nil {
		t.Errorf("failed to get membership")
	}
	assert.False(t, result)
}

func TestUserIsLeaderOfGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, leaderId2)

	assert.True(t, accountManagement.IsLeader(groupId1, leaderId1))

	result, err := accountManagement.IsMember(groupId1, leaderId2)
	if err != nil {
		t.Errorf("failed to get membership")
	}
	assert.True(t, result)
}

func TestUserIsNotLeaderOfGroup(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	assert.False(t, accountManagement.IsLeader(groupId1, usernameNotExists))
}

func TestIsTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	assert.True(t, accountManagement.IsTeam(groupId1))
}

func TestIsNotTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	assert.False(t, accountManagement.IsTeam(groupNameNotExists))
}

func TestUserIsProfessor(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Cheng134", "Professor")
		accountManagement.DeleteOu(adminUser, adminPassword, "Professor")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Professor")
	accountManagement.CreateUserWithRole(adminUser, adminPassword, userID, "Cheng134", "Harry", "Cheng", "Professor", "123", "harry@gmail.com")

	assert.True(t, accountManagement.IsProfessor(userID))
}

func TestUserIsNotProfessor(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Cheng134", "Professor")
		accountManagement.DeleteOu(adminUser, adminPassword, "Professor")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Professor")
	accountManagement.CreateUserWithRole(adminUser, adminPassword, userID, "Cheng134", "Harry", "Cheng", "Professor", "123", "harry@gmail.com")

	assert.False(t, accountManagement.IsProfessor(usernameNotExists))
}

func TestUserIsStakeholder(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Wang134", "Stakeholder")
		accountManagement.DeleteOu(adminUser, adminPassword, "Stakeholder")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Stakeholder")
	accountManagement.CreateUserWithRole(adminUser, adminPassword, userID, "Wang134", "Eric", "Wangg", "Stakeholder", "9865", "eric@gmail.com")

	assert.True(t, accountManagement.IsStakeholder(userID))
}

func TestUserIsNotStakeholder(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Wang134", "Stakeholder")
		accountManagement.DeleteOu(adminUser, adminPassword, "Stakeholder")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Stakeholder")
	accountManagement.CreateUserWithRole(adminUser, adminPassword, userID, "Wang134", "Eric", "Wangg", "Stakeholder", "9865", "eric@gmail.com")

	assert.False(t, accountManagement.IsStakeholder(usernameNotExists))
}
