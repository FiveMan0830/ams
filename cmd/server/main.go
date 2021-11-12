package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	router.Use(cors.Default())
	registerRoutes()
	router.Run(":8080")
}

func registerRoutes() {
	v1 := router.Group("/")
	login(v1)
	teams(v1)
	auth(v1)
	role(v1)
	mySQL(v1)
}
