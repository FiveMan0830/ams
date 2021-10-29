package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/service/user_service"
)

type userApiHandler struct {
	addUserUseCase user_service.AddUserUseCase
}

func RegisterUserApi(
	rg *gin.RouterGroup,
	userRepo *repository.UserRepository,
) {
	addUserUseCase := user_service.NewAddUserUseCase(userRepo)

	h := userApiHandler{
		addUserUseCase: addUserUseCase,
	}

	user := rg.Group("/user")
	user.GET("/get", h.getUser)
	user.POST("/add", h.addUser)
	user.PUT("/edit", h.editUser)
	user.DELETE("/remove", h.removeUser)
}

func (h userApiHandler) getUser(c *gin.Context) {
	input := user_service.AddUserUseCaseInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	h.addUserUseCase.Execute(input)
}

func (h userApiHandler) removeUser(c *gin.Context) {

}

func (h userApiHandler) editUser(c *gin.Context) {

}

func (h userApiHandler) addUser(c *gin.Context) {

}
