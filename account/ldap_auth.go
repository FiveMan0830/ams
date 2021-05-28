package account

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func (lm *LDAPManagement) IsMember(teamName, username string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	teamMemberList, err := lm.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		return false
	}

	for _, teamMember := range teamMemberList {
		if teamMember == username {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsLeader(teamName, username string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	leader, err := lm.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		return false
	}

	if leader == username {
		return true
	} else {
		return false
	}
}

func (lm *LDAPManagement) IsProfessor(username string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	professorList, err := lm.SearchUserWithOu(config.GetAdminUser(), config.GetAdminPassword(), "Professor")

	if err != nil {
		return false
	}

	for _, professor := range professorList {
		if professor == username {
			return true
		}
	}

	return false
}

func (lm *LDAPManagement) IsStakeholder(username string) bool {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

	stakeholderList, err := lm.SearchUserWithOu(config.GetAdminUser(), config.GetAdminPassword(), "Stakeholder")

	if err != nil {
		return false
	}

	for _, stakeholder := range stakeholderList {
		if stakeholder == username {
			return true
		}
	}

	return false
}
