package config

import "os"

var (
	// AdminUser is the user name of admin
	adminUser string
	// AdminPassword is the user password of admin
	adminPassword string
)

func init() {
	adminUser = os.Getenv("LDAP_ADMIN_USER")
	adminPassword = os.Getenv("LDAP_ADMIN_PASSWORD")
}

// GetAdminUser is the getter for getting admin user name from the config.
func GetAdminUser() string {
	return adminUser
}

// GetAdminPassword is the getter for getting admin password from the config.
func GetAdminPassword() string {
	return adminPassword
}
