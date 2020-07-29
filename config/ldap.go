package config

import "os"

var (
	dc      string
	ldapURL string
)

func init() {
	dc = os.Getenv("LDAP_DC")
	ldapURL = os.Getenv("LDAP_URL")
}

// GetDC is a getter for dc config. dc is the domain component of users in the LDAP server.
func GetDC() string {
	return dc
}

// GetLDAPURL is a getter for ldapURL. ldapURL is the url to the LDAP server.
func GetLDAPURL() string {
	return ldapURL
}
