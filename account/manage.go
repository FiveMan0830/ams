package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string)
	CreateGroup(adminUser, adminPasswd, groupname string) ([]*ldap.EntryAttribute, error)
	AddOu(adminUser, adminPasswd, ouname string)
	AddMemberToGroup(adminUser, adminPasswd, groupName, username string)
	GroupExists(adminUser, adminPasswd, groupname string)
	SearchUser(adminUser, adminPasswd, username string)
	SearchGroupMembers(adminUser, adminPasswd, groupname string)
	Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) //adminUser, adminPasswd,
}
