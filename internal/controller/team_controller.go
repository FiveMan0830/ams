package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/service/team_service"
)

type teamApiHandler struct {
	addTeamUseCase        *team_service.AddTeamUseCase
	getTeamUseCase        *team_service.GetTeamUseCase
	assignMembersUseCase  *team_service.AssignMembersUseCase
	expelMembersUseCase   *team_service.ExpelMembersUseCase
	assignSubteamUseCase  *team_service.AssignSubteamUseCase
	expelSubteamUseCase   *team_service.ExpelSubteamUseCase
	updateUserRoleUseCase *team_service.UpdateUserRoleUseCase
	belongToUseCase       *team_service.BelongToUseCase
	logger                *logrus.Logger
}

func RegisterTeamApi(
	rg *gin.RouterGroup,
	teamRepo repository.TeamRepository,
	userRepo repository.UserRepository,
	logger *logrus.Logger,
) {
	addTeamUseCase := team_service.NewAddTeamUseCase(teamRepo)
	getTeamUseCase := team_service.NewGetTeamUseCase(teamRepo)
	assignMembersUseCase := team_service.NewAssignMembersUseCase(teamRepo)
	expelMembersUseCase := team_service.NewExpelMembersUseCase(teamRepo)
	assignSubteamUseCase := team_service.NewAssignSubteamUseCase(teamRepo)
	expelSubteamUseCase := team_service.NewExpelSubteamUseCase(teamRepo)
	updateUserRoleUseCase := team_service.NewUpdateUserRoleUseCase(teamRepo, userRepo)
	belongToUseCase := team_service.NewBelongToUseCase(teamRepo)

	h := teamApiHandler{
		addTeamUseCase:        addTeamUseCase,
		getTeamUseCase:        getTeamUseCase,
		assignMembersUseCase:  assignMembersUseCase,
		expelMembersUseCase:   expelMembersUseCase,
		assignSubteamUseCase:  assignSubteamUseCase,
		expelSubteamUseCase:   expelSubteamUseCase,
		updateUserRoleUseCase: updateUserRoleUseCase,
		belongToUseCase:       belongToUseCase,
		logger:                logger,
	}

	team := rg.Group("/team")
	team.POST("/assign/users", h.addMembers)
	team.DELETE("/expel/users", h.removeMembers)

	team.POST("/assign/teams", h.addSubteams)
	team.DELETE("/expel/teams", h.removeSubteams)

	team.GET("/get", h.getTeam)
	team.POST("/add", h.addTeam)

	team.PUT("/role/edit", h.updateUserRole)
	team.GET("/belong_to", h.belongTo)
}

func (h teamApiHandler) getTeam(c *gin.Context) {
	id := c.Query("teamId")

	input := team_service.GetTeamUseCaseInput{}
	output := team_service.GetTeamUseCaseOuptut{}
	input.Id = id

	if err := h.getTeamUseCase.Execute(input, &output); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h teamApiHandler) addTeam(c *gin.Context) {
	input := team_service.AddTeamUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.addTeamUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h teamApiHandler) addMembers(c *gin.Context) {
	input := team_service.AssignMembersUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.assignMembersUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) removeMembers(c *gin.Context) {
	teamId := c.Param("id")

	input := team_service.ExpelMembersUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	input.TeamId = teamId

	if err := h.expelMembersUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) addSubteams(c *gin.Context) {
	input := team_service.AssignSubteamUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.assignSubteamUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) removeSubteams(c *gin.Context) {
	input := team_service.ExpelSubteamUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.expelSubteamUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) updateUserRole(c *gin.Context) {
	input := team_service.UpdateUserRoleUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.updateUserRoleUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) belongTo(c *gin.Context) {
	input := team_service.BelongToUseCaseInput{UserId: c.Query("userId")}
	output := team_service.BelongToUseCaseOutput{}

	if err := h.belongToUseCase.Execute(input, &output); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
