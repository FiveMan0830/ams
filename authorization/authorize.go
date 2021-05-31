package authorization

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type verifyTokenService struct {
   config config.AuthConfig
}

func NewVerifyTokenService(config config.AuthConfig) VerifyTokenService {
   return &verifyTokenService{config: config}
}

func (vts *verifyTokenService) extractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func (vts *verifyTokenService) verifyToken(r *http.Request) (*jwt.Token, error) {
  tokenString := vts.extractToken(r)
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
     return []byte(vts.config.GetTokenSecret()), nil
  })
  if err != nil {
     return nil, err
  }
  return token, nil
}

func (vts *verifyTokenService) TokenValid(r *http.Request) error {
  token, err := vts.verifyToken(r)
  if err != nil {
     return err
  }
  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
     return err
  }
  return nil
}

func (vts *verifyTokenService) ExtractAccessTokenToUserID(r *http.Request) (string, error) {
  token, err := vts.verifyToken(r)
  if err != nil {
     return "", err
  }
  claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
     userID := claims["user_id"].(string)
     return userID, nil
  }
  return "", err
}

type VerifyTokenService interface {
	TokenValid(r *http.Request) error
	ExtractAccessTokenToUserID(r *http.Request) (string, error)
}