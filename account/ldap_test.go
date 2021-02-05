package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const adminUser string = "admin"
const adminPassword string = "admin"

// User test data
const userIDNull = ""
const userID = "c61965be-8176-4419-b289-4d52617728fb"
const userID2 = "c56515be-1654-7895-1564-1d52513528cf"
const username = "testUser"
const username2 = "testUser2"
const givenName = "testUser"
const surname = "testUser"
const userPassword = "testUser"
const userEmail = "test@gmail.com"

// Group test data
const groupNull = ""
const groupName string = "testGroup"
const groupLeader string = "testLeader"
const groupLeader2 string = "testLeader2"
const groupMember string = "testMember"

// Ou test data
const ouName = "testOu"

func TestCreateUser(t *testing.T) {
	accountManagement := NewLDAPManagement()
	user := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	
	assert.Equal(t, user, nil)
}

func TestUserDuplicate(t *testing.T) {
	accountManagement := NewLDAPManagement()
	user := accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	userError := errors.New("User already exist")
	
	assert.Equal(t, user, userError)
}

func TestSearchUser(t *testing.T) {
	accountManagement := NewLDAPManagement()
	result, err := accountManagement.SearchUser(adminUser, adminPassword, username)

	assert.Equal(t, result, userID)
	assert.Equal(t, err, nil)
}

// func TestUserNotFound(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	result, err := accountManagement.SearchUser(adminUser, adminPassword, username2)
// 	searchError := errors.New("User not found")

// 	assert.Equal(t, result, userIDNull)
// 	assert.Equal(t, err, searchError)
// }

func TestCreateGroup(t *testing.T) {
	accountManagement := NewLDAPManagement()
	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader)

	assert.Equal(t, group, groupName)
	assert.Equal(t, err, nil)
}

func TestGroupNameDuplicate(t *testing.T) {
	accountManagement := NewLDAPManagement()
	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader)
	duplicateError := errors.New("Duplicate Group Name")

	assert.Equal(t, group, groupNull)
	assert.Equal(t, err, duplicateError)
}

func TestGroupLeaderNotExists(t *testing.T) {
	accountManagement := NewLDAPManagement()
	group, err := accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader2)
	leaderError := errors.New("User does not exist")

	assert.Equal(t, group, groupNull)
	assert.Equal(t, err, leaderError)
}
 
func TestCreateOU(t *testing.T) {
	accountManagement := NewLDAPManagement()
	ou := accountManagement.CreateOu(adminUser, adminPassword, ouName)

	assert.Equal(t, ou, nil)
}

func TestTearDown(t *testing.T) {
	accountManagement := NewLDAPManagement()
	deleteUser := accountManagement.DeleteUser(adminUser, adminPassword, username)
	deleteGroup := accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	deleteOu := accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	assert.Equal(t, deleteUser, nil)
	assert.Equal(t, deleteGroup, nil)
	assert.Equal(t, deleteOu, nil)
}


// func TestSearchUser(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	username := "testing"
// 	uid, err := accountManagement.SearchUser(adminUser, adminPass, username)
// 	userID := []string{"c61965be-8176-4419-b289-4d52617728fb"}
// 	if err != nil {
// 		log.Fatalf("Error creating groups %+v", err)
// 	}
// 	for i, user := range uid {
// 		if (user != userID[i]){
// 			log.Fatalf("User not found")
// 		} else {
// 			log.Printf("User %s is found with id = %s", username,user)
// 		}
// 	}
// }

// func TestGetUUIDByUsername(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	username := "testing"
// 	uid, err := accountManagement.GetUUIDByUsername(adminUser, adminPass, username)
// 	userID := "c61965be-8176-4419-b289-4d52617728fb"
// 	if err != nil {
// 		log.Fatalf("Error finding user %+v", err)
// 	}
// 	if (userID != uid){
// 		log.Fatalf("User not found")
// 	} else {
// 		log.Printf("User %s is found with id = %s", username,uid)
// 	}
// }

// func TestGetGroupMembers(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	result,err := accountManagement.GetGroupMembers(adminUser, adminPass, "Testing Test")
// 	memberList := []string{"Test", "ssl1321ois"}
// 	if err != nil {
// 		log.Fatalf("Error Getting Member %+v", err)
// 	}
// 	for i, member := range result {
// 		if (member != memberList[i]){
// 			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
// 		} else {
// 			log.Printf("User %s is in the group", member)
// 		}
// 	}
// }

// func TestSearchGroupLeader(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	groupName := "Testing Test"
// 	result,err := accountManagement.SearchGroupLeader(adminUser, adminPass, groupName)
// 	leader := "Test"
// 	if err != nil {
// 		log.Fatalf("Error Getting Leader %+v", err)
// 	}
// 	if (result != leader){
// 		log.Fatalf("User %s isn't the group leader", result)
// 	} else {
// 		log.Printf("User %s is %s leader", result,groupName)
// 	}

// }

// func TestSearchUserMemberOf(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	groupList := []string{"Testing Test"}
// 	username := "Test"
// 	result,err := accountManagement.SearchUserMemberOf(adminUser, adminPass, username)
// 	if err != nil {
// 		log.Fatalf("Error Getting Group %+v", err)
// 	}
// 	for i, group := range result {
// 		if (group != groupList[i]){
// 			log.Fatalf("User %s aren't in The Group %s, it should be Group %s", username, group, groupList[i])
// 		} else {
// 			log.Printf("User %s is in %s Group", username, group)
// 		}
// 	}
// }

// func TestAddMemberToGroup(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	memberList := []string{"Test", "ssl1321ois" ,"AddTest"}
// 	userToAdd := "AddTest"
// 	groupName := "Testing Test"
// 	result,err := accountManagement.AddMemberToGroup(adminUser, adminPass, groupName, userToAdd)
// 	if err != nil {
// 		log.Fatalf("Error Adding Member %+v", err)
// 	}
// 	for i, member := range result {
// 		if (member != memberList[i]){
// 			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
// 		} else {
// 			log.Printf("User %s is in the group", member)
// 		}
// 	}
// }

// func TestRemoveMemberFromGroup(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	memberList := []string{"Test", "ssl1321ois"}
// 	userToRemove := "AddTest"
// 	groupName := "Testing Test"
// 	result,err := accountManagement.RemoveMemberFromGroup(adminUser, adminPass, groupName, userToRemove)
// 	if err != nil {
// 		log.Fatalf("Error Adding Member %+v", err)
// 	}
// 	for i, member := range result {
// 		if (member != memberList[i]){
// 			log.Fatalf("User %s aren't in the group, it should be %s", member, memberList[i])
// 		} else {
// 			log.Printf("User %s is in the group", member)
// 		}
// 	}
// }

// func TestGetGroups(t *testing.T) {
// 	accountManagement := NewLDAPManagement()
// 	groupList := []string{"Testing Test"}
// 	result,err := accountManagement.GetGroups(adminUser, adminPass)
// 	if err != nil {
// 		log.Fatalf("Error Getting Group %+v", err)
// 	}
// 	for i, group := range result {
// 		if (group != groupList[i]){
// 			log.Fatalf("Group %s not found ", group)
// 		} else {
// 			log.Printf("Group %s is found ", group)
// 		}
// 	}
// }
