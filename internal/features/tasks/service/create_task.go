package tasks_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (h *TasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task: %w", err)
	}

	taskDomain, err := h.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("cant get task from repository: %w", err)
	}
	return taskDomain, nil
}
