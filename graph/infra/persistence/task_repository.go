package persistence

import (
	"github.com/issy20/go-task/graph/domain/model"
	"github.com/issy20/go-task/graph/domain/repository"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

var _ repository.ITaskRepository = &taskRepository{}

func (t *taskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	if err := t.db.Create(task).Error; err != nil {
		return nil, err
	}
	// if err := t.db.Debug().Create()
	return task, nil
}

func (t *taskRepository) GetTaskByUser(userId string) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := t.db.Debug().Model(tasks).Where("user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepository) DeleteTask(id string) (*model.Task, error) {
	var task = &model.Task{}
	if err := t.db.Debug().Model(task).Where("id = ?", id).Take(task).Delete(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
