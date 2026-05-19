package users_service

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (h *UsersService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	userDomain, err := h.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if err := userDomain.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply patch: %w", err)
	}

	responseUser, err := h.usersRepository.PatchUser(ctx, id, userDomain)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user service: %w", err)
	}

	return responseUser, nil
}
