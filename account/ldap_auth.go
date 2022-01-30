package account

import (
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func (lm *LDAPManagement) IsMember(teamId string, userID string) (bool, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, config.GetAdminUser(), config.GetAdminPassword())

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

func (lm *LDAPManagement) IsLeader(teamName, userId string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	leader, err := lm.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		return false
	}

	if leader == userId {
		return true
	} else {
		return false
	}
}

func (lm *LDAPManagement) IsTeam(teamID string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	teamList, err := lm.GetGroups(config.GetAdminUser(), config.GetAdminPassword())

	if err != nil {
		return false
	}

	for _, team := range teamList {
		if team == teamID {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsProfessor(userID string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

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
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

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
