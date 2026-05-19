package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"restapi/internal/core/domain"
	core_errors "restapi/internal/core/errors"

	"github.com/jackc/pgx/v5"
)

func (h *UsersRepository) PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, h.pool.OpTimeout())
	defer cancel()

	query := `UPDATE todoapp.users
			  SET full_name = $1,
			  phone_number=$2,
			  version=version+1
			  WHERE id=$3 AND version=$4
			  RETURNING id, version, full_name, phone_number`

	row := h.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, id, user.Version)
	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id = %d concurrently accessed: %w", id, core_errors.ErrConflict)
		} else {
			return domain.User{}, fmt.Errorf("scan user model: %w", err)
		}
	}
	userDomain := userDomainFromModel(userModel)
	return userDomain, nil
}
