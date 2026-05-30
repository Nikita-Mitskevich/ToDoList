package statistics_repository

import (
	"restapi/internal/core/domain"
	"time"
)

type TaskModel struct {
	ID      int
	Version int

	Name        string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorId int
}

func TaskDomainFromModel(task TaskModel) domain.Task {
	return domain.Task{ID: task.ID, Version: task.Version,
		Name: task.Name, Description: task.Description,
		Completed: task.Completed, CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt, AuthorId: task.AuthorId}
}

func TaskDomainsFromModels(tasks []TaskModel) []domain.Task {
	var taskDomains []domain.Task
	for _, task := range tasks {
		taskDomain := TaskDomainFromModel(task)
		taskDomains = append(taskDomains, taskDomain)
	}
	return taskDomains
}
