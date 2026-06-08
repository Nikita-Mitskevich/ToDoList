package tasks_transport

import (
	"fmt"
	"net/http"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"
	core_http_request "restapi/internal/core/transport/http/request"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_types "restapi/internal/core/transport/http/types"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

type PatchTaskDTO TaskDTO

func PatchTaskDTOFromDomain(task domain.Task) PatchTaskDTO {
	return PatchTaskDTO{ID: task.ID, Version: task.Version, Name: task.Name,
		Description: task.Description, Completed: task.Completed, CreatedAt: task.CreatedAt, CompletedAt: task.CompletedAt,
		AuthorId: task.AuthorId}
}

type TaskPatchStruct struct {
	Name        core_http_types.Nullable[string] `swaggertype:"string" json:"full_name"`
	Description core_http_types.Nullable[string] `swaggertype:"string" json:"description"`
	Completed   core_http_types.Nullable[bool]   `swaggertype:"bool" json:"completed"`
}

func (t *TaskPatchStruct) Validate() error {
	if t.Name.Set {
		if t.Name.Value == nil {
			return fmt.Errorf("Title cant be null")
		}
		if len([]rune(*t.Name.Value)) < 1 || len([]rune(*t.Name.Value)) > 100 {
			return fmt.Errorf("Title must be between 1 and 100")
		}
	}

	if t.Description.Set {
		if t.Description.Value != nil {
			if len([]rune(*t.Description.Value)) < 1 || len([]rune(*t.Description.Value)) > 1000 {
				return fmt.Errorf("Description must be between 1 and 1000")
			}
		}
	}

	if t.Completed.Set {
		if t.Completed.Value == nil {
			return fmt.Errorf("Completed value cant be null")
		}
	}
	return nil
}

// PatchTask godoc
// @Summary Изменить задачу
// @Description Изменить информацию о существующей в системе задаче
// @Description 1. **Поле не передано**: `description` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `description:"сделать домашнее задание по математике"` - устанавливает новое описание в БД
// @Description 3. **Передан null**: `"description": null` - очищает поле в БД (set to NULL)
// @Description Ограничения: `title` и `completed` не могут быть выставлены как null
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Айди изменяемой задачи"
// @Param request body TaskPatchStruct true "PatchTask тело запроса"
// @Success 200 {object} PatchTaskDTO "Успешно измененная задача"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskId, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "get id path value")
		return
	}

	var request TaskPatchStruct
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "decode and validate error")
		return
	}

	taskPatch := TaskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskId, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}

	patchTaskDTO := PatchTaskDTOFromDomain(taskDomain)

	responseHandler.JSONResponse(patchTaskDTO, http.StatusOK)
}

func TaskPatchFromRequest(r TaskPatchStruct) domain.TaskPatch {
	return domain.TaskPatch{Name: r.Name.ToDomain(), Description: r.Description.ToDomain(), Completed: r.Completed.ToDomain()}
}
