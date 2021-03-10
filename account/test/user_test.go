package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestUserDuplicate(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	user := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	userError := errors.New("User already exist")

	assert.Equal(t, user, userError)
}

func TestSearchUser(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUser(adminUser, adminPassword, username)

	assert.Equal(t, result, userID)
	assert.Equal(t, err, nil)
}

func TestUserNotFound(t *testing.T) {
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