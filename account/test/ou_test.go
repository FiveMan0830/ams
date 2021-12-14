package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestCreateOUSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	createOuError := accountManagement.CreateOu(adminUser, adminPassword, ouName2)
	deleteOuError := accountManagement.DeleteOu(adminUser, adminPassword, ouName2)

	assert.Equal(t, nil, createOuError)
	assert.Equal(t, nil, deleteOuError)
}

func TestCreateOUDuplicateName(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	createOuError := accountManagement.CreateOu(adminUser, adminPassword, ouName)
	duplicateOuError := errors.New("This Organization Unit already exists")

	assert.Equal(t, duplicateOuError, createOuError)
}