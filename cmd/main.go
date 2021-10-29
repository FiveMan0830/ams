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

	v2 := router.Group("/v2")

	controller.RegisterUserApi(v2, userRepo)
	controller.RegisterTeamApi(v2)

	router.Run(":10000")
}
