package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
)

//go:generate mockgen -destination=./mock/team_mock.go -package=mock . TeamRepository
type TeamRepository interface {
	AddTeam(team *model.Team) error
	GetTeam(id string) (*model.Team, error)
	GetTeamMembersWithRole(id string) ([]*internal.Member, error)
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

	if err := tr.db.Debug().Preload("Members").
		Where("team.id = ?", id).
		Take(&team).Error; err != nil {
		return nil, errors.New("user not found: " + id)
	}

	return &team, nil
}

func (tr *teamRepository) GetTeamMembersWithRole(id string) ([]*internal.Member, error) {
	members := []*internal.Member{}

	result := tr.db.Debug().
		Model(&model.Team{}).
		Select("user.*, role_relation.role").
		Joins("JOIN role_relation ON role_relation.team_id = team.id").
		Joins("JOIN user ON user.id = role_relation.user_id").
		Where("team.id = ?", id).
		Find(&members)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	return members, nil
}
