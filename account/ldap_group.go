package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
)

// type member struct {
// 	Username string `json:"username"`
// 	Displayname string `json:"displayname"`
// }
type memberRole struct {
	UserID string `json:"id"`
	Role   string `json:"role"`
}

type MemberWithRole struct {
	User
	Role database.Role `json:"role"`
}

// ObjectCategoryGroup is const for ldap attribute
const ObjectCategoryGroup string = "groupOfNames"

// LDAPManagement implement Management interface to connect to LDAP

// CreateGroup is a function for user to create group
func (lm *LDAPManagement) CreateGroup(adminUser, adminPasswd, groupName, username, teamId string) (string, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	baseDN := config.GetDC()
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), []ldap.Control{})

	_, err := lm.GetUserByUsername(adminUser, adminPasswd, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	_, err = lm.GetGroupInDetail(adminUser, adminPasswd, teamId)
	if err == nil {
		return "", errors.New("team already exist")
	}

	addReq.Attribute("objectClass", []string{"top", ObjectCategoryGroup, "UidObject"})
	addReq.Attribute("cn", []string{groupName})
	addReq.Attribute("o", []string{username})
	addReq.Attribute("member", []string{fmt.Sprintf("cn=%s,%s", username, baseDN)})
	addReq.Attribute("uid", []string{teamId})

	if err := conn.Add(addReq); err != nil {
		log.Println("error adding group:", addReq, err)
		return "", err
	}

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(groupName))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{})
	result, err := conn.Search(request)

	if err != nil {
		log.Println("error searching group:", request, err)
		return "", err
	}

	group := strings.Join(result.Entries[0].GetAttributeValues("cn"), "")
	return group, nil
}

// GetGroupMembers is to get members inside of group
func (lm *LDAPManagement) GetGroupMembersRoleDepre(adminUser, adminPasswd, groupName string) ([]*memberRole, error) {
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
		log.Println("error :", err)
		return nil, err
	}

	memberResult := []*memberRole{}
	memberDnList := sr.Entries[0].GetAttributeValues("member")

	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberUUID, err := lm.SearchUser(adminUser, adminPasswd, memberDN)
		teamID, err := lm.SearchGroupUUID(adminUser, adminPasswd, groupName)
		role, err := database.GetRole(memberUUID, teamID)

		if err != nil {
			log.Println("error :", err)
			return nil, err
		}

		mem := new(memberRole)
		mem.UserID = memberUUID
		mem.Role = role.String()

		memberResult = append(memberResult, mem)
	}

	return memberResult, nil
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
		log.Println("error :", err)
		return nil, err
	}

	var memberIDList []string

	memberDnList := sr.Entries[0].GetAttributeValues("member")

	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		memberUUID, err := lm.SearchUser(adminUser, adminPasswd, memberDN)

		if err != nil {
			log.Println("error :", err)
			return nil, err
		}

		memberIDList = append(memberIDList, memberUUID)
	}

	return memberIDList, nil
}

func (lm *LDAPManagement) GetGroupMembersDetail(adminUser, adminPasswd, teamId string) ([]*MemberRole, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	baseDN := config.GetDC()
	searchRequest := ldap.NewSearchRequest(
		baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=groupOfNames)(uid=%s))", teamId),
		[]string{"member"},
		nil,
	)

	sr, err := conn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
		return nil, err
	}

	var memberList []*MemberRole

	memberDnList := sr.Entries[0].GetAttributeValues("member")
	for _, memberDn := range memberDnList {
		memberDn = strings.Replace(memberDn, "cn=", "", -1)
		userName := strings.Replace(memberDn, fmt.Sprintf(",%s", baseDN), "", -1)

		member, _ := lm.GetUserByUsername(adminUser, adminPasswd, userName)
		role, _ := database.GetRole(member.UserID, teamId)
		memberRole := &MemberRole{member, role.String()}

		memberList = append(memberList, memberRole)
	}

	return memberList, nil
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
	leaderID, err := lm.SearchUser(adminUser, adminPasswd, leader)

	return leaderID, nil
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

	return teamID, nil
}

// AddMemberToGroup is for adding member to group
func (lm *LDAPManagement) AddMemberToGroup(adminUser, adminPasswd, teamId, userId string) ([]*MemberRole, error) {
	// check if user exist
	user, err := lm.GetUserByID(adminUser, adminPasswd, userId)
	if err != nil && err.Error() == "user not found" {
		return nil, err
	}

	// check if team exist
	team, err := lm.GetGroupInDetail(adminUser, adminPasswd, teamId)
	if err != nil && err.Error() == "team not found" {
		return nil, err
	}

	isMember, err := lm.IsMember(teamId, userId)
	if err != nil {
		return nil, err
	}

	if !isMember {
		conn, _ := lm.getConnectionWithoutTLS()
		defer conn.Close()
		lm.bindAuth(conn, adminUser, adminPasswd)
		baseDN := config.GetDC()
		modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", team.Name, baseDN), []ldap.Control{})
		modify.Add("member", []string{fmt.Sprintf("cn=%s,%s", user.Username, baseDN)})
		err := conn.Modify(modify)

		if err != nil {
			return nil, errors.New("failed to add user to group")
		}
	} else {
		log.Printf("User %s is already a member of the group %s\n", user.Username, team.Name)
		return nil, errors.New("user already member of the group")
	}

	members, err := lm.GetGroupMembersDetail(adminUser, adminPasswd, teamId)

	return members, nil
}

