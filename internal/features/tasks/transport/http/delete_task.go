package tasks_transport

import (
	"net/http"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

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
