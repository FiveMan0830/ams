package pkg

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000", "http://localhost:30000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "withCredentials"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8000" || origin == "http://localhost:30000" || origin == "http://localhost:3000"
		},
	}))

	engine.GET("/api/v2/cookie", func(c *gin.Context) {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("coo", "kie", 3600, "/", "localhost", false, false)
	})

	return engine
}
