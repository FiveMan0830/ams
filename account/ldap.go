package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

const ObjectCategory_Group string = "groupOfNames"

// LDAPManagement implement Management interface to connect to LDAP
type LDAPManagement struct {
	ldapConn *ldap.Conn
}

// AddGroup is a function for user to create group
func (lm *LDAPManagement) CreateGroup(adminUser, adminPasswd, groupname, groupID, username string) ([]*ldap.EntryAttribute, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupname, config.GetDC()), []ldap.Control{})
	
	if !lm.SearchUserNoConn(adminUser, adminPasswd, username) {
		log.Printf("User %s does not exist", username)
		return nil, errors.New("User does not exist")
	}
	addReq.Attribute("objectClass", []string{"top", ObjectCategory_Group})
	addReq.Attribute("cn", []string{groupname})
	addReq.Attribute("o", []string{username})
	addReq.Attribute("member", []string{fmt.Sprintf("cn=%s,%s", username, baseDN)})
	addReq.Attribute("", []string{groupID})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding group:", addReq, err)
		return nil, err
	}

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(groupname))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"ou"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)
	if err != nil {
		log.Println("error searching group:", request, err)
		return nil, err
	}
	return result.Entries[0].Attributes, nil
}

// AddOu is a function for user to create ou
func (lm *LDAPManagement) AddOu(adminUser, adminPasswd, ouname string) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("ou=%s,%s", ouname, config.GetDC()), []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "organizationalUnit"})
	addReq.Attribute("ou", []string{ouname})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding ou:", addReq, err)
	}
}

// AddUser is a function for user to register
func (lm *LDAPManagement) AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", username, config.GetDC()), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("givenname", []string{givenname})
	addReq.Attribute("sn", []string{surname})
	addReq.Attribute("displayname", []string{givenname + " " + surname})
	addReq.Attribute("userPassword", []string{password})
	addReq.Attribute("uid", []string{userID})
	addReq.Attribute("mail", []string{email})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("error adding service:", addReq, err)
	}
}

// GetGroupMembers is a function to get all the member from a group
func (lm *LDAPManagement) GetGroupMembers(adminUser, adminPasswd, groupName string) ([]string ,error){
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
	var memberIdList []string
	memberDnList := sr.Entries[0].GetAttributeValues("member")
	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberIdList = append(memberIdList, memberDN)
	}
	return memberIdList, nil
}

// GroupExists is a function for get all the groups
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
		return "",err
	} 
	leader := strings.Join(result.Entries[0].GetAttributeValues("o"), "")

	return leader, nil
	
}

// SearchUser is a function to search a user
func (lm *LDAPManagement) SearchUser(adminUser, adminPasswd, search string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"email"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil,err
	}
	log.Printf(result.Entries[0].DN)
	user := result.Entries[0].GetAttributeValues("email")
	return user, nil
}

func (lm *LDAPManagement) SearchUserMemberOf(adminUser, adminPasswd, user string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(&(objectClass=groupOfNames)(member=cn=%s,%s))", user, baseDN)
	searchRequest := ldap.NewSearchRequest(
        baseDN,
		ldap.ScopeWholeSubtree, 
		ldap.NeverDerefAliases, 
		0, 
		0, 
		false,
		filter,
		[]string{"dn", "cn"},
		[]ldap.Control{})

	sr, err := lm.ldapConn.Search(searchRequest)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var groupsList []string
	for _, entry := range sr.Entries {
		groupsList = append(groupsList, entry.GetAttributeValue("cn"))	
	}
	return groupsList, nil
}

// AddMemberToGroup is a function to add a user to a group
func (lm *LDAPManagement) AddMemberToGroup(adminUser, adminPasswd, groupName, username string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	if !lm.GroupExists(adminUser, adminPasswd, groupName) {
		log.Printf("Group %s does not exist, cn or ou is incorrect\n",groupName)
		return nil, errors.New("Group does not exist")
	}
	if !lm.SearchUserNoConn(adminUser, adminPasswd, username) {
		log.Printf("User %s does not exist, hence the user could not be added to the group %s\n", username,groupName)
		return nil, errors.New("User does not exist")
	}
	
	memberExists := false
	membersIdList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	for _, member_username := range membersIdList {
		if member_username == username {
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
			log.Println(fmt.Errorf("failed to query LDAP: %w", err))
			return membersIdList, err
		}
		log.Printf("User %s is added to the group %s\n",  username, groupName)
	} else {
		log.Printf("User %s is already a member of the group %s\n",  username, groupName)
	}

	memberList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	return memberList, nil
}

