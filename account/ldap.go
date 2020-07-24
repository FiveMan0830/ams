package account

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

type LDAPManagement struct {
	ldapConn *ldap.Conn
}

func (lm *LDAPManagement) AddGroup() {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind()

	addReq := ldap.NewAddRequest("CN=testgroup,ou=Groups,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{"testgroup"})
	addReq.Attribute("sAMAccountName", []string{"testgroup"})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x00000004|0x80000000)})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Fatal("error adding group:", addReq, err)
	}
}

func (lm *LDAPManagement) AddUser(username, surname, password string) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind()

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", username), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "person"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("sn", []string{surname})
	addReq.Attribute("userPassword", []string{password})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Fatal("error adding service:", addReq, err)
	}
}

// func Search(l *ldap.Conn){
// 	user := "fooUser"
// 	baseDN := "DC=example,DC=com"
// 	filter := fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(user))

// 	// Filters must start and finish with ()!
// 	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{"sAMAccountName"}, []ldap.Control{})

// 	result, err := l.Search(searchReq)
// 	if err != nil {
// 		return fmt.Errorf("failed to query LDAP: %w", err)
// 	}

// 	log.Println("Got", len(result.Entries), "search results")
// }

func (lm *LDAPManagement) connectWithoutTLS() {
	ldapURL := "ldap://140.124.181.94:389"
	var err error
	lm.ldapConn, err = ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
}

func (lm *LDAPManagement) bind() {
	err := lm.ldapConn.Bind("cn=admin,dc=ssl-drone,dc=csie,dc=ntut,dc=edu,dc=tw", "admin")
	if err != nil {
		log.Fatal(err)
	}
}

func NewLDAPManagement() AccountManagement {
	return &LDAPManagement{}
}
