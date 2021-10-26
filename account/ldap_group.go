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

// ObjectCategoryGroup is const for ldap attribute
const ObjectCategoryGroup string = "groupOfNames"

// LDAPManagement implement Management interface to connect to LDAP

// CreateGroup is a function for user to create group
func (lm *LDAPManagement) CreateGroupDepre(adminUser, adminPasswd, groupName, username, teamID string) (string, error) {
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
func (lm *LDAPManagement) GetGroupMembersUsernameAndDisplayname(adminUser, adminPasswd, groupName string) ([]*member, error) {
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

	memberResult := []*member{}
	memberList := sr.Entries[0].GetAttributeValues("member")

	for _, memberEntry := range memberList {
		memberEntry = strings.Replace(memberEntry, "cn=", "", -1)
		memberEntry = strings.Replace(memberEntry, fmt.Sprintf(",%s", baseDN), "", -1)
		memberDisplayname, err := lm.SearchUserDisplayname(adminUser, adminPasswd, memberEntry)

		if err != nil {
			log.Println("error :", err)
			return nil, err
		}

		mem := new(member)
		mem.Username = memberEntry
		mem.Displayname = memberDisplayname

		memberResult = append(memberResult, mem)
	}

	return memberResult, nil
}

// GetGroupMembers is to get members inside of group
func (lm *LDAPManagement) GetGroupMembersRole(adminUser, adminPasswd, groupName string) ([]*memberRole, error) {
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

	fmt.Println(memberDnList)

	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		fmt.Println(memberDN)
		memberUUID, err := lm.SearchUser(adminUser, adminPasswd, memberDN)
		teamID, err := lm.SearchGroupUUID(adminUser, adminPasswd, groupName)
		role, err := database.GetRole(memberUUID, teamID)

		if err != nil {
			log.Println("error :", err)
			return nil, err
		}

		mem := new(memberRole)
		mem.UserID = memberUUID
		mem.Role = MemberRole(role)

		memberResult = append(memberResult, mem)
	}

	return memberResult, nil
}

// GetGroupMembers is to get members inside of group by group name
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

// AddMemberToGroupDepre is for adding member to group
func (lm *LDAPManagement) AddMemberToGroupDepre(adminUser, adminPasswd, groupName, username string) ([]string, error) {
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

func (lm *LDAPManagement) GetAllGroups() ([]map[string]string, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)
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

	var groupsList []map[string]string
	for _, entry := range sr.Entries {
		groupsList = append(groupsList, map[string]string{
			"name": entry.GetAttributeValue("cn"),
			"id":   entry.GetAttributeValue("uid"),
		})
	}

	return groupsList, nil
}

// DeleteGroup is a function to delete the group
func (lm *LDAPManagement) DeleteGroupDepre(adminUser, adminPasswd, groupName string) error {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(adminUser, adminPasswd)
	// baseDN := config.GetDC()
	// ou := "OISGroup"
	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), nil)
	err := lm.ldapConn.Del(d)

	if err != nil {
		log.Println("Group entry could not be deleted:", d, err)
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

func MemberRole(roleID int) string {
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

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

// get group members by group id
func (lm *LDAPManagement) getGroupMembersByGroupId(teamId string) ([]*User, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=groupOfNames)(userid=%s))", teamId),
		[]string{"member"},
		nil,
	)

	sr, err := lm.ldapConn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
		return nil, err
	}

	var memberList []*User

	memberDnList := sr.Entries[0].GetAttributeValues("member")

	for _, memberDN := range memberDnList {
		memberDN = strings.Replace(memberDN, "cn=", "", -1)
		memberDN = strings.Replace(memberDN, fmt.Sprintf(",%s", baseDN), "", -1)
		user, err := lm.GetUserByCn(memberDN)

		if err != nil {
			log.Println("error :", err)
			return nil, err
		}

		memberList = append(memberList, user)
	}

	return memberList, nil
}

// get group by group id
func (lm *LDAPManagement) GetGroupById(teamId string) (*Group, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf(
		"(&(objectClass=%s)(userid=%s))",
		ldap.EscapeFilter(ObjectCategoryGroup),
		teamId,
	)

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn", "uid", "o", "member"},
		[]ldap.Control{},
	)
	result, err := lm.ldapConn.Search(searchRequest)
	if err != nil {
		return nil, errors.New("failed to search group")
	}
	if len(result.Entries) == 0 {
		return nil, errors.New("can not find group with id " + teamId)
	}

	members, err := lm.getGroupMembersByGroupId(teamId)
	if err != nil {
		return nil, errors.New("failed to get team members")
	}

	leader, err := lm.GetUserByCn(result.Entries[0].GetAttributeValue("o"))
	if err != nil {
		fmt.Println(err)
	}

	group := &Group{
		Id:      result.Entries[0].GetAttributeValue("uid"),
		Name:    result.Entries[0].GetAttributeValue("cn"),
		Members: members,
		Leader:  leader,
	}

	return group, nil
}

// get user's teams
func (lm *LDAPManagement) GetGroupsByUserId(userId string) ([]*Group, error) {
	user, err := lm.GetUserById(userId)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(&(objectClass=groupOfNames)(member=cn=%s,%s))", user.Username, baseDN)
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn", "uid"},
		[]ldap.Control{},
	)
	sr, err := lm.ldapConn.Search(searchRequest)

	if err != nil {
		log.Println("error :", err)
		return nil, err
	}

	teams := []*Group{}
	for _, entry := range sr.Entries {
		teams = append(teams, &Group{
			Id:   entry.GetAttributeValue("uid"),
			Name: entry.GetAttributeValue("cn"),
		})
	}

	return teams, nil
}

