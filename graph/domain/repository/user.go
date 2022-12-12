package repository

import (
	"github.com/issy20/go-task/graph/domain/model"
)

type IUserRepository interface {
	// Login(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
