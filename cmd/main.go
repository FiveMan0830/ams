package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/controller"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	db := pkg.NewMysqlClient()

	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)

	v2 := router.Group("/api/v2")

	logger := pkg.NewLoggerClient()
	controller.RegisterUserApi(v2, userRepo, logger)
	controller.RegisterTeamApi(v2, teamRepo, userRepo, logger)
	controller.RegisterAuthApi(v2, userRepo, logger)

	router.Run(":10000")
}
