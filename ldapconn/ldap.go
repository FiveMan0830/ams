package ldapconn

import (
	"github.com/go-ldap/ldap"
	"log"
	"fmt"
)

// ConnectLDAP connect to the LDAP server without TLS
func ConnectLDAP() *ldap.Conn {
	ldapURL := "ldap://localhost:389"
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

func BindLDAP(l *ldap.Conn) {
	err := l.Bind("cn=admin,dc=csie,dc=ntut,dc=edu,dc=tw", "admin")
	if err != nil {
		log.Fatal(err)
	}
}

func AddGroup(l *ldap.Conn){
	addReq := ldap.NewAddRequest("CN=testgroup,ou=Groups,dc=csie,dc=ntut,dc=edu,dc=tw", []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{"testgroup"})
	addReq.Attribute("sAMAccountName", []string{"testgroup"})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x00000004 | 0x80000000)})

	if err := l.Add(addReq); err != nil {
		log.Fatal("error adding group:", addReq, err)
	}
}

func AddUser(l *ldap.Conn){
	addReq := ldap.NewAddRequest("sn=lab,dc=csie,dc=ntut,dc=edu,dc=tw", []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "person"})
	addReq.Attribute("cn", []string{"fooUser"})
	addReq.Attribute("sn", []string{"lab"})

	if err := l.Add(addReq); err != nil {
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