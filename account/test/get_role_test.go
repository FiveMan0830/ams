package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestGetUserRoleMember(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroupDepre(adminUser, adminPassword, groupName, username)

	result, err := accountManagement.SearchUserRole(groupName, userID)

	assert.Equal(t, 0, result.EnumIndex())
	assert.Equal(t, nil, err)
}

func TestGetUserRoleLeader(t *testing.T) {
	defer teardown()
	setup()

	accountManagement := account.NewLDAPManagement()

	result, err := accountManagement.SearchUserRole(groupName, leaderID2)

	assert.Equal(t, 1, result.EnumIndex())
	assert.Equal(t, nil, err)
}
