package user

import (
	"net/http"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
	core_http_utils "restapi/internal/core/transport/http/utils"
)

// GetUser godoc
// @Summary Получить пользователя
// @Description Получить информацию о существующем пользователе в системе
// @Tags users
// @Produce json
// @Param id path int true "ID получаемого пользователя"
// @Success 200 {object} UserDTOResponse "Успешно полученный пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err,
			"failed to get user id path value",
		)
		return
	}

	user, err := h.UsersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to ger user domain")
	}
	response := UserDTOFromDomain(user)
	responseHandler.JSONResponse(response, http.StatusOK)

}
