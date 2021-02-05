package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	CreateUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) error
	CreateGroup(adminUser, adminPasswd, groupname, username string) (string, error)
	GetGroups(adminUser, adminPasswd string) ([]string, error)
	CreateOu(adminUser, adminPasswd, ouname string) error
	DeleteOu(adminUser, adminPasswd, ouname string) error
	AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchGroupLeader(adminUser, adminPasswd, groupname string) (string, error)
	SearchUser(adminUser, adminPasswd, username string) (string, error)
	DeleteGroup(adminUser, adminPasswd, cn string) error
	Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error)
	GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string, error)
	RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchUserMemberOf(adminUser, adminPasswd, user string) ([]string, error)
	GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error)
	DeleteUser(adminUser, adminPasswd, username string) error 
}
