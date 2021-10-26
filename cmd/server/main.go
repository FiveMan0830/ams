package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func Run() {
	router := gin.Default()
	router.Use(cors.Default())

	accountManager := account.NewLDAPManagement()

	v1 := router.Group("/")
	login(v1)
	registerTeamApi(v1, accountManager)
	teams(v1)
	registerUserApi(v1, accountManager)
	auth(v1)
	role(v1)
	mySQL(v1)

	router.Run(":18080")
}
