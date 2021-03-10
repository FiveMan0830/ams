package test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestOUNameDuplicate(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	ou := accountManagement.CreateOu(adminUser, adminPassword, ouName)
	duplicateError := errors.New("This Organization Unit already exists")

	assert.Equal(t, ou, duplicateError)
}