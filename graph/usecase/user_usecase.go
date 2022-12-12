package usecase

import (
	"errors"

	"github.com/issy20/go-task/graph/domain/model"
	"github.com/issy20/go-task/graph/domain/repository"
	"github.com/issy20/go-task/graph/pkg/auth"
	"github.com/issy20/go-task/graph/utils"
)

type IUserUsecase interface {
	CreateUser(input model.NewUser) (*model.User, error)
	Login(loginUser model.LoginUser) (string, error)
}

type userUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateUser(input model.NewUser) (*model.User, error) {
	uuid := utils.GenRandomUUID()
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        uuid,
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: utils.CurrentTime(),
		UpdatedAt: utils.CurrentTime(),
	}

	return u.UserRepository.CreateUser(user)
}

func (u *userUsecase) Login(loginUser model.LoginUser) (string, error) {
	user, err := u.UserRepository.GetUserByEmail(loginUser.Email)
	if !utils.CheckPasswordHash(loginUser.Password, user.Password) {
		return "", errors.New("fail to authenticate")
	}
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
