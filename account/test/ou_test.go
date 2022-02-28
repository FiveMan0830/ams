package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func TestCreateOUSuccess(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	createOuError := accountManagement.CreateOu(adminUser, adminPassword, ouName2)
	deleteOuError := accountManagement.DeleteOu(adminUser, adminPassword, ouName2)

	assert.Equal(t, nil, createOuError)
	assert.Equal(t, nil, deleteOuError)
}

func TestCreateOUDuplicateName(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	createOuError := accountManagement.CreateOu(adminUser, adminPassword, ouName)
	duplicateOuError := errors.New("this Organization Unit already exists")

	assert.Equal(t, duplicateOuError, createOuError)
}
