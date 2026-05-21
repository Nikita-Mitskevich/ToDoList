package tasks_transport

import (
	"context"
	"restapi/internal/core/domain"
	"restapi/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{tasksService: tasksService}
}

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

func (h *TasksHTTPHandler) Routes() []server.Route {
	return []server.Route{}
}
