package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type AssignSubteamUseCaseInput struct {
	TeamId   string   `json:"teamId" binding:"required"`
	Subteams []string `json:"subteams" binding:"required"`
}

type AssignSubteamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewAssignSubteamUseCase(teamRepo repository.TeamRepository) *AssignSubteamUseCase {
	return &AssignSubteamUseCase{teamRepo}
}

func (uc *AssignSubteamUseCase) Execute(input AssignSubteamUseCaseInput) error {
	teamId := input.TeamId
	subteams := input.Subteams

	// convert data into model.TeamRelation
	teamRelations := []model.TeamRelation{}
	for _, subteamId := range subteams {
		teamRelations = append(teamRelations, model.TeamRelation{
			TeamID:    teamId,
			SubteamID: subteamId,
		})
	}

	if err := uc.teamRepo.AddSubteams(teamRelations); err != nil {
		return err
	}

	return nil
}
