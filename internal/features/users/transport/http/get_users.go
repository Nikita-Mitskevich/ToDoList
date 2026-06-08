package user

import (
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

// GetUsers godoc
// @Summary Получить пользователей
// @Description Получить информацию о всех существующих пользователях в системе с опциональной пагинацией
// @Tags users
// @Produce json
// @Param limit query int false "Размер страницы с пользователями"
// @Param offset query int false "Смещение страницы с пользователями"
// @Success 200 {object} GetUsersResponse "Успешно полученные пользователи"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [get]
func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query param")
		return
	}
	userDomains, err := h.UsersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user domains")
		return
	}

	response := GetUsersResponse(UserDTOfromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetQueryParamInt(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit param: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	offset, err := core_http_utils.GetQueryParamInt(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset param: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return limit, offset, nil
}
