package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func TestCreateUserSuccess(t *testing.T) {
	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

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

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	user, err := accountManagement.CreateUser(adminUser, adminPassword, userId1, username1, givenName1, surname1, userPassword1, userEmail1)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestGetUUIDByUsername(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	result, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, username1)

	assert.Equal(t, userId1, result)
	assert.Equal(t, nil, err)
}

func TestGetUUIDByUsernameNotFound(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	uuid, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, usernameNotExists)
	uuidError := errors.New("User not found")

	assert.Equal(t, null, uuid)
	assert.Equal(t, uuidError, err)
}

// func TestGetAllUsername(t *testing.T) {
// 	defer teardown()
// 	setup()

// 	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

// 	result, err := accountManagement.GetAllUsers(adminUser, adminPassword)

// 	member1 := member{
// 		Username:    username,
// 		Displayname: givenName,
// 	}

// 	assert.Contains(t, result, member1)
// 	assert.Equal(t, nil, err)
// }
