package statistics_repository

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
	"strings"
	"time"
)

func (s *StatisticsRepository) GetTasks(ctx context.Context, userID *int, from *time.Time, to *time.Time) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, s.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`SELECT id, version, title, description, completed, created_at, completed_at, author_user_id FROM todoapp.tasks`)

	conditions := []string{}
	args := []any{}
	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)+1))
		args = append(args, userID)
	}
	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}
	if to != nil {
		conditions = append(conditions, fmt.Sprintf("completed_at<$%d", len(args)+1))
		args = append(args, to)
	}
	if len(args) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	rows, err := s.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)

	}
	defer rows.Close()
	var TaskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Name, &taskModel.Description,
			&taskModel.Completed, &taskModel.CreatedAt, &taskModel.CompletedAt, &taskModel.AuthorId); err != nil {
			return nil, fmt.Errorf("scan tasks: %w", err)
		}
		TaskModels = append(TaskModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan tasks: %w", err)
	}
	taskDomains := TaskDomainsFromModels(TaskModels)
	return taskDomains, nil
}
