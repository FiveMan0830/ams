package account

import (
	// "errors"
	// "fmt"
	// "log"
	// "strings"

	// ldap "github.com/go-ldap/ldap/v3"

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