package persistence

import (
	"github.com/issy20/go-task/graph/domain/model"
	"github.com/issy20/go-task/graph/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

var _ repository.IUserRepository = &userRepository{}

func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{Email: email}
	if err := u.db.Debug().First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
