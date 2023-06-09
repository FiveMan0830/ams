package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

// LDAPManagement implement Management interface to connect to LDAP
type LDAPManagement struct {
	BaseDN string
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DetailTeam struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Leader  *User         `json:"leader"`
	Members []*MemberRole `json:"members"`
}

// CreateUser is a function for user to register
func (lm *LDAPManagement) CreateUser(adminUser, adminPasswd, userId, username, givenName, surName, password, email string) (*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", username, config.GetDC()), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("givenname", []string{givenName})
	addReq.Attribute("sn", []string{surName})
	addReq.Attribute("displayname", []string{givenName + " " + surName})
	addReq.Attribute("userPassword", []string{password})
	addReq.Attribute("uid", []string{userId})
	addReq.Attribute("mail", []string{email})

	if err := conn.Add(addReq); err != nil {
		return nil, errors.New("user already exist")
	}

	user, err := lm.GetUserByID(adminUser, adminPasswd, userId)
	if err != nil {
		return nil, errors.New("failed to get created user")
	}

	return user, nil
}

// CreateUser is a function for user with role to register
func (lm *LDAPManagement) CreateUserWithRole(adminUser, adminPasswd, userID, username, givenname, surname, role, password, email string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=%s,%s", username, role, config.GetDC()), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("givenname", []string{givenname})
	addReq.Attribute("sn", []string{surname})
	addReq.Attribute("displayname", []string{givenname + " " + surname})
	addReq.Attribute("userPassword", []string{password})
	addReq.Attribute("uid", []string{userID})
	addReq.Attribute("mail", []string{email})

	if err := conn.Add(addReq); err != nil {
		return errors.New("User already exist")
	}

	return nil

}

// GetUserByID is for getting user's info from ldap
func (lm *LDAPManagement) GetUserByID(adminUser, adminPasswd, userID string) (*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(userID))
	attributes := []string{"uid", "cn", "displayname", "mail"}

	searchReq := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(searchReq)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	if len(result.Entries) < 1 {
		return nil, errors.New("user not found")
	}

	if len(result.Entries) > 1 {
		return nil, errors.New("too many entries returned")
	}

	user := &User{
		UserID:      userID,
		Username:    result.Entries[0].GetAttributeValue("cn"),
		DisplayName: result.Entries[0].GetAttributeValue("displayName"),
		Email:       result.Entries[0].GetAttributeValue("mail"),
	}

	return user, nil
}

func (lm *LDAPManagement) GetUserByUsername(adminUser, adminPasswd, userName string) (*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(userName))
	attributes := []string{"uid", "cn", "displayname", "mail"}
	searchReq := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(searchReq)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	if len(result.Entries) < 1 {
		return nil, errors.New("user not found")
	}

	if len(result.Entries) > 1 {
		return nil, errors.New("too many entries returned")
	}

	user := &User{
		UserID:      result.Entries[0].GetAttributeValue("uid"),
		Username:    userName,
		DisplayName: result.Entries[0].GetAttributeValue("displayName"),
		Email:       result.Entries[0].GetAttributeValue("mail"),
	}

	return user, nil
}

// DeleteUser is for removing user from ldap
func (lm *LDAPManagement) DeleteUserByUserId(adminUser, adminPasswd, userId string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	user, err := lm.GetUserByID(adminUser, adminPasswd, userId)
	if err != nil {
		return err
	}

	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", user.Username, lm.BaseDN), nil)
	err = conn.Del(d)

	if err != nil {
		log.Println("User could not be deleted :", err)
		return err
	}

	return nil
}

// Delete user who has organization unit
func (lm *LDAPManagement) DeleteUserWithOu(adminUser, adminPasswd, username, role string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=%s,%s", username, role, lm.BaseDN), nil)
	err := conn.Del(d)

	if err != nil {
		log.Println("User could not be deleted :", err)
		return err
	}

	return nil
}

