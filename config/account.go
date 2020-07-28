package config

import "os"

var (
	// AdminUser is the user name of admin
	AdminUser string
	// AdminPassword is the user password of admin
	AdminPassword string
)

func init() {
	AdminUser = os.Getenv("LDAP_ADMIN_USER")
	AdminPassword = os.Getenv("LDAP_ADMIN_PASSWORD")
}
