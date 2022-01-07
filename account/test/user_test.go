package test

import (
	"errors"
	"fmt"
	"testing"

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

	assert.Equal(t, leaderUsername, result[0].Username)
	assert.Equal(t, fmt.Sprintf("%s %s", leaderGivenName, leaderSurname), result[0].Displayname)
	assert.Equal(t, username2, result[1].Username)
	assert.Equal(t, fmt.Sprintf("%s %s", givenName2, surname2), result[1].Displayname)
	assert.Equal(t, nil, err)
}

// func TestGetAllUsername(t *testing.T) {
// 	defer teardown()
// 	setup()

// 	accountManagement := account.NewLDAPManagement()

// 	result, err := accountManagement.SearchAllUser(adminUser, adminPassword)

// 	member1 := member{
// 		Username:    username,
// 		Displayname: givenName,
// 	}

// 	assert.Contains(t, result, member1)
// 	assert.Equal(t, nil, err)
// }

func TestGetNameByUUID(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchNameByUUID(adminUser, adminPassword, userID)

	assert.Equal(t, username, result)
	assert.Equal(t, nil, err)

	group, err2 := accountManagement.SearchNameByUUID(adminUser, adminPassword, groupID)

	assert.Equal(t, groupName, group)
	assert.Equal(t, nil, err2)
}
