package account

import (
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// Role enum
type Role int

const (
	Member Role = iota + 1
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

func (lm *LDAPManagement) SearchUserRole(teamName, username string) (Role, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(config.GetAdminUser(), config.GetAdminPassword())
	
	accountManagement := NewLDAPManagement()

	if accountManagement.IsProfessor(username) && teamName == "" {
		return Professor, nil
	} else if accountManagement.IsStakeholder(username) && teamName == "" {
		return Stakeholder, nil
	} else if accountManagement.IsTeam(teamName) && username == "" {
		return Team, nil 
	} else if accountManagement.IsMember(teamName, username) {
		return Member, nil
	} else if accountManagement.IsLeader(teamName, username) {
		return Leader, nil
	} else {
		return 0, errors.New("Role didn't get!")
	}
}
