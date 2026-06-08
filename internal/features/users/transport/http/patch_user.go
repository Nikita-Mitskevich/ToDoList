package user

import (
	"fmt"
	"net/http"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"
	core_http_request "restapi/internal/core/transport/http/request"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_types "restapi/internal/core/transport/http/types"
	core_http_utils "restapi/internal/core/transport/http/utils"
	"strings"
)

type PatchUserStruct struct {
	FullName    core_http_types.Nullable[string] `swaggertype:"string"`
	PhoneNumber core_http_types.Nullable[string] `swaggertype:"string"`
}

// PatchUser godoc
// @Summary Изменить пользователя
// @Description Изменить информацию о существующем пользователе в системе
// @Description 1. **Поле не передано**: `phone_number` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `phone_number:"+375111113322"` - устанавливает новый номер в БД
// @Description 3. **Передан null**: `"phone_number": null` - очищает поле в БД (set to NULL)
// @Description Ограничения: `full_name` не может быть выставлен null
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "Айди изменяемого пользователя"
// @Param request body PatchUserStruct true "PatchUser тело запроса"
// @Success 200 {object} UserDTOResponse "Успешно измененный пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [patch]
func (r *PatchUserStruct) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("FullName cant be null")
		}
		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("FullName must be between 3 and 100")
		}

	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("PhoneNumber must be between 10 and 15")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("PhoneNumber must starts with '+'")
			}
		}
	}
	return nil
}

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to ger user id")
		return
	}

	var request PatchUserStruct
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.UsersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	responseDTO := UserDTOFromDomain(userDomain)

	responseHandler.JSONResponse(responseDTO, http.StatusOK)

	log.Debug(fmt.Sprintf("PatchUserRequest fields:\nFullName: %v\nPhoneNumber: %v", request.FullName, request.PhoneNumber))

}

func userPatchFromRequest(request PatchUserStruct) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
