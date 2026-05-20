package users_postgres_repository

import (
	"context"
	"fmt"
	"restapi/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, full_name, phone_number
			  FROM todoapp.users
			  ORDER BY id ASC
			  LIMIT $1
			  OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}
	defer rows.Close()
	var users []UserModel
	for rows.Next() {
		var user UserModel
		err = rows.Scan(&user.ID, &user.Version, &user.FullName, &user.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to scan users: %w", err)
		}
		users = append(users, user)
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("next rows: %w", err)
		}
	}

	userDomains := userDomainFronModels(users)
	return userDomains, nil

}
