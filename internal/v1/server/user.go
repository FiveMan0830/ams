package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/controller/middleware"
)

type userApiHandler struct {
	accountManager account.Management
	adminUser      string
	adminPassword  string
}

func registerUserApi(
	rg *gin.RouterGroup,
	accountManager account.Management,
) {
	h := userApiHandler{
		accountManager: accountManager,
		adminUser:      config.GetAdminUser(),
		adminPassword:  config.GetAdminPassword(),
	}

	// rg.GET("/user/:id", h.getUser)
	rg.GET("/profile", middleware.AuthMiddleware(), h.getUserProfile)
	rg.GET("/users", h.getAllUsers)
}

// func (h *userApiHandler) getUser(c *gin.Context) {
// 	userId := c.Param("id")
// 	user, err := h.accountManager.GetUserById(userId)

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

func (h *userApiHandler) getUserProfile(c *gin.Context) {
	userId := c.GetString("uid")

	user, err := h.accountManager.GetUserByID(h.adminUser, h.adminPassword, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to get user profile. userId: %s", userId).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userApiHandler) getAllUsers(c *gin.Context) {
	users, err := h.accountManager.GetAllUsers(h.adminUser, h.adminPassword)
	fmt.Println(users)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// func (h *userApiHandler) getUserTeam(c *gin.Context) {
// 	userId := c.Param("id")
// 	user, err := h.accountManager.GetGroupsByUserId(userId)

// 	if err != nil {
// 		fmt.Println(err)
// 		c.Status(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }
