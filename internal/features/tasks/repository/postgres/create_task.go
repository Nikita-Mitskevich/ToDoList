package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"
	core_postgres_pool "restapi/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `INSERT INTO todoapp.tasks (title, description, completed, created_at, completed_at, author_user_id)
			  VALUES($1, $2, $3, $4, $5, $6)
			  RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id`

	row := r.pool.QueryRow(ctx, query, task.Name, task.Description, task.Completed, task.CreatedAt, task.CompletedAt, task.AuthorId)
	var taskModel TaskModel
	if err := row.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Name, &taskModel.Description, &taskModel.Completed, &taskModel.CreatedAt,
		&taskModel.CompletedAt, &taskModel.AuthorId); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf("%v:user id:%d: %w", err, task.ID, core_errors.ErrConflict)
		} else {
			return domain.Task{}, fmt.Errorf("Scan error: %w", err)
		}
	}
	domainTask := TaskDomainFromModel(taskModel)
	return domainTask, nil
}
