package team_service

import "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"

type UpdateUserRoleUseCaseInput struct {
	TeamId string `json:"teamId"`
	UserId string `json:"userId"`
	Role   int    `json:"role"`
}

type UpdateUserRoleUseCase struct {
	teamRepo repository.TeamRepository
	userRepo repository.UserRepository
}

func NewUpdateUserRoleUseCase(teamRepo repository.TeamRepository, userRepo repository.UserRepository) *UpdateUserRoleUseCase {
	return &UpdateUserRoleUseCase{teamRepo, userRepo}
}

func (uc *UpdateUserRoleUseCase) Execute(input UpdateUserRoleUseCaseInput) error {
	teamId := input.TeamId
	userId := input.UserId
	role := input.Role

	if _, err := uc.userRepo.GetUser(userId); err != nil {
		return err
	}

	if _, err := uc.teamRepo.GetTeam(teamId); err != nil {
		return err
	}

	if err := uc.teamRepo.UpdateUserRole(teamId, userId, role); err != nil {
		return err
	}

	return nil
}
