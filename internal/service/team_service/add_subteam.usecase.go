package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type AddSubteamUseCaseInput struct {
	TeamId   string   `json:"teamId" binding:"required"`
	Subteams []string `json:"subteams" binding:"required"`
}

type AddSubteamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewAddSubteamUseCase(teamRepo repository.TeamRepository) *AddSubteamUseCase {
	return &AddSubteamUseCase{teamRepo}
}

func (uc *AddSubteamUseCase) Execute(input AddSubteamUseCaseInput) error {
	teamId := input.TeamId
	subteams := input.Subteams

	// convert data into model.TeamRelation
	teamRelations := []model.TeamRelation{}
	for _, subteamId := range subteams {
		teamRelations = append(teamRelations, model.TeamRelation{
			TeamID: teamId,
			SubteamID:  subteamId,
		})
	}

	if err := uc.teamRepo.AddSubteams(teamRelations); err != nil {
		return err
	}

	return nil
}
