package server

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// var (
// 	authConfig config.AuthConfig
// )

// func init() {
// 	authConfig = config.NewAuthConfig()
// }

// type LoginRequest struct {
// 	Username string
// 	Password string
// }

// type loginResponse struct {
// 	accessToken string
// }

func login(rg *gin.RouterGroup) {
	login := rg

	// login.Static("/login", "./web")
	login.POST("/auth/login", loginUser)
}

// func loginUserByAccessToken(c *gin.Context) {
// 	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
// 	reqbody := &LoginRequest{}
// 	c.Bind(reqbody)
// 	log.Println(reqbody)
// 	info, err := accountManagement.Login(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username, reqbody.Password)

// 	if err != nil {
// 		log.Println("Login failed: authorization failed: ", err.Error())
// 		c.JSON(401, err)
// 		return
// 	}

// 	userID := info.GetAttributeValue("uid")

// 	jwtGenerator := authorization.NewJWTGenerator(authConfig)
// 	token, err := jwtGenerator.CreateToken(userID)

// 	log.Println("access token: " + token)

// 	if err != nil {
// 		log.Println("login failed: generate token failed -> ", err.Error())
// 		c.JSON(500, "login failed")
// 	}

// 	c.JSON(200, gin.H{"accessToken": token})
// }

func loginUser(c *gin.Context) {
	basicAuthStr := c.GetHeader("Authorization")
	strs := strings.Split(basicAuthStr, " ")
	if strs[0] != "Basic" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid authorization header",
		})
		return
	}

	decodedStr, err := b64.StdEncoding.DecodeString(strs[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := strings.Split(string(decodedStr), ":")[0]
	password := strings.Split(string(decodedStr), ":")[1]

	accountManagement := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	accessToken, err := accountManagement.Login(
		config.GetAdminUser(),
		config.GetAdminPassword(),
		username,
		password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
