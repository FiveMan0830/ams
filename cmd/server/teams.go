package server

import (
	"io/ioutil"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type GetGroupRequest struct {
	GroupName     string
	SelfUsername  string
	InputUsername string
}

type GetGroupsListRequest struct {
	GroupList []string
}

type GetUsersRequest struct {
	Username string
}

type AddMemberRequest struct {
	GroupName string
	Username  string
}

type RemoveMemberRequest struct {
	GroupName string
	Leader    string
	Username  string
}

func teams(rg *gin.RouterGroup) {
	team := rg

	team.POST("/team/create", createTeam)
	team.GET("/team", getTeam)
	team.POST("/team/get/members", getTeamMember)
	team.POST("/team/get/leader", getTeamLeader)
	team.POST("/team/isleader", isLeader)
	team.POST("/team/get/belonging-teams", getBelongingTeams)
	team.POST("/team/get/uuid/user", getUUIDOfUser)
	team.POST("/team/get/uuid/team", getUUIDOfTeam)
	team.POST("/team/get/name", getName)
	team.POST("/team/delete", deleteTeam)
	team.POST("/team/add/member", addMember)
	team.POST("/team/remove/member", removeMember)
	team.POST("/team/leader/handover", handoverLeader)
	team.POST("/team/get/member/name", getTeamMemberUsernameAndDisplayname)
	team.GET("/all/username", getAllUsername)
	team.POST("/team/member/role", getRoleOfTeamMembers)

	// Richard requested API
	team.POST("/get/name", getName)
}

func createTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)
	teamID := uuid.New().String()
	info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.SelfUsername, teamID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, info)
}

func getTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	GroupList, err := accountManagement.GetGroups(config.GetAdminUser(), config.GetAdminPassword())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, GroupList)
}

func getTeamMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	c.Bind(reqbody)
	memberList, err := accountManagement.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, memberList)
}

func getTeamLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	c.Bind(reqbody)
	leader, err := accountManagement.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, leader)
}

func isLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)

	result := accountManagement.IsLeader(reqbody.GroupName, reqbody.SelfUsername)

	c.JSON(200, result)
}

func getBelongingTeams(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	c.Bind(reqbody)
	GroupList, err := accountManagement.SearchUserMemberOf(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, GroupList)
}

func getUUIDOfUser(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	uuid, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, uuid)
}

func getUUIDOfTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	uuid, err := accountManagement.SearchGroupUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, uuid)
}

func getName(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)

	name, err := accountManagement.SearchNameByUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, name)
}

func deleteTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)
	err := accountManagement.DeleteGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	if err != nil {
		c.JSON(500, err)
		return
	}
}

func addMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &AddMemberRequest{}
	c.Bind(reqbody)

	if !accountManagement.IsMember(reqbody.GroupName, reqbody.Username) {
		memberList, err := accountManagement.AddMemberToGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, memberList)
	} else {
		c.JSON(403, "User is not member of the team!")
		return
	}
}

func removeMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &RemoveMemberRequest{}
	c.Bind(reqbody)

	if accountManagement.IsLeader(reqbody.GroupName, reqbody.Leader) {
		memberList, err := accountManagement.RemoveMemberFromGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, memberList)
	} else {
		c.JSON(403, "User is not leader of the team!")
		return
	}
}

func handoverLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)

	if accountManagement.IsLeader(reqbody.GroupName, reqbody.SelfUsername) || accountManagement.IsProfessor(reqbody.SelfUsername) {
		err := accountManagement.UpdateGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.InputUsername)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}

	} else {
		c.JSON(403, "User is not professor or leader of the team!")
		return
	}

}

func getTeamMemberUsernameAndDisplayname(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)

	memberList, err := accountManagement.GetGroupMembersUsernameAndDisplayname(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, memberList)
}

func getAllUsername(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	userList, err := accountManagement.SearchAllUser(config.GetAdminUser(), config.GetAdminPassword())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, userList)
}

// input is unitID of team 
func getRoleOfTeamMembers(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	c.Bind(reqbody)

	teamName, err := accountManagement.SearchNameByUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))
	memberList, err := accountManagement.GetGroupMembersRole(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, memberList)
}