// RemoveMemberFromGroup is a function to remove a user from a group
func (lm *LDAPManagement) RemoveMemberFromGroup(adminUser, adminPasswd, groupName, username string)([]string,error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	
	if !lm.GroupExists(adminUser, adminPasswd, groupName) {
		log.Printf("Group %s does not exist, cn or ou is incorrect\n", groupName)
		return nil, errors.New("Group does not exist")
	}
	memberExists := false
	membersIdList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	for _, member_username := range membersIdList {
		if member_username == username {
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
		log.Printf("User %s is removed from the group %s\n",  username, groupName)
	} else {
		log.Printf("User %s is not a member of the group %s\n",  username, groupName)
	}
	membersList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	return  membersList, nil
}

// GetGroups is a function to get all the group
func (lm *LDAPManagement) GetGroups(adminUser, adminPasswd string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	ou := "test"
	if ou == "" {
		log.Fatal("ou is a required paramater for getting a list of groups")
	}
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(objectClass=%s)", ldap.EscapeFilter(ObjectCategory_Group))
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
	baseDN := config.GetDC()
	ou := "OISGroup"
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=%s,%s", groupName, ou, baseDN), nil)
	err := lm.ldapConn.Del(d)
	if err != nil {
		log.Println("Group entry could not be deleted :", err)
		return err
	} else {
		log.Printf("Group %s is deleted\n", groupName)
		return nil
	}
}

// GetUUIDByUsername is a function to get username to get UUID
func (lm *LDAPManagement) GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error)  {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(username))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"uid"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return "", err
	}
	log.Printf(result.Entries[0].DN)
	user := strings.Join(result.Entries[0].GetAttributeValues("uid"), "")
	return user, nil
}

// Login is a function for user to login and get information
func (lm *LDAPManagement) Login(adminUser, adminPasswd, username, password string) ([]*ldap.EntryAttribute, error) {
	if err := lm.connectWithoutTLS(); err != nil {
		return nil, err
	}
	defer lm.ldapConn.Close()

	// First bind with a read only user
	if adminPasswd != "" {
		if err := lm.bind(adminUser, adminPasswd); err != nil {
			return nil, err
		}
	}

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(username))

	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		0, 0, 0, false, filter,
		[]string{"description", "sn", "cn", "displayname", "userPassword", "uid", "mail", "givenname"},
		[]ldap.Control{})

	result, err := lm.ldapConn.Search(searchReq)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	if len(result.Entries) < 1 {
		return nil, errors.New("User not found")
	}

	if len(result.Entries) > 1 {
		return nil, errors.New("Too many entries returned")
	}

	userdn := result.Entries[0].DN

	// Bind as the user to verify their password
	err = lm.ldapConn.Bind(userdn, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result.Entries[0].Attributes, nil
}

func (lm *LDAPManagement) connectWithoutTLS() error {
	ldapURL := config.GetLDAPURL()
	var err error
	lm.ldapConn, err = ldap.DialURL(ldapURL)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (lm *LDAPManagement) bind(username, password string) error {
	err := lm.ldapConn.Bind(fmt.Sprintf("cn=%s,%s", username, config.GetDC()), password)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// NewLDAPManagement is a factory method to generate LDAPManagement
func NewLDAPManagement() Management {
	return &LDAPManagement{}
}

// GetMemberNoConn is a function to get all the member from a group
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
	var memberIdList []string
	memberDnList := sr.Entries[0].GetAttributeValues("member")
	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberIdList = append(memberIdList, memberDN)
	}
	return memberIdList
}

// SearchUserNoConn is a function to search a user
func (lm *LDAPManagement) SearchUserNoConn(adminUser, adminPasswd, search string) bool {
	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
	}
	if len(result.Entries) < 1 {
		return false
	} else {
		return true
	}
}

// GroupExists is a function for get all the groups
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
	} else {
		return true
	}
}
