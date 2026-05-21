package tasks_transport

import (
	"net/http"
	core_logger "restapi/internal/core/logger"
	core_http_request "restapi/internal/core/transport/http/request"
	core_http_response "restapi/internal/core/transport/http/response"
	"time"
)

type CreateTaskRequest struct {
	Name string `validate:"required,min=3,max=100"`
	Description *string `validate:"omitempty,min=1,max=1000"`
}

type CreateTaskResponse struct {
	ID int
	Version int

	Name string
	Description *string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time

	AuthorId int
}

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler:=core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("task create request incoming")

	var taskDTO CreateTaskRequest
	if err:=core_http_request.DecodeAndValidateRequest(r, &taskDTO); err!=nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate task")
	}

	



}