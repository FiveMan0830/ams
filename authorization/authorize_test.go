package authorization

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeWithToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjQ3MjMyMDAsInVzZXJfaWQiOiJ1c2VySWQifQ.klQyPw89jAZZR_B3OMvW158z2-g5wKleGUCBtyBFjm8"

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Println("getting new request failed")
		t.Fail()
	}
	req.Header.Add("Authorization" , "Bearer " + token)

	service := NewVerifyJWTService(&mockConfig{
		tokenSecret: "tokenSecret",
		expiredDays: 30,
	})

	err = service.TokenValid(req)
	assert.Nil(t, err)
}

func TestExtractAccessTokenToUserID(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjQ3MjMyMDAsInVzZXJfaWQiOiJ1c2VySWQifQ.klQyPw89jAZZR_B3OMvW158z2-g5wKleGUCBtyBFjm8"

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Println("getting new request failed")
		t.Fail()
	}
	req.Header.Add("Authorization" , "Bearer " + token)

	service := NewVerifyJWTService(&mockConfig{
		tokenSecret: "tokenSecret",
		expiredDays: 30,
	})

	var userID string
	userID, err = service.ExtractAccessTokenToUserID(req)
	assert.Nil(t, err)
	assert.Equal(t, "userId", userID)
}