package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	rg.POST("user", middleware.AuthMiddleware(), middleware.AdminMiddleware(), h.createUser)
	rg.DELETE("/user/:userId", middleware.AuthMiddleware(), middleware.AdminMiddleware(), h.deleteUser)
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

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *userApiHandler) createUser(c *gin.Context) {
	type CreateUserReq struct {
		Username  string `json:"username" binding:"required"`
		Givenname string `json:"givenname" binding:"required"`
		Surname   string `json:"surname" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

	req := &CreateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accountManager := account.NewLDAPManagement()
	user, err := accountManager.CreateUser(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		uuid.New().String(),
		req.Username,
		req.Givenname,
		req.Surname,
		req.Password,
		req.Email,
	)
	if err != nil {
		if err.Error() == "user already exist" {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userApiHandler) deleteUser(c *gin.Context) {
	userId, ok := c.Params.Get("userId")

	if !ok || userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	accountManager := account.NewLDAPManagement()
	err := accountManager.DeleteUserByUserId(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		userId,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
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
