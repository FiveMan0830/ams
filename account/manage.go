package account

// Management is a interface to help user manage accounts
type Management interface {
	IsMember(teamId string, userId string) (bool, error)
	IsLeader(teamName, userId string) bool
	IsTeam(teamName string) bool
	IsProfessor(username string) bool
	IsStakeholder(username string) bool

	CreateUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) (*User, error)
	CreateUserWithRole(adminUser, adminPasswd, userID, username, givenname, surname, role, password, email string) error
	CreateGroup(adminUser, adminPasswd, groupname, username, teamID string) (string, error)

	GetGroupInDetail(adminUser, adminPasswd, teamId string) (*DetailTeam, error)
	GetAllGroupsInDetail(adminUser, adminPassword string) ([]*DetailTeam, error)
	GetTeamLeader(adminUser, adminPasswd, teamId string) (*User, error)
	AddMemberToGroup(adminUser, adminPasswd, teamId, userId string) ([]*MemberRole, error)

	GetAllUsers(adminUser, adminPasswd string) ([]*User, error)

	CreateOu(adminUser, adminPasswd, ouname string) error
	DeleteOu(adminUser, adminPasswd, ouname string) error
	SearchGroupLeader(adminUser, adminPasswd, groupName string) (string, error)
	SearchUser(adminUser, adminPasswd, username string) (string, error)
	SearchUserDisplayname(adminUser, adminPasswd, search string) (string, error)
	SearchUserWithOu(adminUser, adminPasswd, role string) ([]string, error)
	SearchNameByUUID(adminUser, adminPasswd, userID string) (string, error)
	SearchUserDn(adminUser, adminPasswd, search string) (string, error)
	DeleteGroup(adminUser, adminPasswd, cn string) error
	Login(adminUser, adminPasswd, username, password string) (string, error)
	GetGroupMembersDetail(adminUser, adminPasswd, teamId string) ([]*MemberRole, error)
	RemoveMemberFromGroup(adminUser, adminPasswd, teamId, userId string) ([]*MemberRole, error)
	GetUserBelongingTeams(adminUser, adminPasswd, username string) ([]*Team, error)
	GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error)
	DeleteUserByUserId(adminUser, adminPasswd, userId string) error
	DeleteUserWithOu(adminUser, adminPasswd, username, role string) error
	UpdateTeamLeader(adminUser, adminPasswd, groupName, newLeader string) error
	// SearchUserRole(teamName, username string) (Role, error)
	GetUserByID(adminUser, adminPasswd, userID string) (*User, error)
	GetUserByUsername(adminUser, adminPasswd, userName string) (*User, error)
}

type User struct {
	UserID      string `json:"userId"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type MemberRole struct {
	*User
	Role string `json:"role"`
}
