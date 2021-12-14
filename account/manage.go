package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	IsMember(teamName, username string) bool
	IsLeader(teamName, username string) bool
	IsTeam(teamName string) bool
	IsProfessor(username string) bool
	IsStakeholder(username string) bool
	CreateUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) error
	CreateUserWithOu(adminUser, adminPasswd, userID, username, givenname, surname, role, password, email string) error
	CreateGroup(adminUser, adminPasswd, groupname, username, teamID string) (string, error)
	GetGroups(adminUser, adminPasswd string) ([]string, error)
	CreateOu(adminUser, adminPasswd, ouname string) error
	DeleteOu(adminUser, adminPasswd, ouname string) error
	AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchGroupLeader(adminUser, adminPasswd, groupname string) (string, error)
	SearchAllUser(adminUser, adminPasswd string) ([]*member, error)
	SearchUser(adminUser, adminPasswd, username string) (string, error)
	SearchUserDisplayname(adminUser, adminPasswd, search string) (string, error)
	SearchUserWithOu(adminUser, adminPasswd, role string) ([]string, error)
	SearchNameByUUID(adminUser, adminPasswd, userID string) (string, error)
	SearchUserDn(adminUser, adminPasswd, search string) (string, error)
	DeleteGroup(adminUser, adminPasswd, cn string) error
	Login(adminUser, adminPasswd, username, password string) (*ldap.Entry, error)
	GetGroupMembersUsernameAndDisplayname(adminUser, adminPasswd, groupName string) ([]*member, error)
	GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string, error)
	GetGroupMembersRole(adminUser, adminPasswd, groupName string) ([]*memberRole, error)
	RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchUserMemberOf(adminUser, adminPasswd, user string) ([]*team, error)
	GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error)
	DeleteUser(adminUser, adminPasswd, username string) error
	DeleteUserWithOu(adminUser, adminPasswd, username, role string) error
	SearchGroupUUID(adminUser, adminPasswd, groupName string) (string, error)
	UpdateGroupLeader(adminUser, adminPasswd, groupName, newLeader string) error
	SearchUserRole(teamName, username string) (Role, error)
	GetUserByID(adminUser, adminPasswd, userID string) (*User, error)
}

type User struct {
	UserID      string
	Username    string
	DisplayName string
	Email       string
}
