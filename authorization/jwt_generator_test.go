package authorization

import (
	"testing"
	"time"

	// "errors"

	"github.com/stretchr/testify/assert"
)

type mockConfig struct {
	tokenSecret string
	expiredDays int
}

func (mc *mockConfig) GetTokenSecret() string {
	return mc.tokenSecret
}

func (mc *mockConfig) GetTokenExpiredTime() int64 {
	return time.Date(2021, 5, 28, 0, 0, 0, 0, time.Local).
		Add(time.Hour * 24 * time.Duration(mc.expiredDays)).
		Unix()
}

func TestCreateToken(t *testing.T) {
	tokenGenerator := NewJWTGenerator(&mockConfig{
		tokenSecret: "tokenSecret",
		expiredDays: 30,
	})
	token, _ := tokenGenerator.CreateToken("userId")
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjQ3MjMyMDAsInVzZXJfaWQiOiJ1c2VySWQifQ.klQyPw89jAZZR_B3OMvW158z2-g5wKleGUCBtyBFjm8"

	assert.Equal(t, expectedToken, token)
}