func (lm *LDAPManagement) GetGroupByGroupName(groupName string) (*Group, error) {
	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(groupName))
	request := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"uid", "o"},
		[]ldap.Control{},
	)
	result, err := lm.ldapConn.Search(request)
	if err != nil {
		return nil, errors.New("failed to search group")
	}
	if len(result.Entries) == 0 {
		return nil, errors.New("group not found: " + groupName)
	}

	members, err := lm.getGroupMembersByGroupId(result.Entries[0].GetAttributeValue("uid"))
	if err != nil {
		return nil, errors.New("failed to get team members")
	}

	leader, err := lm.GetUserByCn(result.Entries[0].GetAttributeValue("o"))
	if err != nil {
		fmt.Println(err)
	}

	team := &Group{
		Id:      result.Entries[0].GetAttributeValue("uid"),
		Name:    groupName,
		Members: members,
		Leader:  leader,
	}

	return team, nil
}

func (lm *LDAPManagement) CreateGroup(groupId string, groupName string, leaderName string) (*Group, error) {
	_, err := lm.GetUserByCn(leaderName)
	if err != nil {
		return nil, err
	}

	group, err := lm.GetGroupByGroupName(groupName)
	if group != nil {
		return nil, errors.New("team already exist")
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()), []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", ObjectCategoryGroup, "UidObject"})
	addReq.Attribute("cn", []string{groupName})
	addReq.Attribute("o", []string{leaderName})
	addReq.Attribute("member", []string{fmt.Sprintf("cn=%s,%s", leaderName, baseDN)})
	addReq.Attribute("uid", []string{groupId})

	if err := lm.ldapConn.Add(addReq); err != nil {
		log.Println("failed to create group:", addReq, err)
		return nil, err
	}

	group, err = lm.GetGroupByGroupName(groupName)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (lm *LDAPManagement) DeleteGroup(groupName string) (*Group, error) {
	group, err := lm.GetGroupByGroupName(groupName)
	if err != nil {
		return nil, err
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	req := ldap.NewDelRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, config.GetDC()),
		nil,
	)
	if err := lm.ldapConn.Del(req); err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to delete group" + err.Error())
	}

	return group, nil
}

func (lm *LDAPManagement) IsTeamMember(userName string, groupName string) (bool, error) {
	team, err := lm.GetGroupByGroupName(groupName)
	fmt.Println("team", team)
	if err != nil {
		return false, errors.New("team does not exist")
	}

	user, err := lm.GetUserByCn(userName)
	if err != nil {
		return false, errors.New("user does not exist")
	}

	for _, member := range team.Members {
		if member.UserId == user.UserId {
			return true, nil
		}
	}

	return false, nil
}

func (lm *LDAPManagement) IsTeamLeader(userName string, groupName string) (bool, error) {
	team, err := lm.GetGroupByGroupName(groupName)
	fmt.Println("team", team)
	if err != nil {
		return false, errors.New("team does not exist")
	}

	user, err := lm.GetUserByCn(userName)
	if err != nil {
		return false, errors.New("user does not exist")
	}

	if user.UserId == team.Leader.UserId {
		return true, nil
	}

	return false, nil
}

func (lm *LDAPManagement) AddUserToGroup(userName string, groupName string) ([]*User, error) {

	isMember, err := lm.IsTeamMember(userName, groupName)
	if err != nil {
		fmt.Println("error here")
		return nil, err
	}

	if isMember {
		// group, err := lm.GetGroupByGroupName(adminUser, adminPasswd, groupName)
		// fmt.Println("group", group)
		// if err != nil {
		// 	return nil, err
		// }
		// return group.Members, nil
		return nil, errors.New("the user is already a member of the team")
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	modify := ldap.NewModifyRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN),
		[]ldap.Control{},
	)
	modify.Add("member", []string{fmt.Sprintf("cn=%s,%s", userName, baseDN)})

	err = lm.ldapConn.Modify(modify)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to add user to group")
	}

	group, _ := lm.GetGroupByGroupName(groupName)
	fmt.Println(group)

	return group.Members, nil
}

func (lm *LDAPManagement) DeleteUserFromTeam(userName string, groupName string) ([]*User, error) {

	isMember, err := lm.IsTeamMember(userName, groupName)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("the user is not a member of the team")
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	modify := ldap.NewModifyRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN),
		[]ldap.Control{},
	)
	modify.Delete("member", []string{fmt.Sprintf("cn=%s,%s", userName, baseDN)})

	err = lm.ldapConn.Modify(modify)
	if err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return nil, err
	}

	group, _ := lm.GetGroupByGroupName(groupName)

	return group.Members, nil
}

func (lm *LDAPManagement) UpdateGroupLeader(newLeaderName string, groupName string) error {
	_, err := lm.GetUserByCn(newLeaderName)
	if err != nil {
		return err
	}

	_, err = lm.GetGroupByGroupName(groupName)
	if err != nil {
		return err
	}

	lm.connectWithoutTLS()
	defer lm.ldapConn.Close()
	lm.bind(lm.adminUser, lm.adminPasswd)

	baseDN := config.GetDC()
	modify := ldap.NewModifyRequest(
		fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, baseDN),
		[]ldap.Control{},
	)
	modify.Replace("o", []string{newLeaderName})
	if err = lm.ldapConn.Modify(modify); err != nil {
		log.Println(fmt.Errorf("failed to query LDAP: %w", err))
		return errors.New("failed to update new leader!")
	}

	return nil
}
