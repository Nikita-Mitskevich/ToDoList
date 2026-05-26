package tasks_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"
)

func (h *TasksService) GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	tasksDomains, err := h.tasksRepository.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("tasks service: %w", err)
	}
	return tasksDomains, nil

}
