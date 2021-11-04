package server

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/authorization"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
// )

// func auth(rg *gin.RouterGroup) {
// 	auth := rg
// 	auth.GET("/authorize", authorize)
// }

// func authorize(c *gin.Context) {
// 	verifyTokenService := authorization.NewVerifyJWTService(config.NewAuthConfig())
// 	userID, err := verifyTokenService.ExtractAccessTokenToUserID(c.Request)

// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(401, "Unauthorized")
// 		return
// 	}

// 	var user *account.User
// 	accountManagement := account.NewLDAPManagement()
// 	user, err = accountManagement.GetUserByID(config.GetAdminUser(), config.GetAdminPassword(), userID)

// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(404, "User not found")
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"userId":      user.UserId,
// 		"username":    user.Username,
// 		"displayName": user.DisplayName,
// 		"email":       user.Email,
// 	})
// }
