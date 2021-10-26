package server

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"

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

type teamApiHandler struct {
	accountManager account.Management
	adminUser      string
	adminPassword  string
}

func registerTeamApi(
	rg *gin.RouterGroup,
	accountManager account.Management,
) {
	h := teamApiHandler{
		accountManager: accountManager,
		adminUser:      config.GetAdminUser(),
		adminPassword:  config.GetAdminPassword(),
	}

	rg.GET("/teams", h.getAllTeams)
	rg.GET("/team/:id", h.getTeamById)

	rg.POST("/team/add", h.createTeam)
	rg.DELETE("/team/delete", h.deleteTeam)
	rg.POST("/team/member/add", h.addUserToTeam)

	// DANGERGOUS
	// Only the team leader can perform these operation.
	// should check user identity by access token in the future
	rg.DELETE("/team/member/delete", h.deleteUserFromTeam)
	rg.PUT("/team/leader", h.handoverLeader)

	// rg.GET("/team/:id/member", h.getTeamMembers)
}

func (h *teamApiHandler) getAllTeams(c *gin.Context) {
	teams, err := h.accountManager.GetAllGroups()

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teams": teams,
		"total": len(teams),
	})
}

func (h *teamApiHandler) getTeamById(c *gin.Context) {
	teamId := c.Param("id")

	team, err := h.accountManager.GetGroupById(teamId)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *teamApiHandler) createTeam(c *gin.Context) {
	type createTeamReq struct {
		Name       string `json:"name" binding:"required"`
		LeaderName string `json:"leaderName" binding:"required"`
	}

	// check req body
	req := createTeamReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check user exist
	user, err := h.accountManager.GetUserByCn(req.LeaderName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create team
	teamId := uuid.New().String()
	_, err = h.accountManager.CreateGroup(teamId, req.Name, req.LeaderName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user-role map to db
	database.InsertRole(user.UserId, teamId, 1)

	c.Status(http.StatusCreated)
}

func (h *teamApiHandler) deleteTeam(c *gin.Context) {
	type deleteTeamReq struct {
		Name string `json:"name" binding:"required"`
	}

	// check req body
	req := deleteTeamReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// get team and check if it exist
	team, err := h.accountManager.GetGroupByGroupName(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// delete team from mysql
	database.DeleteTeam(team.Id)

	// delete team from ldap
	_, err = h.accountManager.DeleteGroup(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *teamApiHandler) addUserToTeam(c *gin.Context) {
	type addUserToTeamReq struct {
		UserName string `json:"userName" binding:"required"`
		TeamName string `json:"teamName" binding:"required"`
	}

	// check req body
	req := addUserToTeamReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user to ldap db
	members, err := h.accountManager.AddUserToGroup(req.UserName, req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user-role map to mysql
	user, err := h.accountManager.GetUserByCn(req.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	team, err := h.accountManager.GetGroupByGroupName(req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.InsertRole(user.UserId, team.Id, 0)

	c.JSON(http.StatusOK, gin.H{
		"members": members,
	})
}

func (h *teamApiHandler) deleteUserFromTeam(c *gin.Context) {
	type deleteUserFromTeamReq struct {
		UserName string `json:"userName" binding:"required"`
		TeamName string `json:"teamName" binding:"required"`
	}

	// check req body
	req := deleteUserFromTeamReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// delete user from ldap db
	members, err := h.accountManager.DeleteUserFromTeam(req.UserName, req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// delete user-role map from mysql
	user, err := h.accountManager.GetUserByCn(req.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	team, err := h.accountManager.GetGroupByGroupName(req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DeleteRole(user.UserId, team.Id)

	c.JSON(http.StatusOK, gin.H{
		"members": members,
	})
}

func (h *teamApiHandler) handoverLeader(c *gin.Context) {
	type handoverLeaderReq struct {
		NewLeader string `json:"newLeader" binding:"required"`
		TeamName  string `json:"teamName" binding:"required"`
	}

	// check req body
	req := handoverLeaderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.accountManager.GetUserByCn(req.NewLeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	team, err := h.accountManager.GetGroupByGroupName(req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(user)
	fmt.Println(team)
	database.UpdateLeader(team.Leader.UserId, user.UserId, team.Id)

	err = h.accountManager.UpdateGroupLeader(req.NewLeader, req.TeamName)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// func (h *teamApiHandler) getTeamMembers(c *gin.Context) {
// 	teamId := c.Param("id")

// 	members, err := h.accountManager.GetGroupMembersByGroupId(
// 		h.adminUser,
// 		h.adminPassword,
// 		teamId,
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 		c.Status(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"members": members,
// 		"total":   len(members),
// 	})
// }

func teams(rg *gin.RouterGroup) {
	team := rg

	team.POST("/team/create", createTeam)                     ///
	team.GET("/team", getTeam)                                ///
	team.POST("/team/get/members", getTeamMember)             ///
	team.POST("/team/get/leader", getTeamLeader)              ///
	team.POST("/team/isleader", isLeader)                     ///
	team.POST("/team/get/belonging-teams", getBelongingTeams) ///
	team.POST("/team/get/uuid/user", getUUIDOfUser)
	team.POST("/team/get/uuid/team", getUUIDOfTeam)
	team.POST("/team/", getName)
	team.POST("/team/delete", deleteTeam)              ///
	team.POST("/team/add/member", addMember)           ///
	team.POST("/team/remove/member", removeMember)     ///
	team.POST("/team/leader/handover", handoverLeader) ///
	team.POST("/team/get/member/name", getTeamMemberUsernameAndDisplayname)
	team.GET("/all/username", getAllUsername)
	team.POST("/team/member/role", getRoleOfTeamMembers) ///

	// Richard requested API
	team.POST("/get/name", getName)
}

func createTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)
	teamID := uuid.New().String()
	info, err := accountManagement.CreateGroupDepre(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.SelfUsername, teamID)
	leaderID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.SelfUsername)

	fmt.Println(reqbody.GroupName)
	fmt.Println(reqbody.SelfUsername)
	fmt.Println(teamID)
	fmt.Println(info)
	fmt.Println(leaderID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	database.InsertRole(leaderID, teamID, 1)

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
	teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))
	leaderID, err := database.GetTeamLeader(teamID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, leaderID)
}

func isLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)

	leaderID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.SelfUsername)

	if err != nil {
		c.JSON(500, err)
		return
	}

	result := accountManagement.IsLeader(reqbody.GroupName, leaderID)

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
	teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)
	database.DeleteTeam(teamID)

	err = accountManagement.DeleteGroupDepre(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	fmt.Println("delete r" + reqbody.GroupName)
	fmt.Println("delete ID " + teamID)
	if err != nil {
		c.JSON(500, err)
		return
	}

	fmt.Println(teamID)

}

func addMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &AddMemberRequest{}
	c.Bind(reqbody)

	if !accountManagement.IsMember(reqbody.GroupName, reqbody.Username) {
		memberList, err := accountManagement.AddMemberToGroupDepre(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		userID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)
		teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		database.InsertRole(userID, teamID, 0)

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

	userID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Leader)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	targerUserID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	isLeader, err := database.GetRole(userID, teamID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if isLeader == 1 {
		memberList, err := accountManagement.RemoveMemberFromGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		database.DeleteRole(targerUserID, teamID)

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

	userID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.SelfUsername)
	teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)
	isLeader, err := database.GetRole(userID, teamID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if isLeader == 1 || isLeader == 2 {
		err := accountManagement.UpdateGroupLeaderDepre(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.InputUsername)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		oldLeaderID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.SelfUsername)
		newLeaderID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.InputUsername)
		teamID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		database.UpdateLeader(oldLeaderID, newLeaderID, teamID)

		c.JSON(200, "")
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

	fmt.Println("Team name: ", teamName)

	memberList, err := accountManagement.GetGroupMembersRole(config.GetAdminUser(), config.GetAdminPassword(), teamName)

	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, memberList)
}
