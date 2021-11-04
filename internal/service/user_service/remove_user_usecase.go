package user_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type RemoveUserUseCaseInput struct {
	Id string `json:"id" binding:"required"`
}

type RemoveUserUseCase struct {
	userRepo repository.UserRepository
}

func NewRemoveUserUseCase(userRepo repository.UserRepository) RemoveUserUseCase {
	return RemoveUserUseCase{
		userRepo: userRepo,
	}
}

func (uc RemoveUserUseCase) Execute(input RemoveUserUseCaseInput) error {
	if err := uc.userRepo.RemoveUser(input.Id); err != nil {
		return err
	}
	return nil
}
