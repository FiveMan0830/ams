package account

import (
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func (lm *LDAPManagement) IsMember(teamId string, userID string) (bool, error) {
	members, err := lm.GetGroupMembersDetail(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
	)
	if err != nil {
		return false, errors.New("failed to get members")
	}

	for _, member := range members {
		if member.UserID == userID {
			return true, nil
		}
	}

	return false, nil
}

func (lm *LDAPManagement) IsLeader(teamId, userId string) bool {
	leader, err := lm.GetTeamLeader(config.GetAdminUser(), config.GetAdminPassword(), teamId)

	if err != nil {
		return false
	}

	if leader.UserID == userId {
		return true
	} else {
		return false
	}
}

func (lm *LDAPManagement) IsTeam(teamID string) bool {
	teamList, err := lm.GetAllGroupsInDetail(config.GetAdminUser(), config.GetAdminPassword())

	if err != nil {
		return false
	}

	for _, team := range teamList {
		if team.Id == teamID {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsProfessor(userID string) bool {
	professorList, err := lm.SearchUserWithOu(config.GetAdminUser(), config.GetAdminPassword(), "Professor")

	if err != nil {
		return false
	}

	for _, professor := range professorList {
		if professor == userID {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsStakeholder(userID string) bool {
	stakeholderList, err := lm.SearchUserWithOu(config.GetAdminUser(), config.GetAdminPassword(), "Stakeholder")

	if err != nil {
		return false
	}

	for _, stakeholder := range stakeholderList {
		if stakeholder == userID {
			return true
		}
	}

	return false
}
