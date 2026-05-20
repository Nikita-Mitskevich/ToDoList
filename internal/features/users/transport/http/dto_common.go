package user

import "restapi/internal/core/domain"

type UserDTOResponse struct {
	ID          int
	Version     int
	Name        string
	PhoneNumber *string
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
