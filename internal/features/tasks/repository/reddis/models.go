package tasks_reddis_repository

import (
	"fmt"
	"restapi/internal/core/domain"
	"strconv"
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

type TaskModels []TaskModel

func TaskDomainFromModel(task TaskModel) domain.Task {
	return domain.Task{ID: task.ID, Version: task.Version,
		Name: task.Name, Description: task.Description,
		Completed: task.Completed, CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt, AuthorId: task.AuthorId}
}

func TaskDomainsFromModels(tasks TaskModels) []domain.Task {
	var taskDomains []domain.Task
	for _, task := range tasks {
		taskDomain := TaskDomainFromModel(task)
		taskDomains = append(taskDomains, taskDomain)
	}
	return taskDomains
}

func TaskDomainToModel(task domain.Task) TaskModel {
	return TaskModel{ID: task.ID, Version: task.Version,
		Name: task.Name, Description: task.Description,
		Completed: task.Completed, CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt, AuthorId: task.AuthorId}
}
func TaskDomainsToModels(tasks []domain.Task) TaskModels {
	var taskDomains TaskModels
	for _, task := range tasks {
		taskDomain := TaskDomainToModel(task)
		taskDomains = append(taskDomains, taskDomain)
	}
	return taskDomains
}

func GetKeyValue(taskId int) string {
	return fmt.Sprintf("task:%d", taskId)
}

func GetTasksKeyValue(userId *int) string {
	if userId != nil {
		return fmt.Sprintf("tasks:%d", *userId)
	}
	return "tasks:all"
}

func TasksListField(limit *int, offset *int) string {
	prt := func(v *int) string {
		if v == nil {
			return "nil"
		}
		return strconv.Itoa(*v)
	}
	return fmt.Sprintf("%s:%s", prt(limit), prt(offset))
}
