package account

import (
	"fmt"
	"log"

	ldap "github.com/go-ldap/ldap/v3"
)

func (lm *LDAPManagement) UpdateTeamLeader(adminUser, adminPasswd, teamId, newLeaderId string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	// check if the team exist
	team, err := lm.GetGroupInDetail(adminUser, adminPasswd, teamId)
	if err != nil {
		return err
	}

	// check if the user exist
	newLeader, err := lm.GetUserByID(adminUser, adminPasswd, newLeaderId)
	if err != nil {
		return err
	}

	// check if the user is the member
	var isMember bool
	for _, member := range team.Members {
		if member.UserID == newLeaderId {
			isMember = true
			break
		}
	}
	if !isMember {
		return fmt.Errorf("user %s is not a member of team %s", newLeaderId, teamId)
	}

	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", team.Name, lm.BaseDN), []ldap.Control{})
	modify.Replace("o", []string{fmt.Sprintf(newLeader.Username)})
	err = conn.Modify(modify)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return fmt.Errorf(
			"failed to update leader to %s for team %s", newLeaderId, teamId,
		)
	}

	return nil
}
