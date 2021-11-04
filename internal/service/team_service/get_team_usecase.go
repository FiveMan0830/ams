package team_service

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type GetTeamUseCaseInput struct {
	Id string `json:"id" binding:"required"`
}

type GetTeamUseCaseOuptut struct {
	Id      string             `json:"id"`
	Name    string             `json:"name"`
	Members []*internal.Member `json:"members"`
}

type GetTeamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewGetTeamUseCase(teamRepo repository.TeamRepository) GetTeamUseCase {
	return GetTeamUseCase{
		teamRepo: teamRepo,
	}
}

func (uc GetTeamUseCase) Execute(input GetTeamUseCaseInput, output *GetTeamUseCaseOuptut) error {
	team, err := uc.teamRepo.GetTeam(input.Id)
	members, err := uc.teamRepo.GetTeamMembersWithRole(input.Id)

	if err != nil {
		return err
	}

	output.Id = team.ID
	output.Name = team.Name
	output.Members = members

	return nil
}
