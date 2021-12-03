package user_service

import (
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
	hashedPassword := pkg.HashWithSHA256(input.Password)
	encodedPassword := pkg.EncodeWithBase64(hashedPassword)

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
