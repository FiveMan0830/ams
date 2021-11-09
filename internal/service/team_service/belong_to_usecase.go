package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type BelongToUseCaseInput struct {
	UserId string `json:"userid"`
}

type BelongToUseCaseOutput struct {
	Teams []model.Team `json:"teams"`
}

type BelongToUseCase struct {
	teamRepo repository.TeamRepository
}

func NewBelongToUseCase(teamRepo repository.TeamRepository) *BelongToUseCase {
	return &BelongToUseCase{teamRepo}
}

func (uc *BelongToUseCase) Execute(input BelongToUseCaseInput, output *BelongToUseCaseOutput) error {
	userId := input.UserId

	teams, err := uc.teamRepo.GetBelongTeam(userId)

	if err != nil {
		return err
	}

	output.Teams = teams

	return nil
}
