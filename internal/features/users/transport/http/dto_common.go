package user

import "restapi/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	Name        string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{ID: user.ID, Version: user.Version, Name: user.FullName, PhoneNumber: user.PhoneNumber}
}

type UserDTORequest UserDTOResponse

func UserDTOfromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = UserDTOFromDomain(user)
	}
	return usersDTO
}
