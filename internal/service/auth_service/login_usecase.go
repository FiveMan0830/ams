package auth_service

import (
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

type LoginUseCaseInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginUseCaseOutput struct {
	AccessToken string `json:"accessToken"`
}

type LoginUseCase struct {
	userRepo repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{userRepo}
}

func (uc *LoginUseCase) Execute(
	input LoginUseCaseInput,
	output *LoginUseCaseOutput,
) (string, error) {
	// check account
	user, err := uc.userRepo.GetUserByAccount(input.Account)
	if err != nil {
		return "", err
	}

	// check password
	hashedPassword := pkg.HashWithSHA256(input.Password)
	encodedPassword := pkg.EncodeWithBase64(hashedPassword)
	if user.Password != encodedPassword {
		return "", errors.New("incorrect password")
	}

	token, err := pkg.NewJWTClient(config.NewAuthConfig()).CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"uid":         user.ID,
	// 	"account":     user.Account,
	// 	"displayName": user.DisplayName,
	// 	"email":       user.Email,
	// })

	// jwtSecret := os.Getenv("JWT_SECRET")
	// tokenString, err := token.SignedString([]byte(jwtSecret))
	// if err != nil {
	// 	panic(err)
	// }

	return token, nil
}
