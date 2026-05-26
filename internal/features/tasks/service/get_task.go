package tasks_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (s *TasksService) GetTask(ctx context.Context, taskId int) (domain.Task, error) {
	taskDomain, err := s.tasksRepository.GetTask(ctx, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task from repository: %w", err)
	}
	return taskDomain, nil
}
