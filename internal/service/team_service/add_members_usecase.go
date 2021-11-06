package team_service

import (
	"fmt"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type AddMembersUseCaseInput struct {
	TeamId string `json:"teamId" binding:"required"`
	Users  []struct {
		UserId string `json:"id"`
		Role   int    `json:"role"`
	} `json:"users" binding:"required"`
}

type AddMembersUseCase struct {
	teamRepo repository.TeamRepository
}

func NewAddMembersUseCase(teamRepo repository.TeamRepository) *AddMembersUseCase {
	return &AddMembersUseCase{teamRepo}
}

func (uc *AddMembersUseCase) Execute(input AddMembersUseCaseInput) error {
	teamId := input.TeamId
	users := input.Users

	fmt.Println("team id: ", teamId)
	for _, userId := range users {
		fmt.Println("user: ", userId)
	}

	// convert data into model.RoleRelation
	members := []model.RoleRelation{}
	for _, user := range users {
		members = append(members, model.RoleRelation{
			TeamID: teamId,
			UserID: user.UserId,
			Role:   user.Role,
		})
	}

	if err := uc.teamRepo.AddMembers(members); err != nil {
		return err
	}

	return nil
}
