package repository

import (
	"errors"

	"gorm.io/gorm"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
)

//go:generate mockgen -destination=./mock/user_mock.go -package=mock . UserRepository
type UserRepository interface {
	AddUser(user *model.User) error
	GetUser(id string) (*model.User, error)
	RemoveUser(id string) error
	UpdateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// constructor for user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&model.User{})
	return &userRepository{db}
}

func (ur *userRepository) AddUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUser(id string) (*model.User, error) {
	user := model.User{}
	if err := ur.db.
		Where("id = ?", id).
		Take(&user).Error; err != nil {
		return nil, errors.New("user not found: " + id)
	}
	return &user, nil
}

func (ur *userRepository) RemoveUser(id string) error {
	result := ur.db.Where("id = ?", id).Delete(&model.User{})

	if result.Error != nil {
		return errors.New("failed to remove user")
	}

	return nil
}

func (ur *userRepository) UpdateUser(user *model.User) error {
	result := ur.db.Updates(user)

	if result.Error != nil {
		return errors.New("failed to remove user")
	}

	return nil
}
