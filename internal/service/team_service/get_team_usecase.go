package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type GetTeamUseCaseInput struct {
	Id string `json:"id" binding:"required"`
}

type GetTeamUseCaseOuptut struct {
	Team *model.Team `json:"team"`
}

type GetTeamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewGetTeamUseCase(teamRepo repository.TeamRepository) *GetTeamUseCase {
	return &GetTeamUseCase{teamRepo}
}

func (uc *GetTeamUseCase) Execute(input GetTeamUseCaseInput, output *GetTeamUseCaseOuptut) error {
	team, err := uc.teamRepo.GetTeam(input.Id)
	if err != nil {
		return err
	}

	members, err := uc.teamRepo.GetTeamMembersWithRole(input.Id)
	if err != nil {
		return err
	}

	team.Members = members
	output.Team = team

	return nil
}
