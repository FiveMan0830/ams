package user_service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository/mock"
)

func TestAddUserUseCaseWithSingleUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock.NewMockUserRepository(ctrl)

	mockUserRepo.
		EXPECT().
		AddUser(gomock.Eq(&model.User{
			ID:          "00000000-0000-0000-0000-000000000000",
			Account:     "dummy_user_account",
			DisplayName: "dummy_user_display_name",
			Password:    "dummy_user_password",
			Email:       "dummy_user_email",
		})).
		Return(nil)

	// uc := NewAddUserUseCase(mockUserRepo)
	// if err := uc.Execute(); err != nil {
	//
	// }
}

func TestAddUserUseCaseTwiceWithIdenticalUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock.NewMockUserRepository(ctrl)

	mockUserRepo.
		EXPECT().
		AddUser(gomock.Eq(&model.User{
			ID:          "00000000-0000-0000-0000-000000000000",
			Account:     "dummy_user_account",
			DisplayName: "dummy_user_display_name",
			Password:    "dummy_user_password",
			Email:       "dummy_user_email",
		})).
		Return(nil)

	// uc := NewAddUserUseCase(mockUserRepo)
	// if err := uc.Execute(); err != nil {
	//
	// }

	// output
}
