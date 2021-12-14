package account

import (
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// Role enum
type Role int

const (
	Member Role = iota 
	Leader
	Professor
	Stakeholder
	Team
)

func (r Role) String() string {
	return [...]string{"Member", "Leader", "Professor", "Stakeholder", "Team"}[r - 1]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func (lm *LDAPManagement) SearchUserRole(teamName, userID string) (Role, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())
	
	accountManagement := NewLDAPManagement()

	if accountManagement.IsLeader(teamName, userID) {
		return Leader, nil
	} else if accountManagement.IsMember(teamName, userID) {
		return Member, nil
	} else {
		return 5, errors.New("Role didn't get!")
	}
}
