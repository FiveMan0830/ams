package account

import (
	"testing"
	"log"
 )

 func TestCreateGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	groupName,err := accountManagement.CreateGroup("admin","admin","TestTeam0231")
	if err != nil {
		log.Fatalf("Error creating groups %+v", err)
	}
	log.Printf("Groups: %+v", groupName)
 }