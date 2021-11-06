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
	addMembersUseCase     *team_service.AddMembersUseCase
	removeMembersUseCase  *team_service.RemoveMembersUseCase
	addSubteamUseCase     *team_service.AddSubteamUseCase
	removeSubteamUseCase  *team_service.RemoveSubteamUseCase
	updateUserRoleUseCase *team_service.UpdateUserRoleUseCase
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
	addMembersUseCase := team_service.NewAddMembersUseCase(teamRepo)
	removeMembersUseCase := team_service.NewRemoveMembersUseCase(teamRepo)
	addSubteamUseCase := team_service.NewAddSubteamUseCase(teamRepo)
	removeSubteamUseCase := team_service.NewRemoveSubteamUseCase(teamRepo)
	updateUserRoleUseCase := team_service.NewUpdateUserRoleUseCase(teamRepo, userRepo)

	h := teamApiHandler{
		addTeamUseCase:        addTeamUseCase,
		getTeamUseCase:        getTeamUseCase,
		addMembersUseCase:     addMembersUseCase,
		removeMembersUseCase:  removeMembersUseCase,
		addSubteamUseCase:     addSubteamUseCase,
		removeSubteamUseCase:  removeSubteamUseCase,
		updateUserRoleUseCase: updateUserRoleUseCase,
		logger:                logger,
	}

	team := rg.Group("/team")
	team.POST("/assign/users", h.addMembers)
	team.DELETE("/unassign/users", h.removeMembers)

	team.POST("/assign/teams", h.addSubteams)
	team.DELETE("/unassign/teams", h.removeSubteams)

	team.GET("/:id", h.getTeam)
	team.POST("/", h.addTeam)

	team.PUT("/role/edit", h.updateUserRole)
}

func (h teamApiHandler) getTeam(c *gin.Context) {
	id := c.Param("id")

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
	input := team_service.AddMembersUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.addMembersUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) removeMembers(c *gin.Context) {
	teamId := c.Param("id")

	input := team_service.RemoveMembersUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	input.TeamId = teamId

	if err := h.removeMembersUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) addSubteams(c *gin.Context) {
	input := team_service.AddSubteamUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.addSubteamUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h teamApiHandler) removeSubteams(c *gin.Context) {
	input := team_service.RemoveSubteamUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.removeSubteamUseCase.Execute(input); err != nil {
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
