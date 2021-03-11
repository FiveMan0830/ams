package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestCreateUserSuccess(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	createUserErr := accountManagement.CreateUser(adminUser, adminPassword, userID3, username3, givenName3, surname3, userPassword3, userEmail3)
	result, searchUserErr := accountManagement.SearchUser(adminUser, adminPassword, username3)
	deleteUserErr := accountManagement.DeleteUser(adminUser, adminPassword, username3)

	assert.Equal(t, createUserErr, nil)
	assert.Equal(t, searchUserErr, nil)
	assert.Equal(t, deleteUserErr, nil)
	
	assert.Equal(t, result, userID3)
}

func TestCreateDuplicateUser(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	duplicateUser := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	duplicateError := errors.New("User already exist")

	assert.Equal(t, duplicateUser, duplicateError)
}

func TestSearchUserSuccess(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, username)

	assert.Equal(t, result, userID)
	assert.Equal(t, err, nil)
}

func TestSearchUserNotFound(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, usernameNotExists)
	searchError := errors.New("User not found")

	assert.Equal(t, result, null)
	assert.Equal(t, err, searchError)
}

func TestGetUUIDByUsername(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	uuid, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, username)

	assert.Equal(t, uuid, userID)
	assert.Equal(t, err, nil)
}

func TestGetUUIDByUsernameNotFound(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	uuid, err := accountManagement.GetUUIDByUsername(adminUser, adminPassword, usernameNotExists)
	uuidError := errors.New("User not found")

	assert.Equal(t, uuid, null)
	assert.Equal(t, err, uuidError)
}