package account

import (
	"log"
	"testing"
)

const adminUser string = "admin"
const adminPass string = "admin"

func TestCreateGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	groupName, err := accountManagement.CreateGroup(adminUser, adminPass, "TestTeam0231", "ssl1321ois")
	if err != nil {
		log.Fatalf("Error creating groups %+v", err)
	}
	log.Printf("Groups: %+v", groupName)
}

func TestDeleteGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	err := accountManagement.DeleteGroup(adminUser, adminPass, "TestTeam0231")
	if err != nil {
		log.Fatalf("Error Deleting groups %+v", err)
	}
}

func TestCreateOu(t *testing.T) {
	accountManagement := NewLDAPManagement()
	err := accountManagement.AddOu(adminUser, adminPass, "OIS")
	if err != nil {
		log.Fatalf("Error creating Organization Unit %+v", err)
	}
}

func TestDeleteOu(t *testing.T) {
	accountManagement := NewLDAPManagement()
	err := accountManagement.DeleteOu(adminUser, adminPass, "OIS")
	if err != nil {
		log.Fatalf("Error Deleting Organization Unit %+v", err)
	}
}

 func TestAddUser(t *testing.T) {
	accountManagement := NewLDAPManagement()
	userID := "c61965be-8176-4419-b289-4d52617728fb"
	username := "testing"
	givenname := "test"
	surname := "test"
	password := "test"
	email := "test@mail.com"
	err := accountManagement.AddUser(adminUser,adminPass,userID,username,givenname,surname,password,email)
	if err != nil {
		log.Fatalf("Error Adding User %+v", err)
	}
}

func TestSearchUser(t *testing.T) {
	accountManagement := NewLDAPManagement()
	username := "testing"
	uid, err := accountManagement.SearchUser(adminUser, adminPass, username)
	userID := []string{"c61965be-8176-4419-b289-4d52617728fb"}
	if err != nil {
		log.Fatalf("Error creating groups %+v", err)
	}
	for i, user := range uid {
		if (user != userID[i]){
			log.Fatalf("User not found")
		} else {
			log.Printf("User %s is found with id = %s", username,user)
		}
	}
}

func TestGetUUIDByUsername(t *testing.T) {
	accountManagement := NewLDAPManagement()
	username := "testing"
	uid, err := accountManagement.GetUUIDByUsername(adminUser, adminPass, username)
	userID := "c61965be-8176-4419-b289-4d52617728fb"
	if err != nil {
		log.Fatalf("Error finding user %+v", err)
	}
	if (userID != uid){
		log.Fatalf("User not found")
	} else {
		log.Printf("User %s is found with id = %s", username,uid)
	}
}

func TestRemoveUser(t *testing.T) {
	accountManagement := NewLDAPManagement()
	err := accountManagement.RemoveUser(adminUser, adminPass, "testing")
	if err != nil {
		log.Fatalf("Error Removing User %+v", err)
	}
}

func TestGetGroupMembers(t *testing.T) {
	accountManagement := NewLDAPManagement()
	result,err := accountManagement.GetGroupMembers(adminUser, adminPass, "Testing Test")
	memberList := []string{"Test", "ssl1321ois"}
	if err != nil {
		log.Fatalf("Error Getting Member %+v", err)
	} 
	for i, member := range result {
		if (member != memberList[i]){
			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
		} else {
			log.Printf("User %s is in the group", member)
		}
	}
}

func TestSearchGroupLeader(t *testing.T) {
	accountManagement := NewLDAPManagement()
	groupName := "Testing Test"
	result,err := accountManagement.SearchGroupLeader(adminUser, adminPass, groupName)
	leader := "Test"
	if err != nil {
		log.Fatalf("Error Getting Leader %+v", err)
	} 
	if (result != leader){
		log.Fatalf("User %s isn't the group leader", result)
	} else {
		log.Printf("User %s is %s leader", result,groupName)
	}
	
}

func TestSearchUserMemberOf(t *testing.T) {
	accountManagement := NewLDAPManagement()
	groupList := []string{"Testing Test"}
	username := "Test"
	result,err := accountManagement.SearchUserMemberOf(adminUser, adminPass, username)
	if err != nil {
		log.Fatalf("Error Getting Group %+v", err)
	} 
	for i, group := range result {
		if (group != groupList[i]){
			log.Fatalf("User %s aren't in The Group %s, it should be Group %s", username, group, groupList[i])
		} else {
			log.Printf("User %s is in %s Group", username, group)
		}
	}
}

func TestAddMemberToGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	memberList := []string{"Test", "ssl1321ois" ,"AddTest"}
	userToAdd := "AddTest"
	groupName := "Testing Test"
	result,err := accountManagement.AddMemberToGroup(adminUser, adminPass, groupName, userToAdd)
	if err != nil {
		log.Fatalf("Error Adding Member %+v", err)
	} 
	for i, member := range result {
		if (member != memberList[i]){
			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
		} else {
			log.Printf("User %s is in the group", member)
		}
	}
}

func TestRemoveMemberFromGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	memberList := []string{"Test", "ssl1321ois"}
	userToRemove := "AddTest"
	groupName := "Testing Test"
	result,err := accountManagement.RemoveMemberFromGroup(adminUser, adminPass, groupName, userToRemove)
	if err != nil {
		log.Fatalf("Error Adding Member %+v", err)
	} 
	for i, member := range result {
		if (member != memberList[i]){
			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
		} else {
			log.Printf("User %s is in the group", member)
		}
	}
}

func TestGetGroups(t *testing.T) {
	accountManagement := NewLDAPManagement()
	groupList := []string{"Testing Test"}
	result,err := accountManagement.GetGroups(adminUser, adminPass)
	if err != nil {
		log.Fatalf("Error Getting Group %+v", err)
	} 
	for i, group := range result {
		if (group != groupList[i]){
			log.Fatalf("Group %s not found ", group)
		} else {
			log.Printf("Group %s is found ", group)
		}
	}
}








