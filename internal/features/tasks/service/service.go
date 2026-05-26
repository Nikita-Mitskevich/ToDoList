package tasks_service

import (
	"context"
	"restapi/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskId int) (domain.Task, error)
	DeleteTask(ctx context.Context, taskId int) error
	PatchTask(ctx context.Context, taskId int, taskPatch domain.Task) (domain.Task, error)
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
