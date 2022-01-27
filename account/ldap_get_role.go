package account

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
	return [...]string{"Member", "Leader", "Professor", "Stakeholder", "Team"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
}

// func (lm *LDAPManagement) SearchUserRole(teamId string, userId string) (Role, error) {
// 	lm.connectWithoutTLS()
// 	defer lm.ldapConn.Close()
// 	lm.bind(config.GetAdminUser(), config.GetAdminPassword())

// 	accountManagement := NewLDAPManagement()

// 	isLeader := accountManagement.IsLeader(teamId, userId)
// 	isMember, err := accountManagement.IsMember(teamId, userId)

// 	if accountManagement.IsLeader(teamName, userID) {
// 		return Leader, nil
// 	} else if accountManagement.IsMember(teamName, userID) {
// 		return Member, nil
// 	} else {
// 		return 5, errors.New("Role didn't get!")
// 	}
// }
