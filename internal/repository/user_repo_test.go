package repository

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
)

func setupUserRepo() (*gorm.DB, UserRepository) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	userRepo := NewUserRepository(db)

	return db, userRepo
}

func TestAddUser(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_password",
	}
	repo.AddUser(user)

	result := model.User{}
	db.Where("id = ?", user.ID).First(&result)
	defer db.Delete(user)

	assert.Equal(t, result.ID, user.ID)
	assert.Equal(t, result.Account, user.Account)
	assert.Equal(t, result.DisplayName, user.DisplayName)
	assert.Equal(t, result.Password, user.Password)
	assert.Equal(t, result.Email, user.Email)
}

func TestGetUser(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_password",
	}
	db.Create(user)

	result, err := repo.GetUser(user.ID)
	if err != nil {
		t.Errorf("failed to get user %s", user.ID)
	}
	defer db.Delete(user)

	assert.Equal(t, result.ID, user.ID)
	assert.Equal(t, result.Account, user.Account)
	assert.Equal(t, result.DisplayName, user.DisplayName)
	assert.Equal(t, result.Password, user.Password)
	assert.Equal(t, result.Email, user.Email)
}

func TestGetUserByAccount(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_password",
	}
	db.Create(user)

	result, err := repo.GetUserByAccount(user.Account)
	if err != nil {
		t.Errorf("failed to get user %s", user.Account)
	}
	defer db.Delete(user)

	assert.Equal(t, result.ID, user.ID)
	assert.Equal(t, result.Account, user.Account)
	assert.Equal(t, result.DisplayName, user.DisplayName)
	assert.Equal(t, result.Password, user.Password)
	assert.Equal(t, result.Email, user.Email)
}

func TestEditPartOfUserData(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_password",
	}
	db.Create(user)

	updatedUser := &model.User{
		ID:          "dummy_id",
		DisplayName: "new_dummy_display_name",
	}
	if err := repo.UpdateUser(updatedUser); err != nil {
		t.Errorf("failed to update user %s", updatedUser.ID)
	}

	result := model.User{}
	db.Where("id = ?", user.ID).First(&result)
	defer db.Delete(user)

	assert.Equal(t, result.DisplayName, updatedUser.DisplayName)
}

func TestEditAllOfUserData(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_email",
	}
	db.Create(user)

	updatedUser := &model.User{
		ID:          "dummy_id",
		DisplayName: "new_dummy_display_name",
		Email:       "new_dummy_email",
	}
	if err := repo.UpdateUser(updatedUser); err != nil {
		t.Errorf("failed to update user %s", updatedUser.ID)
	}

	result := model.User{}
	db.Where("id = ?", user.ID).First(&result)
	defer db.Delete(user)

	assert.Equal(t, result.DisplayName, updatedUser.DisplayName)
	assert.Equal(t, result.Email, updatedUser.Email)
}

func TestRemoveUser(t *testing.T) {
	db, repo := setupUserRepo()

	user := &model.User{
		ID:          "dummy_id",
		Account:     "dummy_account",
		DisplayName: "dummy_display_name",
		Password:    "dummy_password",
		Email:       "dummy_email",
	}
	db.Create(user)

	if err := repo.RemoveUser(user.ID); err != nil {
		t.Errorf("failed to remove user %s", user.ID)
	}

	result := model.User{}
	err := db.Where("id = ?", user.ID).First(&result).Error

	assert.Equal(t, true, errors.Is(err, gorm.ErrRecordNotFound))
}
