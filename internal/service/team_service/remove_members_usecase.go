package team_service

import "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"

type RemoveMembersUseCaseInput struct {
	TeamId  string   `json:"teamId"`
	UserIds []string `json:"userIds" binding:"required"`
}

type RemoveMembersUseCase struct {
	teamRepo repository.TeamRepository
}

func NewRemoveMembersUseCase(teamRepo repository.TeamRepository) *RemoveMembersUseCase {
	return &RemoveMembersUseCase{teamRepo}
}

func (uc *RemoveMembersUseCase) Execute(input RemoveMembersUseCaseInput) error {
	teamId := input.TeamId
	userIds := input.UserIds

	if err := uc.teamRepo.RemoveMembers(teamId, userIds); err != nil {
		return err
	}

	return nil
}
