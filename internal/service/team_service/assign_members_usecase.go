package team_service

import (
	"fmt"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
)

type AssignMembersUseCaseInput struct {
	TeamId string `json:"teamId" binding:"required"`
	Users  []struct {
		UserId string `json:"id"`
		Role   int    `json:"role"`
	} `json:"users" binding:"required"`
}

type AssignMembersUseCase struct {
	teamRepo repository.TeamRepository
}

func NewAssignMembersUseCase(teamRepo repository.TeamRepository) *AssignMembersUseCase {
	return &AssignMembersUseCase{teamRepo}
}

func (uc *AssignMembersUseCase) Execute(input AssignMembersUseCaseInput) error {
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
