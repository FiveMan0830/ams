package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
)

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

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, lm.BaseDN), []ldap.Control{})

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
	addReq.Attribute("member", []string{fmt.Sprintf("cn=%s,%s", username, lm.BaseDN)})
	addReq.Attribute("uid", []string{teamId})

	if err := conn.Add(addReq); err != nil {
		log.Println("error adding group:", addReq, err)
		return "", err
	}

	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(groupName))
	request := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"cn"},
		[]ldap.Control{},
	)
	result, err := conn.Search(request)

	if err != nil {
		log.Println("error searching group:", request, err)
		return "", err
	}

	group := strings.Join(result.Entries[0].GetAttributeValues("cn"), "")
	return group, nil
}

func (lm *LDAPManagement) GetGroupMembersDetail(adminUser, adminPasswd, teamId string) ([]*MemberRole, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	filter := fmt.Sprintf("(&(objectClass=groupOfNames)(uid=%s))", teamId)
	searchRequest := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
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
		userName := strings.Replace(memberDn, fmt.Sprintf(",%s", lm.BaseDN), "", -1)

		member, _ := lm.GetUserByUsername(adminUser, adminPasswd, userName)
		role, _ := database.GetRole(member.UserID, teamId)
		memberRole := &MemberRole{member, role.String()}

		memberList = append(memberList, memberRole)
	}

	return memberList, nil
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

	// check if the user is the member
	isMember, err := lm.IsMember(teamId, userId)
	if err != nil {
		return nil, err
	}
	if isMember {
		log.Printf("User %s is already a member of the group %s\n", user.Username, team.Name)
		return nil, errors.New("user already member of the group")
	}

	// query
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", team.Name, lm.BaseDN), []ldap.Control{})
	modify.Add("member", []string{fmt.Sprintf("cn=%s,%s", user.Username, lm.BaseDN)})
	err = conn.Modify(modify)

	if err != nil {
		return nil, errors.New("failed to add user to group")
	}

	members, _ := lm.GetGroupMembersDetail(adminUser, adminPasswd, teamId)

	return members, nil
}

// RemoveMemberFromGroup is a function to remove a user from a group
func (lm *LDAPManagement) RemoveMemberFromGroup(adminUser, adminPasswd, teamId, userId string) ([]*MemberRole, error) {
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
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", team.Name, lm.BaseDN), []ldap.Control{})
	modify.Delete("member", []string{fmt.Sprintf("cn=%s,%s", user.Username, lm.BaseDN)})
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

// get leader of the team
func (lm *LDAPManagement) GetTeamLeader(adminUser, adminPasswd, teamId string) (*User, error) {
	conn, _ := lm.getConnectionWithoutTLS()
	lm.bindAuth(conn, adminUser, adminPasswd)
	defer conn.Close()

	filter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(teamId))
	request := ldap.NewSearchRequest(
		lm.BaseDN,
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

	filter := fmt.Sprintf("(&(objectClass=groupOfNames)(uid=%s))", teamId)
	searchRequest := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
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

	leader, err := lm.GetTeamLeader(adminUser, adminPasswd, teamId)
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

	filter := fmt.Sprintf("(objectClass=%s)", ldap.EscapeFilter(ObjectCategoryGroup))
	searchRequest := ldap.NewSearchRequest(
		lm.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
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

	d := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=OISGroup,%s", groupName, lm.BaseDN), nil)
	err := conn.Del(d)

	if err != nil {
		log.Println("Group entry could not be deleted :", d, err)
		return err
	}

	return nil
}