// Get all users
func (lm *LDAPManagement) GetAllUsers(adminUser, adminPasswd string) ([]*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := "(&(objectClass=organizationalPerson))"
	attributes := []string{"cn", "uid", "displayName", "mail"}
	request := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(request)

	if err != nil {
		return nil, errors.New("search failed")
	} else if len(result.Entries) < 1 {
		return nil, errors.New("there is no user")
	}

	userList := []*User{}

	for _, entry := range result.Entries {
		userList = append(userList, &User{
			UserID:      entry.GetAttributeValue("uid"),
			Username:    entry.GetAttributeValue("cn"),
			DisplayName: entry.GetAttributeValue("displayName"),
			Email:       entry.GetAttributeValue("mail"),
		})
	}

	return userList, nil
}

// search user who has organization unit
func (lm *LDAPManagement) SearchUserWithOu(adminUser, adminPasswd, role string) ([]string, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(ou:dn:=%s))", ldap.EscapeFilter(role))
	attributes := []string{"uid"}
	request := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(request)

	if err != nil {
		log.Println("Search Failed")
		return nil, errors.New("search Failed")
	} else if len(result.Entries) < 1 {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}

	var userList []string

	for _, entry := range result.Entries {
		userList = append(userList, entry.GetAttributeValue("uid"))
	}

	return userList, nil
}

// GetUserBelongingTeams is for search group that user belong
func (lm *LDAPManagement) GetUserBelongingTeams(adminUser, adminPasswd, username string) ([]*Team, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	user, err := lm.GetUserByUsername(adminUser, adminPasswd, username)
	if err != nil {
		return nil, err
	}

	filter := fmt.Sprintf("(&(objectClass=groupOfNames)(member=cn=%s,%s))", user.Username, lm.BaseDN)
	attributes := []string{"dn", "cn", "uid"}
	searchRequest := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	sr, err := conn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
		return nil, err
	}

	teams := []*Team{}

	for _, entry := range sr.Entries {
		tm := &Team{
			Id:   entry.GetAttributeValue("uid"),
			Name: entry.GetAttributeValue("cn"),
		}

		teams = append(teams, tm)
	}

	return teams, nil
}

// GetUUIDByUsername is a function to get username to get UUID
func (lm *LDAPManagement) GetUUIDByUsername(adminUser, adminPasswd, username string) (string, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(username))
	attributes := []string{"uid"}
	request := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(request)

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
func (lm *LDAPManagement) Login(adminUser, adminPasswd, username, password string) (string, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(username))
	attributes := []string{"cn", "uid"}

	searchReq := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := conn.Search(searchReq)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return "", err
	}

	if len(result.Entries) < 1 {
		return "", errors.New("user not found")
	}

	if len(result.Entries) > 1 {
		return "", errors.New("too many entries returned")
	}

	userdn := result.Entries[0].DN
	userId := result.Entries[0].GetAttributeValue("uid")

	// Bind as the user to verify their password
	err = conn.Bind(userdn, password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// generate JWT
	token, err := pkg.NewJWTClient(config.NewAuthConfig()).CreateToken(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SearchNameByUUID is for search name of user or group by their UUID
func (lm *LDAPManagement) SearchNameByUUID(adminUser, adminPasswd, search string) (string, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(search))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{})
	result, err := conn.Search(request)

	if err != nil {
		return "", errors.New("search Failed")
	} else if len(result.Entries) < 1 {
		return "", errors.New("User not found")
	}

	user := strings.Join(result.Entries[0].GetAttributeValues("cn"), "")

	return user, nil
}

func (lm *LDAPManagement) getConnectionWithoutTLS() (*ldap.Conn, error) {
	ldapUrl := config.GetLDAPURL()
	conn, err := ldap.DialURL(ldapUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func (lm *LDAPManagement) bindAuth(conn *ldap.Conn, username string, password string) error {
	err := conn.Bind(fmt.Sprintf("cn=%s,%s", username, config.GetDC()), password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type LDAPManagerConfig struct {
	BaseDN string
}

// NewLDAPManagement is a factory method to generate LDAPManagement
func NewLDAPManagement(config LDAPManagerConfig) Management {
	return &LDAPManagement{BaseDN: config.BaseDN}
}
