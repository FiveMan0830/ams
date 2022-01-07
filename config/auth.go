package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type authConfig struct {
	tokenSecret string
	expiredDays int
}

func NewAuthConfig() AuthConfig {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	expiredDays, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_DAYS"))
	if err != nil {
		log.Fatalln("getting environment TOKEN_EXPIRED_DAYS failed: " + err.Error())
	}

	auth := &authConfig{
		tokenSecret: tokenSecret,
		expiredDays: expiredDays,
	}

	return auth
}

func (ac *authConfig) TokenSecret() string {
	return ac.tokenSecret
}

func (ac *authConfig) TokenExpiredTime() int64 {
	return time.Now().Add(time.Hour * 24 * time.Duration(ac.expiredDays)).Unix()
}

type AuthConfig interface {
	TokenSecret() string
	TokenExpiredTime() int64
}
