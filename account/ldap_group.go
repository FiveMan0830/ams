package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// ObjectCategoryGroup is const for ldap attribute
const ObjectCategoryGroup string = "groupOfNames"

// LDAPManagement implement Management interface to connect to LDAP

// CreateGroup is a function for user to create group
func (lm *LDAPManagement) CreateGroup(adminUser, adminPasswd, groupName, username, teamID string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), []ldap.Control{})

	if !lm.SearchUserNoConn(adminUser, adminPasswd, username) {
		return "", errors.New("User does not exist")
	}
	if lm.GroupExists(adminUser, adminPasswd, groupName) {
		return "", errors.New("Duplicate Group Name")
	}
	addReq.Attribute("objectClass", []string{"top", ObjectCategoryGroup, "UidObject"})
	addReq.Attribute("cn", []string{groupName})
	addReq.Attribute("o", []string{username})
	addReq.Attribute("member", []string{fmt.Sprintf("cn=%s,%s", username, baseDN)})
	addReq.Attribute("uid", []string{teamID})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding group:", addReq, err)
		return "", err
	}

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(groupName))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)
	if err != nil {
		log.Println("error searching group:", request, err)
		return "", err
	}
	group := strings.Join(result.Entries[0].GetAttributeValues("cn"), "")
	return group, nil
}

// GetGroupMembers is to get members inside of group
func (lm *LDAPManagement) GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=groupOfNames))",
		[]string{"member"},
		nil,
	)
	sr, err := lm.ldapConn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var memberIDList []string
	memberDnList := sr.Entries[0].GetAttributeValues("member")
	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberIDList = append(memberIDList, memberDN)
	}
	return memberIDList, nil
}

// SearchGroupLeader is for searching the group leader
func (lm *LDAPManagement) SearchGroupLeader(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", search, baseDN),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"o"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return "", err
	}
	leader := strings.Join(result.Entries[0].GetAttributeValues("o"), "")

	return leader, nil
}

// SearchGroupUUID is for searching the group UUID by using group name
func (lm *LDAPManagement) SearchGroupUUID(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", search, baseDN),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"uid"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return "", err
	}
	teamID := strings.Join(result.Entries[0].GetAttributeValues("uid"), "")
	log.Println("team id : " + teamID)
	return teamID, nil

}

// AddMemberToGroup is for adding member to group
func (lm *LDAPManagement) AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	if !lm.GroupExists(adminUser, adminPasswd, groupName) {
		return nil, errors.New("Group does not exist")
	}
	if !lm.SearchUserNoConn(adminUser, adminPasswd, username) {
		return nil, errors.New("User does not exist")
	}

	memberExists := false
	membersIDList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	for _, memberUsername := range membersIDList {
		if memberUsername == username {
			memberExists = true
			break
		}
	}

	if !memberExists {
		baseDN := config.GetDC()
		modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN), []ldap.Control{})
		modify.Add("member", []string{fmt.Sprintf("cn=%s,%s", username, baseDN)})
		err := lm.ldapConn.Modify(modify)
		if err != nil {
			return membersIDList, errors.New("Failed to add user to group")
		}
	} else {
		log.Printf("User %s is already a member of the group %s\n", username, groupName)
		return nil, errors.New("User already member of the group")
	}

	memberList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	return memberList, nil
}

// RemoveMemberFromGroup is a function to remove a user from a group
func (lm *LDAPManagement) RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	if !lm.GroupExists(adminUser, adminPasswd, groupName) {
		return nil, errors.New("Group does not exist")
	}
	memberExists := false
	membersIDList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	for _, memberUsername := range membersIDList {
		if memberUsername == username {
			memberExists = true
			break
		}
	}
	if memberExists {
		baseDN := config.GetDC()
		modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN), []ldap.Control{})
		modify.Delete("member", []string{fmt.Sprintf("cn=%s,%s", username, baseDN)})
		err := lm.ldapConn.Modify(modify)
		if err != nil {
			log.Println(fmt.Errorf("failed to query LDAP: %w", err))
			return nil, err
		}
	} else {
		return nil, errors.New("User is not a member of group")
	}
	membersList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	return membersList, nil
}

// GetGroups is for getting all group inside of ldap
func (lm *LDAPManagement) GetGroups(adminUser, adminPasswd string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	ou := "test"
	if ou == "" {
		log.Fatal("ou is a required paramater for getting a list of groups")
	}
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(objectClass=%s)", ldap.EscapeFilter(ObjectCategoryGroup))
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{},
	)
	sr, err := lm.ldapConn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	var groupsList []string
	for _, entry := range sr.Entries {
		groupsList = append(groupsList, entry.GetAttributeValue("cn"))
	}

	return groupsList, err
}

// DeleteGroup is a function to delete the group
func (lm *LDAPManagement) DeleteGroup(adminUser, adminPasswd, groupName string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	// baseDN := config.GetDC()
	// ou := "OISGroup"
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), nil)
	err := lm.ldapConn.Del(d)
	if err != nil {
		log.Println("Group entry could not be deleted :", d, err)
		return err
	}

	return nil
}

// GetMemberNoConn is for searching member without connection
func (lm *LDAPManagement) GetMemberNoConn(adminUser, adminPasswd, groupName string) []string {
	baseDN := config.GetDC()
	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=groupOfNames))",
		[]string{"member"},
		nil,
	)
	sr, err := lm.ldapConn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	var memberIDList []string
	memberDnList := sr.Entries[0].GetAttributeValues("member")
	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberIDList = append(memberIDList, memberDN)
	}
	return memberIDList
}

// GroupExists is for checking the group is exists or not in ldap  
func (lm *LDAPManagement) GroupExists(adminUser, adminPasswd, search string) bool {
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"ou"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
	}
	if len(result.Entries) < 1 {
		return false
	}

	return true
}