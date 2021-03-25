package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// LDAPManagement implement Management interface to connect to LDAP
type LDAPManagement struct {
	ldapConn *ldap.Conn
}

// CreateUser is a function for user to register
func (lm *LDAPManagement) CreateUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email string) error {
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
		return errors.New("User already exist")
	}

	return nil
}

// DeleteUser is for removing user from ldap
func (lm *LDAPManagement) DeleteUser(adminUser, adminPasswd, username string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	baseDN := config.GetDC()
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", username, baseDN), nil)
	err := lm.ldapConn.Del(d)
	
	if err != nil {
		log.Println("User could not be deleted :", err)
		return err
	} 
	
	return nil
}

// SearchUser is a function to search a user
func (lm *LDAPManagement) SearchUser(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"uid"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		return "", errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}
	user := strings.Join(result.Entries[0].GetAttributeValues("uid"), "")
	return user, nil
}

// SearchNameByUUID is for search name of user or group by their UUID
func (lm *LDAPManagement) SearchNameByUUID(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		return "", errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}
	user := strings.Join(result.Entries[0].GetAttributeValues("cn"), "")
	return user, nil
}

// SearchUserMemberOf is for search group that user belong
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

<<<<<<< HEAD
=======
// AddMemberToGroup is a function to add a user to a group
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
			return membersIdList, errors.New("Failed to add user to group")
		}
	} else {
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
	} else {
		return nil, errors.New("User is not a member of group")
	}
	membersList := lm.GetMemberNoConn(adminUser, adminPasswd, groupName)
	return membersList, nil
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
	// baseDN := config.GetDC()
	// ou := "OISGroup"
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), nil)
	err := lm.ldapConn.Del(d)
	if err != nil {
		log.Println("Group entry could not be deleted :", d, err)
		return err
	} else {
		return nil
	}
}

>>>>>>> master
// GetUUIDByUsername is a function to get username to get UUID
func (lm *LDAPManagement) GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error) {
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
		return "", fmt.Errorf("failed to query LDAP: %w", err)
	}

	if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}

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
	}

	return true
}