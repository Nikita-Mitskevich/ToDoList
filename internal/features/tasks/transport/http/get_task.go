package tasks_transport

import (
	"net/http"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

type GetTaskResponseDTO TaskDTO

func GetTaskDTOFromDomain(task domain.Task) GetTaskResponseDTO {
	return GetTaskResponseDTO{ID: task.ID, Version: task.Version, Name: task.Name,
		Description: task.Description, Completed: task.Completed, CreatedAt: task.CreatedAt, CompletedAt: task.CompletedAt,
		AuthorId: task.AuthorId}
}

// GetTask godoc
// @Summary Получить задачу
// @Description Получить информацию о существующей задаче
// @Tags tasks
// @Produce json
// @Param id path int true "ID получаемой задачи"
// @Success 200 {object} GetTaskResponseDTO  "Успешно полученная задача"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseWriter := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseWriter.ErrorResponse(err, "get int path value")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseWriter.ErrorResponse(err, "failed to get task from service")
		return
	}

	taskDTO := GetTaskDTOFromDomain(taskDomain)
	responseWriter.JSONResponse(taskDTO, http.StatusOK)

}