// RemoveMemberFromGroup is a function to remove a user from a group
func (lm *LDAPManagement) RemoveMemberFromGroup(adminUser, adminPasswd, teamId, userId string) ([]*MemberRole, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	// check if the team exist
	team, err := lm.GetGroupInDetail(adminUser, adminPasswd, teamId)
	if err != nil {
		return nil, err
	}

	// check if the user exist
	user, err := lm.GetUserByID(adminUser, adminPasswd, userId)
	if err != nil {
		return nil, err
	}

	// check if the user is the member
	var isMember bool
	for _, member := range team.Members {
		if member.UserID == userId {
			isMember = true
			break
		}
	}
	if !isMember {
		return nil, fmt.Errorf("user %s is not a member of team %s", userId, teamId)
	}

	// query
	baseDN := config.GetDC()
	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", team.Name, baseDN), []ldap.Control{})
	modify.Delete("member", []string{fmt.Sprintf("cn=%s,%s", user.Username, baseDN)})
	err = conn.Modify(modify)
	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	// get new members
	members, err := lm.GetGroupMembersDetail(adminUser, adminPasswd, teamId)
	if err != nil {
		return nil, err
	}

	return members, nil
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
		[]string{"cn", "uid"},
		[]ldap.Control{},
	)
	sr, err := lm.ldapConn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
	}

	var groupsList []string

	for _, entry := range sr.Entries {
		groupsList = append(groupsList, entry.GetAttributeValue("uid"))
	}

	return groupsList, err
}

func (lm *LDAPManagement) SearchLeaderByTeamId(adminUser, adminPasswd, teamId string) (*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	lm.bindAuth(conn, adminUser, adminPasswd)
	defer conn.Close()

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(teamId))
	request := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"o"},
		[]ldap.Control{},
	)
	result, err := conn.Search(request)

	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	leaderId := strings.Join(result.Entries[0].GetAttributeValues("o"), "")
	leader, err := lm.GetUserByUsername(adminUser, adminPasswd, leaderId)
	if err != nil {
		log.Println(fmt.Errorf("failed to get team leader"))
		return nil, err
	}

	return leader, nil
}

func (lm *LDAPManagement) GetGroupInDetail(adminUser, adminPasswd, teamId string) (*DetailTeam, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	baseDN := config.GetDC()
	searchRequest := ldap.NewSearchRequest(
		baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=groupOfNames)(uid=%s))", teamId),
		[]string{"member", "cn"},
		nil,
	)

	sr, err := conn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
		return nil, err
	}

	if len(sr.Entries) < 1 {
		return nil, errors.New("team not found")
	}

	cn := sr.Entries[0].GetAttributeValue("cn")

	leader, err := lm.SearchLeaderByTeamId(adminUser, adminPasswd, teamId)
	if err != nil {
		return nil, err
	}

	members, err := lm.GetGroupMembersDetail(adminUser, adminPasswd, teamId)
	if err != nil {
		return nil, err
	}

	detailTeam := DetailTeam{
		Id:      teamId,
		Name:    cn,
		Leader:  leader,
		Members: members,
	}

	return &detailTeam, nil
}

func (lm *LDAPManagement) GetAllGroupsInDetail(adminUser, adminPasswd string) ([]*DetailTeam, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(objectClass=%s)", ldap.EscapeFilter(ObjectCategoryGroup))
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn", "uid"},
		[]ldap.Control{},
	)
	sr, err := conn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
	}

	groupsList := []*DetailTeam{}

	for _, entry := range sr.Entries {
		teamId := entry.GetAttributeValue("uid")

		detailTeam, _ := lm.GetGroupInDetail(adminUser, adminPasswd, teamId)

		groupsList = append(groupsList, detailTeam)
	}

	return groupsList, err
}

// DeleteGroup is a function to delete the group
func (lm *LDAPManagement) DeleteGroup(adminUser, adminPasswd, groupName string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	// baseDN := config.GetDC()
	// ou := "OISGroup"
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), nil)
	err := conn.Del(d)

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
		log.Println("error :", err)
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

func getRoleOfMember(roleID int) string {
	if roleID == 0 {
		return "MEMBER"
	} else if roleID == 1 {
		return "LEADER"
	} else if roleID == 2 {
		return "PROFESSOR"
	} else if roleID == 3 {
		return "STAKEHOLDER"
	} else {
		return "NO_ROLE"
	}
}
