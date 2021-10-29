package repository

import (
	"gorm.io/gorm"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&model.User{})
	return &UserRepository{db}
}

func (ur *UserRepository) AddUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
