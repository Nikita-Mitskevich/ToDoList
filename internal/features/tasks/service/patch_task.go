package tasks_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, taskId int, taskPatch domain.TaskPatch) (domain.Task, error) {
	taskDomain, err := s.tasksRepository.GetTask(ctx, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task from repository: %w", err)
	}

	if err := taskDomain.ApplyPatch(taskPatch); err != nil {
		return domain.Task{}, fmt.Errorf("apply patch: %w", err)
	}

	fmt.Println(taskDomain.Name)
	taskDomain, err = s.tasksRepository.PatchTask(ctx, taskId, taskDomain)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	return taskDomain, nil
}
