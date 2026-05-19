package users_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (h *UsersService) GetUser(ctx context.Context, id int) (domain.User, error) {

	user, err := h.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user from repository: %w", err)
	}
	return user, nil
}
