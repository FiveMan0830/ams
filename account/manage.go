package account

import "github.com/go-ldap/ldap/v3"

// Management is a interface to help user manage accounts
type Management interface {
	AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string)
	AddGroup(username, password, groupname string)
	AddOu(username, password, ouname string)
	SearchGroup(username, password, groupname string)
	Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) //adminUser, adminPasswd,
}
