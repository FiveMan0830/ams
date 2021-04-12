package test

import (
	"testing"
	// "errors"

	"github.com/stretchr/testify/assert"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func TestUserIsMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	
	assert.True(t, accountManagement.IsMember(groupName, username))

	accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
}

func TestUserIsNotMemberOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupName, username)
	
	assert.False(t, accountManagement.IsMember(groupName, usernameNotExists))

	accountManagement.RemoveMemberFromGroup(adminUser, adminPassword, groupName, username)
}

func TestUserIsLeaderOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	assert.True(t, accountManagement.IsLeader(groupName, leaderUsername2))
	assert.True(t, accountManagement.IsMember(groupName, leaderUsername2))
}

func TestUserIsNotLeaderOfGroup(t *testing.T) {
	accountManagement := account.NewLDAPManagement()

	assert.False(t, accountManagement.IsLeader(groupName, usernameNotExists))
}



