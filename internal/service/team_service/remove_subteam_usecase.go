package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type RemoveSubteamUseCaseInput struct {
	TeamId     string   `json:"teamId" binding:"required"`
	SubteamIds []string `json:"subteamIds" binding:"required"`
}

type RemoveSubteamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewRemoveSubteamUseCase(teamRepo repository.TeamRepository) *RemoveSubteamUseCase {
	return &RemoveSubteamUseCase{teamRepo}
}

func (uc *RemoveSubteamUseCase) Execute(input RemoveSubteamUseCaseInput) error {
	teamId := input.TeamId
	subteams := input.SubteamIds

	if err := uc.teamRepo.RemoveSubteams(teamId, subteams); err != nil {
		return err
	}

	return nil
}
