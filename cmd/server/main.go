package server

import (
	"github.com/gin-gonic/gin"
	cors "github.com/gin-contrib/cors"
)

var (
	router = gin.Default()
)

func Run() {
	router.Use(cors.Default())
	getRoutes()
	router.Run()
}

func getRoutes() {
	v1 := router.Group("/")
	login(v1)
	teams(v1)
}
