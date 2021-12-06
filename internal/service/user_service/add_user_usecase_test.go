package user_service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository/mock"
)

// definition of user matcher
type userMatcher struct{}

func NewUserMatcher() gomock.Matcher { return &userMatcher{} }

func (um *userMatcher) Matches(x interface{}) bool {
	_, ok := x.(*model.User)
	return ok
}

func (um *userMatcher) String() string { return "match user" }

func TestAddUserUseCaseWithSingleUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock.NewMockUserRepository(ctrl)

	mockUserRepo.
		EXPECT().
		AddUser(NewUserMatcher()).Return(nil)

	input := AddUserUseCaseInput{
		Account:     "dummy_user_account",
		DisplayName: "dummy_user_display_name",
		Password:    "dummy_user_password",
		Email:       "dummy_user_email",
	}
	uc := NewAddUserUseCase(mockUserRepo)
	if err := uc.Execute(input); err != nil {
		t.Errorf("failed to execute use case")
	}
}
