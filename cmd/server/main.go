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
	getRoutes()
	router.Run()
}

func getRoutes() {
	v1 := router.Group("/")
	login(v1)
	teams(v1)
	auth(v1)
}
