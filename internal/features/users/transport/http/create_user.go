package user

import (
	"net/http"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"
	core_http_request "restapi/internal/core/transport/http/request"
	core_http_response "restapi/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Name        string  `validate:"required,min=3,max=100" example:"Vasya Kalupin" json:"full_name"`
	PhoneNumber *string `validate:"omitempty,min=10,max=15,startswith=+" example:"+375123456789" json:"phone_number"`
}

type CreateUserResponse UserDTOResponse

// CreateUser godoc
// @Summary Создать пользователя
// @Description Создать нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "CreateUser тело запроса"
// @Success 201 {object} CreateUserResponse "Успешно созданный пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [post]
func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("user create request incoming")
	var user CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &user); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(user)

	userDomain, err := h.UsersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusCreated)

}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.Name, dto.PhoneNumber)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{ID: user.ID, Version: user.Version, Name: user.FullName, PhoneNumber: user.PhoneNumber}
}
