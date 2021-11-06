package team_service

import (
	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type AddTeamUseCaseInput struct {
	Name string `json:"name" binding:"required"`
}

type AddTeamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewAddTeamUseCase(teamRepo repository.TeamRepository) *AddTeamUseCase {
	return &AddTeamUseCase{teamRepo}
}

func (uc *AddTeamUseCase) Execute(input AddTeamUseCaseInput) error {
	id := uuid.New()
	team := model.Team{
		ID:   id.String(),
		Name: input.Name,
	}

	if err := uc.teamRepo.AddTeam(&team); err != nil {
		return err
	}

	return nil
}
