package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type ExpelSubteamUseCaseInput struct {
	TeamId     string   `json:"teamId" binding:"required"`
	SubteamIds []string `json:"subteamIds" binding:"required"`
}

type ExpelSubteamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewExpelSubteamUseCase(teamRepo repository.TeamRepository) *ExpelSubteamUseCase {
	return &ExpelSubteamUseCase{teamRepo}
}

func (uc *ExpelSubteamUseCase) Execute(input ExpelSubteamUseCaseInput) error {
	teamId := input.TeamId
	subteams := input.SubteamIds

	if err := uc.teamRepo.RemoveSubteams(teamId, subteams); err != nil {
		return err
	}

	return nil
}
