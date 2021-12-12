package controller

import (
	b64 "encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/service/auth_service"
)

type authApiHandler struct {
	loginUseCase *auth_service.LoginUseCase
}

func RegisterAuthApi(
	rg *gin.RouterGroup,
	userRepo repository.UserRepository,
	logger *logrus.Logger,
) {
	loginUseCase := auth_service.NewLoginUseCase(userRepo)

	h := authApiHandler{
		loginUseCase: loginUseCase,
	}

	auth := rg.Group("/auth")
	auth.POST("/login", h.login)
}

func (h authApiHandler) login(c *gin.Context) {
	basicAuthStr := c.GetHeader("Authorization")
	strs := strings.Split(basicAuthStr, " ")
	if strs[0] != "Basic" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid authorization header",
		})
		return
	}

	decodedStr, err := b64.StdEncoding.DecodeString(strs[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	account := strings.Split(string(decodedStr), ":")[0]
	password := strings.Split(string(decodedStr), ":")[1]
	input := auth_service.LoginUseCaseInput{
		Account:  account,
		Password: password,
	}
	output := auth_service.LoginUseCaseOutput{}

	accessToken, err := h.loginUseCase.Execute(input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errors.New("incorrect account or password").Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", accessToken, 86400, "/", "localhost", false, true)
	c.SetCookie("has_token", "1", 86400, "/", "localhost", false, false)
	c.String(http.StatusOK, "Login succeeded")

	// c.String(http.StatusOK, accessToken)
}
