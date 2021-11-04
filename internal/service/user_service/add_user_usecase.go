package user_service

import (
	"encoding/base64"

	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

type AddUserUseCaseInput struct {
	Account     string `json:"account" binding:"required" validate:"max=255"`
	DisplayName string `json:"displayName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required" validate:"email,max=255"`
}

type AddUserUseCase struct {
	userRepo repository.UserRepository
}

func NewAddUserUseCase(userRepo repository.UserRepository) AddUserUseCase {
	return AddUserUseCase{
		userRepo: userRepo,
	}
}

func (uc AddUserUseCase) Execute(input AddUserUseCaseInput) error {
	hashedPassword := uc.hashPassword(input.Password)
	encodedPassword := uc.encodePassword(hashedPassword)

	id := uuid.New()
	user := model.User{
		ID:          id.String(),
		Account:     input.Account,
		DisplayName: input.DisplayName,
		Password:    encodedPassword,
		Email:       input.Email,
	}
	if err := uc.userRepo.AddUser(&user); err != nil {
		return err
	}

	return nil
}

func (AddUserUseCase) hashPassword(password string) []byte {
	hasher := pkg.NewSHA256Client()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		panic(err)
	}
	return hasher.Sum(nil)
}

func (AddUserUseCase) encodePassword(password []byte) string {
	encoder := pkg.NewBase64Client()
	defer encoder.Close()

	return base64.StdEncoding.EncodeToString(password)
}
