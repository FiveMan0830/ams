package user_service

import (
	"encoding/base64"

	"github.com/google/uuid"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/repository"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/pkg"
)

type AddUserUseCaseInput struct {
	Account     string
	DisplayName string
	Password    string
	Email       string
}

type AddUserUseCase struct {
	userRepo *repository.UserRepository
}

func NewAddUserUseCase(userRepo *repository.UserRepository) AddUserUseCase {
	return AddUserUseCase{
		userRepo: userRepo,
	}
}

func (uc AddUserUseCase) Execute(input AddUserUseCaseInput) error {
	id := uuid.New()

	hasher := pkg.NewSHA256Client()
	encoder := pkg.NewBase64Client()
	defer encoder.Close()

	_, err := hasher.Write([]byte(input.Password))
	if err != nil {
		panic(err)
	}

	encodedPassword := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	user := model.User{
		ID:          id.String(),
		Account:     input.Account,
		DisplayName: input.DisplayName,
		Password:    encodedPassword,
	}
	if err := uc.userRepo.AddUser(&user); err != nil {
		return err
	}
	return nil
}
