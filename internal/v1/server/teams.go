package server

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"

	// "os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/controller/middleware"

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

type RemoveMemberRequest struct {
	GroupName string
	Leader    string
	Username  string
}

func teams(rg *gin.RouterGroup) {
	team := rg

	team.POST("/team/create", middleware.AuthMiddleware(), middleware.AdminMiddleware(), createTeam)
	team.POST("/team/delete", middleware.AuthMiddleware(), middleware.AdminMiddleware(), deleteTeam)
	team.GET("/teams", getAllTeams)
	team.GET("/team/:teamId", getTeam)
	team.GET("/team/:teamId/members", getTeamMember)
	team.POST("/team/:teamId/member", addMember)
	team.DELETE("/team/:teamId/member", middleware.AuthMiddleware(), removeMember)
	team.POST("/team/:teamId/leader/handover", middleware.AuthMiddleware(), handoverLeader)

	team.POST("/team/get/belonging-teams", getBelongingTeams)
}

func createTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()

	type createTeamReq struct {
		TeamName string `json:"teamName" binding:"required"`
		Leader   string `json:"teamLeader" binding:"required"`
	}

	// check request
	req := &createTeamReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTeamId := uuid.New().String()
	result, err := accountManagement.CreateGroup(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.TeamName,
		req.Leader,
		newTeamId,
	)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to create team. %s", err.Error()),
		})
		return
	}

	leaderID, err := accountManagement.GetUUIDByUsername(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.Leader,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get leader. %s", err.Error()),
		})
		return
	}

	database.InsertRole(leaderID, newTeamId, 1)

	c.JSON(200, gin.H{"result": result})
}

func getAllTeams(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	groupList, err := accountManagement.GetAllGroupsInDetail(config.GetAdminUser(), config.GetAdminPassword())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groupList)
}

func getTeam(c *gin.Context) {
	teamId, ok := c.Params.Get("teamId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "teamId missed",
		})
		return
	}

	accountManagement := account.NewLDAPManagement()
	team, err := accountManagement.GetGroupInDetail(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, team)
}

func getTeamMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()

	teamId, _ := c.Params.Get("teamId")

	memberList, err := accountManagement.GetGroupMembersDetail(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
	)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberList)
}

func getBelongingTeams(c *gin.Context) {
	type GetBelongingTeamsReq struct {
		Username string `json:"username" binding:"required"`
	}

	req := &GetBelongingTeamsReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accountManagement := account.NewLDAPManagement()

	teams, err := accountManagement.GetUserBelongingTeams(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.Username,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func deleteTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()

	type deleteTeamReq struct {
		TeamName string `json:"teamName" binding:"required"`
	}
	req := &deleteTeamReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	teamId, err := accountManagement.GetUUIDByUsername(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.TeamName,
	)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DeleteTeam(teamId)

	err = accountManagement.DeleteGroup(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.TeamName,
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"result": fmt.Sprintf("team deleted: %s", teamId),
	})
}

func addMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()

	type AddMemberReq struct {
		UserId string `json:"userId" binding:"required"`
	}

	req := AddMemberReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamId, _ := c.Params.Get("teamId")

	_, err := accountManagement.GetUserByID(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		req.UserId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = accountManagement.GetGroupInDetail(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isMember, err := accountManagement.IsMember(teamId, req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if isMember {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": errors.New("the user is already a member of the team"),
		})
		return
	}

	memberList, err := accountManagement.AddMemberToGroup(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
		req.UserId,
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.InsertRole(req.UserId, teamId, 0)

	c.JSON(http.StatusOK, gin.H{
		"members": memberList,
	})
}

func removeMember(c *gin.Context) {
	type RemoveMemberReq struct {
		UserId string `json:"userId" binding:"required"`
	}

	req := &RemoveMemberReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid := c.GetString("uid")
	teamId, _ := c.Params.Get("teamId")

	// only team lead can remove member
	role, err := database.GetRole(uid, teamId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if role != 1 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "only team lead can remove member from team",
		})
		return
	}

	database.DeleteRole(req.UserId, teamId)

	accountManagement := account.NewLDAPManagement()

	members, err := accountManagement.RemoveMemberFromGroup(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
		req.UserId,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"members": members,
	})
}

func handoverLeader(c *gin.Context) {
	type HandOverLeaderReq struct {
		NewLeaderId string `json:"newLeaderId" binding:"required"`
	}
	req := &HandOverLeaderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// user that perform the operation
	userId := c.GetString("uid")
	fmt.Println("userId:", userId)

	// target team id
	teamId, _ := c.Params.Get("teamId")
	fmt.Println("teamId:", teamId)

	// only team lead can remove member
	role, err := database.GetRole(userId, teamId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if role != 1 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "only team lead can transfer the ownership",
		})
		return
	}

	accountManagement := account.NewLDAPManagement()

	err = accountManagement.UpdateTeamLeader(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		teamId,
		req.NewLeaderId,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.UpdateLeader(userId, req.NewLeaderId, teamId)

	c.Status(http.StatusNoContent)
}
