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

	accountManagement.AddMemberToGroup(adminUser, adminPassword, groupId1, userId2)
	accountManagement.UpdateTeamLeader(adminUser, adminPassword, groupId1, userId2)

	leader, err := accountManagement.SearchLeaderByTeamId(adminUser, adminPassword, groupId1)
	if err != nil {
		t.Errorf("failed to get team leader")
	}

	assert.Equal(t, userId2, leader.UserID)
}
