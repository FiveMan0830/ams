package account

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func (lm *LDAPManagement) IsMember(teamName, userID string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	teamMemberList, err := lm.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		return false
	}

	for _, teamMember := range teamMemberList {
		if teamMember == userID {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsLeader(teamName, username string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	userID, err := lm.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(),username)
	if err != nil {
		return false
	}

	leader, err := lm.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), teamName)
	
	if err != nil {
		return false
	}

	if leader == userID {
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
		return false;
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
