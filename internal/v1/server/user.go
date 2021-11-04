package server

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
// )

// type userApiHandler struct {
// 	accountManager account.Management
// 	adminUser      string
// 	adminPassword  string
// }

// func registerUserApi(
// 	rg *gin.RouterGroup,
// 	accountManager account.Management,
// ) {
// 	h := userApiHandler{
// 		accountManager: accountManager,
// 		adminUser:      config.GetAdminUser(),
// 		adminPassword:  config.GetAdminPassword(),
// 	}

// 	user := rg.Group("/user/:id/")
// 	{
// 		user.GET("/", h.getUser)
// 		user.GET("/team", h.getUserTeam)
// 	}
// }

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
