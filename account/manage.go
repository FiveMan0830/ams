package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string)
	CreateGroup(adminUser, adminPasswd, groupname string) ([]*ldap.EntryAttribute, error)
	GetGroups(adminUser, adminPasswd string) ([]string, error)
	AddOu(adminUser, adminPasswd, ouname string)
	AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
	GroupExists(adminUser, adminPasswd, groupname string) (bool)
	SearchUser(adminUser, adminPasswd, username string) (bool)
	DeleteGroup(adminUser, adminPasswd, cn string) (error) 
	Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) 
	GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string, error)
	RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string) ([]string, error)
}
