package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestCreateUserSuccess(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	user, err := accountManagement.CreateUser(adminUser, adminPassword, userId3, username3, givenName3, surname3, userPassword3, userEmail3)

	assert.Nil(t, err)
	assert.Equal(t, userId3, user.UserID)
	assert.Equal(t, userEmail3, user.Email)
	assert.Equal(t, username3, user.Username)

	err = accountManagement.DeleteUserByUserId(adminUser, adminPassword, userId3)

	assert.Nil(t, err)
}

func TestCreateDuplicateUser(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	user, err := accountManagement.CreateUser(adminUser, adminPassword, userId1, username1, givenName1, surname1, userPassword1, userEmail1)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestSearchUserSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, username1)

	assert.Equal(t, userId1, result)
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

	result, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, username1)

	assert.Equal(t, userId1, result)
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

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId2)

	result, err := accountManagement.GetGroupMembersUsernameAndDisplayname(adminUser, adminPassword, groupName)

	fmt.Println(result)

	assert.Equal(t, leaderUsername1, result[0].Username)
	assert.Equal(t, fmt.Sprintf("%s %s", leaderGivenName1, leaderSurname1), result[0].Displayname)
	assert.Equal(t, username2, result[1].Username)
	assert.Equal(t, fmt.Sprintf("%s %s", givenName2, surname2), result[1].Displayname)
	assert.Equal(t, nil, err)
}

// func TestGetAllUsername(t *testing.T) {
// 	defer teardown()
// 	setup()

// 	accountManagement := account.NewLDAPManagement()

// 	result, err := accountManagement.GetAllUsers(adminUser, adminPassword)

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

	result, err := accountManagement.SearchNameByUUID(adminUser, adminPassword, userId1)

	assert.Equal(t, username1, result)
	assert.Equal(t, nil, err)

	group, err2 := accountManagement.SearchNameByUUID(adminUser, adminPassword, groupId1)

	assert.Equal(t, groupName, group)
	assert.Equal(t, nil, err2)
}
