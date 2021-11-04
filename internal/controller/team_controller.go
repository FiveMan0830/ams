package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/service/team_service"
)

type teamApiHandler struct {
	addTeamUseCase team_service.AddTeamUseCase
	getTeamUseCase team_service.GetTeamUseCase
	logger         *logrus.Logger
}

func RegisterTeamApi(
	rg *gin.RouterGroup,
	teamRepo repository.TeamRepository,
	logger *logrus.Logger,
) {
	addTeamUseCase := team_service.NewAddTeamUseCase(teamRepo)
	getTeamUseCase := team_service.NewGetTeamUseCase(teamRepo)

	h := teamApiHandler{
		addTeamUseCase: addTeamUseCase,
		getTeamUseCase: getTeamUseCase,
		logger:         logger,
	}

	team := rg.Group("/team")
	team.GET("/:id/get", h.getTeam)
	team.POST("/add", h.addTeam)

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
		return
	}

	c.Status(http.StatusCreated)
}
