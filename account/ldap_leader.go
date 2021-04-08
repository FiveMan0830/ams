package account

import (
	"errors"
	"fmt"
	"log"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func (lm *LDAPManagement) UpdateGroupLeader(adminUser, adminPasswd, groupName, newLeader string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	if !lm.GroupExists(adminUser, adminPasswd, groupName) {
		return errors.New("Group does not exist")
	}
	if !lm.SearchUserNoConn(adminUser, adminPasswd, newLeader) {
		return errors.New("User does not exist")
	}

	baseDN := config.GetDC()
	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN), []ldap.Control{})
	modify.Replace("o", []string{fmt.Sprintf(newLeader)})
	err := lm.ldapConn.Modify(modify)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return errors.New("Failed to update new leader!")
	} else {
		return errors.New("User is not a member of group")
	}
}