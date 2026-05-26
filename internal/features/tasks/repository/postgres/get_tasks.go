package tasks_postgres_repository

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

type TaskModels []TaskModel

func (r *TasksRepository) GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, title, description, completed, created_at, completed_at, author_user_id FROM todoapp.tasks
			  %s
			  LIMIT $1 OFFSET $2`

	args := []any{limit, offset}
	if userId != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id = $3")
		args = append(args, userId)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()

	var taskModels TaskModels
	for rows.Next() {
		var task TaskModel
		if err := rows.Scan(&task.ID, &task.Version, &task.Name, &task.Description,
			&task.Completed, &task.CreatedAt, &task.CompletedAt, &task.AuthorId); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		taskModels = append(taskModels, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}
	taskDomains := TaskDomainsFromModels(taskModels)
	return taskDomains, nil
}
