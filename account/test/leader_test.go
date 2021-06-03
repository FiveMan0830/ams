package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestHandoverLeader(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, "OIS", "stella83")
	accountManagement.UpdateGroupLeader(adminUser, adminPassword, "OIS", "stella83")
	
	result, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, "OIS")

	assert.Equal(t, nil, err)
	assert.Equal(t, "a56515be-5783-8738-1564-1d52513528cz", result)
}
