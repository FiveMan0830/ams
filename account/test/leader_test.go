package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestHandoverLeader(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	accountManagement.UpdateGroupLeader(adminUser, adminPassword, "OIS", "stella83")
	
	result, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, "OIS")

	assert.Equal(t, "stella83", result)
	assert.Equal(t, err, nil)
}
