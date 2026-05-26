package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"
	core_postgres_pool "restapi/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PatchTask(ctx context.Context, taskId int, taskPatch domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `UPDATE todoapp.tasks SET
	version = version + 1,
	title = $1,
	description = $2,
	completed = $3,
	created_at = $4,
	completed_at = $5
	WHERE id = $6 AND version = $7
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id
	`
	row := r.pool.QueryRow(ctx, query, taskPatch.Name, taskPatch.Description, taskPatch.Completed,
		taskPatch.CreatedAt, taskPatch.CompletedAt, taskId, taskPatch.Version)

	var patchTaskModel TaskModel

	if err := row.Scan(&patchTaskModel.ID, &patchTaskModel.Version, &patchTaskModel.Name,
		&patchTaskModel.Description, &patchTaskModel.Completed, &patchTaskModel.CreatedAt,
		&patchTaskModel.CompletedAt, &patchTaskModel.AuthorId); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with id = %d concurrently accessed: %w", taskId, core_errors.ErrConflict)
		} else {
			return domain.Task{}, fmt.Errorf("scan task model: %w", err)
		}
	}
	taskDomain := TaskDomainFromModel(patchTaskModel)
	return taskDomain, nil
}
