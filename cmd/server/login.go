package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/authorization"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

var (
	authConfig config.AuthConfig
)

func init() {
	authConfig = config.NewAuthConfig()
}

type LoginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	accessToken string
}

func login(rg *gin.RouterGroup) {
	login := rg
	login.Static("/login", "./web")
	login.POST("/login", loginUser)
}

func loginUser(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &LoginRequest{}
	c.Bind(reqbody)
	log.Println(reqbody)
	info, err := accountManagement.Login(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username, reqbody.Password)

	if err != nil {
		log.Println("Login failed: authorization failed: ", err.Error())
		c.JSON(401, err)
		return
	}

	userID := info.GetAttributeValue("uid")

	jwtGenerator := authorization.NewJWTGenerator(authConfig)
	token, err := jwtGenerator.CreateToken(userID)

	log.Println("access token: " + token)

	if err != nil {
		log.Println("login failed: generate token failed -> ", err.Error())
		c.JSON(500, "login failed")
	}

	c.JSON(200, gin.H{"accessToken": token})
}
