package resolver

import (
	"github.com/issy20/go-task/graph/infra/persistence"
	"github.com/issy20/go-task/graph/usecase"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct {
// 	UserService repository.UserService
// }

type Resolver struct {
	UserUsecase usecase.IUserUsecase
	TaskUsecase usecase.ITaskUsecase
}

func NewResolver(db *gorm.DB) *Resolver {
	userRepository := persistence.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	taskRepository := persistence.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	return &Resolver{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	}
}
