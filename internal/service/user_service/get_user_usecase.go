package user_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type GetUserUseCaseInput struct {
	Id string `json:"id" binding:"required"`
}

type GetUserUseCaseOutput struct {
	ID          string `json:"id"`
	Account     string `json:"account"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type GetUserUseCase struct {
	userRepo repository.UserRepository
}

func NewGetUserUseCase(userRepo repository.UserRepository) GetUserUseCase {
	return GetUserUseCase{
		userRepo: userRepo,
	}
}

func (uc GetUserUseCase) Execute(input GetUserUseCaseInput, output *GetUserUseCaseOutput) error {
	user, err := uc.userRepo.GetUser(input.Id)
	if err != nil {
		return err
	}

	output.ID = user.ID
	output.Account = user.Account
	output.DisplayName = user.DisplayName
	output.Email = user.Email

	return nil
}
