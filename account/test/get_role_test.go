package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestGetUserRoleMember(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)

	result, err := accountManagement.SearchUserRole(groupName, userID)

	assert.Equal(t, 1, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleLeader(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUserRole(groupName, leaderID2)

	assert.Equal(t, 2, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleTeam(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUserRole(groupName, "")

	assert.Equal(t, 5, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleProfessor(t *testing.T) {
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
	

	result, err := accountManagement.SearchUserRole("", userID)

	assert.Equal(t, 3, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleStakeholder(t *testing.T) {
	accountManagement := account.NewLDAPManagement()
	userID := uuid.New().String()

	defer func() {
		accountManagement.DeleteUserWithOu(adminUser, adminPassword, "Wang134", "Stakeholder")
		accountManagement.DeleteOu(adminUser, adminPassword, "Stakeholder")
		teardown()
	}()

	setup()
	accountManagement.CreateOu(adminUser, adminPassword, "Stakeholder")
	accountManagement.CreateUserWithOu(adminUser, adminPassword, userID, "Wang134", "Eric", "Wang", "Stakeholder", "9865", "eric@gmail.com")
	
	result, err := accountManagement.SearchUserRole("", userID)

	assert.Equal(t, 4, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleException(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUserRole(groupName, userID2)

	assert.Equal(t, 0, result.EnumIndex())
	assert.Equal(t,  errors.New("Role didn't get!"), err)
}