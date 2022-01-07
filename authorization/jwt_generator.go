package authorization

import (
	"github.com/dgrijalva/jwt-go"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type TokenGenerator interface {
	CreateToken(userId string) (string, error)
}

type jwtGenerator struct {
	config config.AuthConfig
}

func NewJWTGenerator(config config.AuthConfig) TokenGenerator {
	return &jwtGenerator{
		config: config,
	}
}

func (jg *jwtGenerator) CreateToken(userID string) (string, error) {
	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = jg.config.TokenExpiredTime()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jg.config.TokenSecret()))
	if err != nil {
		return "", err
	}
	return token, nil
}
