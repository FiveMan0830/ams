package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

// func TestCreateUserSuccess(t *testing.T) {
// 	accountManagement := account.NewLDAPManagement()

// 	createUserErr := accountManagement.CreateUser(adminUser, adminPassword, userID3, username3, givenName3, surname3, userPassword3, userEmail3)
// 	result, searchUserErr := accountManagement.SearchUser(adminUser, adminPassword, username3)
// 	deleteUserErr := accountManagement.DeleteUser(adminUser, adminPassword, username3)

// 	assert.Equal(t, createUserErr, nil)
// 	assert.Equal(t, searchUserErr, nil)
// 	assert.Equal(t, deleteUserErr, nil)
	
// 	assert.Equal(t, result, userID3)
// }

func TestCreateDuplicateUser(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	duplicateUser := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	duplicateError := errors.New("User already exist")

	assert.Equal(t, duplicateError, duplicateUser)
}

func TestSearchUserSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, username)

	assert.Equal(t, userID, result)
	assert.Equal(t, nil, err)
}

func TestSearchUserNotFound(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, usernameNotExists)
	searchError := errors.New("User not found")

	assert.Equal(t, null, result)
	assert.Equal(t, searchError, err)
}

func TestGetUUIDByUsername(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, username)

	assert.Equal(t, userID, result)
	assert.Equal(t, nil, err)
}

func TestGetUUIDByUsernameNotFound(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	uuid, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, usernameNotExists)
	uuidError := errors.New("User not found")

	assert.Equal(t, null, uuid)
	assert.Equal(t, uuidError, err)
}

func TestGetListOfMemberUsernameAndDisplaynameByTeamName(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)

	result, err := accountManagement.GetGroupMembersUsernameAndDisplayname(adminUser, adminPassword, groupName)

	assert.Equal(t, "david93", result[0].Username)
	assert.Equal(t, "David Wang", result[0].Displayname)
	assert.Equal(t, "audi98", result[1].Username)
	assert.Equal(t, "Audi Wu", result[1].Displayname)
	assert.Equal(t, nil, err)
}

func TestGetAllUsername(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchAllUser(adminUser, adminPassword)

	member1 := new(member)
	member1.Displayname = "test"
	member1.Username = "test"

	assert.Equal(t, result[0].Username, member1.Username)
	assert.Equal(t, nil, err)
}

func TestGetNameByUUID(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchNameByUUID(adminUser, adminPassword, "c61965be-8176-4419-b289-4d52617728fb")

	assert.Equal(t, "fiveman123", result)
	assert.Equal(t, nil, err)

	group, err2 := accountManagement.SearchNameByUUID(adminUser, adminPassword, "d23475kl-4862-7456-8473-2c53916648fn")

	assert.Equal(t, "OIS", group)
	assert.Equal(t, nil, err2)
}