package main

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/controller"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

func main() {
	engine := pkg.NewGinEngine()

	db := pkg.NewMysqlClient()

	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)

	v2 := engine.Group("/api/v2")

	logger := pkg.NewLoggerClient()
	controller.RegisterUserApi(v2, userRepo, logger)
	controller.RegisterTeamApi(v2, teamRepo, userRepo, logger)
	controller.RegisterAuthApi(v2, userRepo, logger)

	engine.Run(":10000")
}
