package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/issy20/go-task/graph/domain/model"
	"github.com/issy20/go-task/graph/domain/repository"
	"github.com/issy20/go-task/graph/middleware"
	"github.com/issy20/go-task/graph/utils"
)

type ITaskUsecase interface {
	CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error)
	GetTaskByUser(ctx context.Context, userID string) ([]*model.Task, error)
	DeleteTask(ctx context.Context, id string) (*model.Task, error)
}

type taskUsecase struct {
	TaskRepository repository.ITaskRepository
}

func NewTaskUsecase(repo repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{repo}
}

func (t *taskUsecase) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	uuid := utils.GenRandomUUID()
	user := middleware.ForContext(ctx)
	log.Println(user)

	if user == nil {
		return &model.Task{}, fmt.Errorf("access denied")
	}

	task := &model.Task{
		ID:        uuid,
		Title:     input.Title,
		CreatedAt: utils.CurrentTime(),
		UpdatedAt: utils.CurrentTime(),
		UserID:    user.ID,
	}

	return t.TaskRepository.CreateTask(task)
}

func (t *taskUsecase) GetTaskByUser(ctx context.Context, userID string) ([]*model.Task, error) {
	user := middleware.ForContext(ctx)

	return t.TaskRepository.GetTaskByUser(user.ID)
}

func (t *taskUsecase) DeleteTask(ctx context.Context, id string) (*model.Task, error) {
	return t.TaskRepository.DeleteTask(id)
}
