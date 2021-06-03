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

type member struct {
	Username string `json:"username"`
	Displayname string `json:"displayname"`
	// Role int `json:"role"`
}

type team struct {
	Name string `json:"name"`
	UUID string `json:"id"`
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

// CreateUser is a function for user with role to register
func (lm *LDAPManagement) CreateUserWithOu(adminUser, adminPasswd, userID, username, givenname, surname, role, password, email string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=%s,%s", username, role, config.GetDC()), []ldap.Control{})
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

// GetUserByID is for getting user's info from ldap
func (lm *LDAPManagement) GetUserByID(adminUser, adminPasswd, userID string) (*User, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(userID))
	log.Println(userID)
	
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		0, 0, 0, false, filter,
		[]string{"uid", "cn", "displayname",  "mail"},
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

	user := &User{
		UserID: userID,
		Username: result.Entries[0].GetAttributeValue("cn"),
		DisplayName: result.Entries[0].GetAttributeValue("displayName"),
		Email: result.Entries[0].GetAttributeValue("mail"),
	}

	return user, nil
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

func (lm *LDAPManagement) DeleteUserWithOu(adminUser, adminPasswd, username, role string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	baseDN := config.GetDC()
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=%s,%s", username, role, baseDN), nil)
	err := lm.ldapConn.Del(d)

	if err != nil {
		log.Println("User could not be deleted :", err)
		return err
	}
	return nil
}

// SearchUser is a function to search a user
func (lm *LDAPManagement) SearchAllUser(adminUser, adminPasswd string) ([]*member, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		[]string{"cn"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		return nil, errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		return nil, errors.New("There is no user")
	}

	userList := []*member{}
	
	for _, entry := range result.Entries {
		mem := new(member)
		mem.Username = entry.GetAttributeValue("cn")
		memberDisplayname, err := lm.SearchUserDisplayname(adminUser, adminPasswd, mem.Username)
		mem.Displayname = memberDisplayname
		fmt.Println(mem.Username + ", " + mem.Displayname)
		
		userList = append(userList, mem)
		if err != nil {
			return nil, errors.New("Search Failed")
		}
	}

	return userList, nil
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

func (lm *LDAPManagement) SearchUserDisplayname(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"displayName"},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		return "", errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}

	user := strings.Join(result.Entries[0].GetAttributeValues("displayName"), "")

	return user, nil
}

// SearchUser is a function to search a user that have a role
func (lm *LDAPManagement) SearchUserWithOu(adminUser, adminPasswd, role string) ([]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(ou:dn:=%s))", ldap.EscapeFilter(role))
	request := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{"uid"},
		[]ldap.Control{})

	result, err := lm.ldapConn.Search(request)

	if err != nil {
		log.Println("Search Failed")
		return nil, errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		log.Println("User not found")
		return nil, errors.New("User not found")
	}

	var userList []string

	for _, entry := range result.Entries {
		userList = append(userList, entry.GetAttributeValue("uid"))
	}

	return userList, nil
}

// SearchUser is a function to search a user dn
func (lm *LDAPManagement) SearchUserDn(adminUser, adminPasswd, search string) (string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{},
		[]ldap.Control{})
	result, err := lm.ldapConn.Search(request)

	if err != nil {
		return "", errors.New("Search Failed")
	} else if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}

	// user := strings.Join(result.Entries[0], "")

	return result.Entries[0].DN, nil
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
func (lm *LDAPManagement) SearchUserMemberOf(adminUser, adminPasswd, user string) ([]*team, error) {
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
		[]string{"dn", "cn", "uid"},
		[]ldap.Control{})

	sr, err := lm.ldapConn.Search(searchRequest)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	teams := []*team{}

	for _, entry := range sr.Entries {
		tm := new(team)
		tm.Name = entry.GetAttributeValue("cn")
		tm.UUID = entry.GetAttributeValue("uid")
		fmt.Println(tm.Name + ", " + tm.UUID)
		
		teams = append(teams, tm)
	}

	return teams, nil
}

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
func (lm *LDAPManagement) Login(adminUser, adminPasswd, username, password string) (*ldap.Entry, error) {
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

	return result.Entries[0], nil
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


