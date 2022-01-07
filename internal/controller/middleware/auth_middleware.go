package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		fmt.Println(bearerToken)

		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "auth header missed",
			})
			return
		}

		strArr := strings.Split(bearerToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid access token",
			})
			return
		}
		tokenString := strArr[1]

		jwtClient := pkg.NewJWTClient(config.NewAuthConfig())
		token, err := jwtClient.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		c.Set("token", token)

		c.Next()
	}
}
