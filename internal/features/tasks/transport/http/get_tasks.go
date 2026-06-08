package tasks_transport

import (
	"fmt"
	"net/http"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

type GetTaskResponse []TaskDTO

func TasksDTOFromDomains(tasks []domain.Task) GetTaskResponse {
	var tasksDTO GetTaskResponse
	for _, task := range tasks {
		tasksDTO = append(tasksDTO, TaskDTO{ID: task.ID, Version: task.Version, Name: task.Name,
			Description: task.Description, Completed: task.Completed, CreatedAt: task.CreatedAt, CompletedAt: task.CompletedAt,
			AuthorId: task.AuthorId})
	}
	return tasksDTO
}

// GetTasks godoc
// @Summary Получить задачи
// @Description Получить информацию о всех существующих задачах в системе с опциональной пагинацией
// @Tags tasks
// @Produce json
// @Param limit query int false "Размер страницы с задачами"
// @Param offset query int false "Смещение страницы с задачами"
// @Success 200 {object} GetTaskResponse "Успешно полученные задачи"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks [get]
func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userId, limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to ger query params")
	}

	taskDomains, err := h.tasksService.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks from service")
	}

	tasksDTO := TasksDTOFromDomains(taskDomains)
	responseHandler.JSONResponse(tasksDTO, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	userId, err := core_http_utils.GetQueryParamInt(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get user id param: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	limit, err := core_http_utils.GetQueryParamInt(r, "limit")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get limit param: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	offset, err := core_http_utils.GetQueryParamInt(r, "offset")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get offset param: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return userId, limit, offset, nil
}
