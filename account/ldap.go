package account

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

// LDAPManagement implement Management interface to connect to LDAP
type LDAPManagement struct {
	ldapConn *ldap.Conn
}

// AddGroup is a function for user to create group
func (lm *LDAPManagement) AddGroup(username, password, groupname string) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(username, password)

	addReq := ldap.NewAddRequest(fmt.Sprintf("CN=%s,ou=Groups,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", groupname), []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{groupname})
	addReq.Attribute("sAMAccountName", []string{groupname})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x00000004|0x80000000)})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding group:", addReq, err)
	}
}

// AddUser is a function for user to register
func (lm *LDAPManagement) AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", username), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("givenname", []string{givenname})
	addReq.Attribute("sn", []string{surname})
	addReq.Attribute("displayname", []string{givenname + " " + surname})
	addReq.Attribute("userPassword", []string{password})
	addReq.Attribute("uid", []string{userID})
	addReq.Attribute("mail", []string{email})

	// addReq.Attribute("userAccountControl", []string{fmt.Sprintf("%d", 0x0202)})
	// addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	//
	// addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", 0x00000000)})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding service:", addReq, err)
	}
}

// Login is a function for user to login and get information
func (lm *LDAPManagement) Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	// lm.bind(username, password)
	lm.bind(adminUser, adminPasswd)
	baseDN := "dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw"
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(username))

	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		0, 0, 0, false, filter,
		[]string{"description", "sn", "cn", "displayname", "userPassword", "uid", "mail"},
		[]ldap.Control{})

	result, err := lm.ldapConn.Search(searchReq)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	userdn := result.Entries[0].DN

	// Bind as the user to verify their password
	err = lm.ldapConn.Bind(userdn, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result.Entries[0].Attributes, nil

	// for _, attribute := range result.Entries[0].Attributes {
	// 	fmt.Printf("%s: %v\n", attribute.Name, attribute.Values)
	// }
}

func (lm *LDAPManagement) connectWithoutTLS() {
	ldapURL := "ldap://140.124.181.94:389"
	var err error
	lm.ldapConn, err = ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
}

func (lm *LDAPManagement) bind(username, password string) {
	err := lm.ldapConn.Bind(fmt.Sprintf("cn=%s,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", username), password)

	if err != nil {
		log.Fatal(err)
	}
}

// NewLDAPManagement is a factory method to generate LDAPManagement
func NewLDAPManagement() Management {
	return &LDAPManagement{}
}
