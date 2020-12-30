package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type LoginRequest struct {
	Username string
	Password string
}

func main() {
	router := gin.Default()
	accountManagement := account.NewLDAPManagement()
	router.Static("/", "./web")

	router.POST("/login", func(c *gin.Context) {
		reqbody := &LoginRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		info, err := accountManagement.Login(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username, reqbody.Password)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, info)
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}