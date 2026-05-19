package users_service

import (
	"context"
	"fmt"
)

func (h *UsersService) DeleteUser(ctx context.Context, id int) error {
	err := h.usersRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user from repository: %w", err)
	}
	return nil
}
