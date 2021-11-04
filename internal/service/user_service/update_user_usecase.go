package user_service

import (
	"encoding/base64"
	"errors"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

type UpdateUserUseCaseInput struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
}

type UpdateUserUseCase struct {
	userRepo repository.UserRepository
}

func NewUpdateUserUseCase(userRepo repository.UserRepository) UpdateUserUseCase {
	return UpdateUserUseCase{
		userRepo: userRepo,
	}
}

func (uc UpdateUserUseCase) Execute(input UpdateUserUseCaseInput) error {
	user, err := uc.userRepo.GetUser(input.Id)
	if err != nil {
		return errors.New("user not found: " + input.Id)
	}

	user.DisplayName = input.DisplayName
	user.Email = input.Email

	if err := uc.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (uc UpdateUserUseCase) updatePassword(password string) string {
	if password == "" {
		return ""
	}

	hasher := pkg.NewSHA256Client()
	encoder := pkg.NewBase64Client()
	defer encoder.Close()

	_, err := hasher.Write([]byte(password))
	if err != nil {
		panic(err)
	}
	encodedPassword := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return encodedPassword
}
