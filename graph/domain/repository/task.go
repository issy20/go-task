package repository

import "github.com/issy20/go-task/graph/domain/model"

type ITaskRepository interface {
	CreateTask(task *model.Task) (*model.Task, error)
	DeleteTask(id string) (*model.Task, error)
	GetTaskByUser(userId string) ([]*model.Task, error)
}
