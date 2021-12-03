package user_service

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository/mock"
)

func TestGetUserUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock.NewMockUserRepository(ctrl)

	userId := "dummy_user_id"
	input := GetUserUseCaseInput{userId}

	mockUserRepo.
		EXPECT().
		GetUser(gomock.Eq(userId)).
		Return(&model.User{
			ID:          "dummy_user_id",
			Account:     "dummy_user_account",
			DisplayName: "dummy_user_display_name",
			Password:    "dummy_user_password",
			Email:       "dummy_user_email",
		}, nil)

	uc := NewGetUserUseCase(mockUserRepo)

	output := GetUserUseCaseOutput{}
	if err := uc.Execute(input, &output); err != nil {
		t.Errorf("failed to execute use case")
	}

	assert.Equal(t, "dummy_user_id", output.ID)
	assert.Equal(t, "dummy_user_account", output.Account)
	assert.Equal(t, "dummy_user_display_name", output.DisplayName)
	assert.Equal(t, "dummy_user_email", output.Email)
}
