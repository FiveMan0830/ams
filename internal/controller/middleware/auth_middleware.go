package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

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

		c.Set("uid", token.Claims.(jwt.MapClaims)["uid"])
		c.Set("tokenExp", token.Claims.(jwt.MapClaims)["exp"])

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("uid")

		accountManager := account.NewLDAPManagement(account.LDAPManagerConfig{BaseDN: config.GetDC()})
		user, err := accountManager.GetUserByID(
			config.GetAdminUser(),
			config.GetAdminPassword(),
			userId,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		if user.Username != "ams_admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "only admin can perform this action",
			})
			return
		}

		c.Next()
	}
}
