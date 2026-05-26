package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"
	core_postgres_pool "restapi/internal/core/repository/postgres/pool"
)

type GetTaskModel TaskModel

func GetTaskDomainFromModel(task GetTaskModel) domain.Task {
	return domain.Task{ID: task.ID, Version: task.Version,
		Name: task.Name, Description: task.Description,
		Completed: task.Completed, CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt, AuthorId: task.AuthorId}
}

func (r *TasksRepository) GetTask(ctx context.Context, taskId int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, title, description, completed, created_at, completed_at, author_user_id FROM todoapp.tasks
		    WHERE id = $1`

	row := r.pool.QueryRow(ctx, query, taskId)

	var taskModel GetTaskModel
	if err := row.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Name, &taskModel.Description,
		&taskModel.Completed, &taskModel.CreatedAt, &taskModel.CompletedAt, &taskModel.AuthorId); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("scan task: %w", core_errors.ErrNotFound)
		} else {
			return domain.Task{}, fmt.Errorf("scan error: %w", err)
		}
	}

	taskDomain := GetTaskDomainFromModel(taskModel)
	return taskDomain, nil
}
