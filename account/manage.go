package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string)
	CreateGroup(adminUser, adminPasswd, groupname, username string) ([]*ldap.EntryAttribute, error)
	GetGroups(adminUser, adminPasswd string) ([]string, error)
	AddOu(adminUser, adminPasswd, ouname string)
	AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchGroupLeader(adminUser, adminPasswd, groupname string) ([]string, error)
	SearchUser(adminUser, adminPasswd, username string) (bool)
	DeleteGroup(adminUser, adminPasswd, cn string) (error) 
	Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) 
	GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string, error)
	RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	SearchUserMemberOf(adminUser, adminPasswd, user string) ([]string, error)
}
