package tasks_transport

import (
	"net/http"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

// DeleteTask godoc
// @Summary Удалить задачу
// @Description Удалить существующую задачу из системы
// @Tags tasks
// @Param id path int true "ID удаляемой задачи"
// @Success 204 "Успешное удаление задачи"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseWriter := core_http_response.NewHTTPResponseHandler(log, w)

	taskId, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseWriter.ErrorResponse(err, "get int path value")
		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskId); err != nil {
		responseWriter.ErrorResponse(err, "failed to delete task")
		return
	}

	responseWriter.NoContentResponse()
}
