package main

import (
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/controller"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/v1/server"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

func main() {
	engine := pkg.NewGinEngine()

	// dependencies for api v1
	am := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})

	// dependencies for api v2
	db := pkg.NewMysqlClient()
	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)

	// register api v1
	v1 := engine.Group("/api/v1")
	server.RegisterV1Api(v1, am)

	engine.GET("/health-check", func(c *gin.Context) {
		c.String(200, "the server is healthy\n")
	})

	// register api v2
	v2 := engine.Group("/api/v2")
	logger := pkg.NewLoggerClient()
	controller.RegisterUserApi(v2, userRepo, logger)
	controller.RegisterTeamApi(v2, teamRepo, userRepo, logger)
	controller.RegisterAuthApi(v2, userRepo, logger)

	engine.Run(":8080")
}
