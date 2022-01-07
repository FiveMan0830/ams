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

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username2)
	accountManagement.UpdateGroupLeader(adminUser, adminPassword, groupName, username2)

	result, err := accountManagement.SearchGroupLeader(adminUser, adminPassword, groupName)

	assert.Equal(t, nil, err)
	assert.Equal(t, userID2, result)
}
