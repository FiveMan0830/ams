package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
)

//go:generate mockgen -destination=./mock/team_mock.go -package=mock . TeamRepository
type TeamRepository interface {
	AddTeam(team *model.Team) error
	GetTeam(id string) (*model.Team, error)
	GetTeamMembersWithRole(id string) ([]*model.Member, error)
	AddMembers(relations []model.RoleRelation) error
	RemoveMembers(teamId string, userIds []string) error
	AddSubteams(subteams []model.TeamRelation) error
	RemoveSubteams(teamId string, subteams []string) error
	UpdateUserRole(teamId string, userId string, role int) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	db.AutoMigrate(&model.Team{})
	return &teamRepository{db}
}

func (tr *teamRepository) AddTeam(team *model.Team) error {
	if err := tr.db.Create(team).Error; err != nil {
		return err
	}
	return nil
}

func (tr *teamRepository) GetTeam(id string) (*model.Team, error) {
	team := model.Team{}

	if err := tr.db.Debug().
		Where("team.id = ?", id).
		Take(&team).Error; err != nil {
		return nil, errors.New("team not found: " + id)
	}

	return &team, nil
}

func (tr *teamRepository) GetTeamMembersWithRole(id string) ([]*model.Member, error) {
	members := []*model.Member{}
	results := []map[string]interface{}{}

	if err := tr.db.
		Table("team").
		Select("user.*, role_relation.role").
		Joins("JOIN role_relation ON role_relation.team_id = team.id").
		Joins("JOIN user ON user.id = role_relation.user_id").
		Where("team.id = ?", id).
		Find(&results).Error; err != nil {
		return nil, err
	}

	for _, result := range results {
		fmt.Println(result)
		members = append(members, &model.Member{
			User: model.User{
				ID:          result["id"].(string),
				Account:     result["account"].(string),
				DisplayName: result["display_name"].(string),
				Email:       result["email"].(string),
			},
			Role: result["role"].(int64),
		})
	}

	return members, nil
}

func (tr *teamRepository) AddMembers(members []model.RoleRelation) error {
	if err := tr.db.Create(&members).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) RemoveMembers(teamId string, userIds []string) error {
	members := []model.RoleRelation{}
	for _, userId := range userIds {
		members = append(members, model.RoleRelation{TeamID: teamId, UserID: userId})
	}

	if err := tr.db.
		Delete(members).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) AddSubteams(subteams []model.TeamRelation) error {
	if err := tr.db.Create(&subteams).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) RemoveSubteams(teamId string, subteams []string) error {
	teamRelation := []model.TeamRelation{}
	for _, subteamId := range subteams {
		teamRelation = append(teamRelation, model.TeamRelation{TeamID: teamId, SubteamID: subteamId})
	}

	if err := tr.db.
		Delete(teamRelation).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) UpdateUserRole(teamId string, userId string, role int) error {
	if err := tr.db.
		Model(&model.RoleRelation{TeamID: teamId, UserID: userId}).
		Update("role", role).Error; err != nil {
		return err
	}
	// result := tr.db.
	//	Model(&model.RoleRelation{TeamID: teamId, UserID: userId}).
	//	Update("role", role)


	return nil
}
