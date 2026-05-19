package users_postgres_repository

import (
	"context"
	"fmt"
	core_errors "restapi/internal/core/errors"
)

func (h *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, h.pool.OpTimeout())
	defer cancel()
	query := `DELETE FROM todoapp.users
			  WHERE id = $1`

	commandTag, err := h.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id: %d: %w", id, core_errors.ErrNotFound)
	}

	return nil

}
