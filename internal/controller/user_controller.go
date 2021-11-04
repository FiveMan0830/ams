package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/service/user_service"
)

type userApiHandler struct {
	addUserUseCase    user_service.AddUserUseCase
	getUserUseCase    user_service.GetUserUseCase
	removeUserUseCase user_service.RemoveUserUseCase
	updateUserUseCase user_service.UpdateUserUseCase
	logger            *logrus.Logger
}

func RegisterUserApi(
	rg *gin.RouterGroup,
	userRepo repository.UserRepository,
	logger *logrus.Logger,
) {
	addUserUseCase := user_service.NewAddUserUseCase(userRepo)
	getUserUseCase := user_service.NewGetUserUseCase(userRepo)
	removeUserUseCase := user_service.NewRemoveUserUseCase(userRepo)
	updateUserUseCase := user_service.NewUpdateUserUseCase(userRepo)

	h := userApiHandler{
		addUserUseCase:    addUserUseCase,
		getUserUseCase:    getUserUseCase,
		removeUserUseCase: removeUserUseCase,
		updateUserUseCase: updateUserUseCase,
		logger:            logger,
	}

	user := rg.Group("/user")
	user.GET("/:id/get", h.getUser)
	user.POST("/add", h.addUser)
	user.PUT("/:id/edit", h.editUser)
	user.DELETE("/:id/remove", h.removeUser)
}

func (h userApiHandler) getUser(c *gin.Context) {
	id := c.Param("id")

	input := user_service.GetUserUseCaseInput{Id: id}
	output := user_service.GetUserUseCaseOutput{}
	if err := h.getUserUseCase.Execute(input, &output); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h userApiHandler) removeUser(c *gin.Context) {
	id := c.Param("id")
	input := user_service.RemoveUserUseCaseInput{}
	input.Id = id

	if err := h.removeUserUseCase.Execute(input); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h userApiHandler) editUser(c *gin.Context) {
	input := user_service.UpdateUserUseCaseInput{}
	c.ShouldBindJSON(&input)
	input.Id = c.Param("id")

	if err := h.updateUserUseCase.Execute(input); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h userApiHandler) addUser(c *gin.Context) {
	input := user_service.AddUserUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.addUserUseCase.Execute(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
