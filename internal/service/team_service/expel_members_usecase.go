package team_service

import "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"

type ExpelMembersUseCaseInput struct {
	TeamId  string   `json:"teamId"`
	UserIds []string `json:"userIds" binding:"required"`
}

type ExpelMembersUseCase struct {
	teamRepo repository.TeamRepository
}

func NewExpelMembersUseCase(teamRepo repository.TeamRepository) *ExpelMembersUseCase {
	return &ExpelMembersUseCase{teamRepo}
}

func (uc *ExpelMembersUseCase) Execute(input ExpelMembersUseCaseInput) error {
	teamId := input.TeamId
	userIds := input.UserIds

	if err := uc.teamRepo.RemoveMembers(teamId, userIds); err != nil {
		return err
	}

	return nil
}
