package core_pgx_pool

import (
	"errors"
	"fmt"
	core_postgres_pool "restapi/internal/core/repository/postgres/pool"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxQueryRow struct {
	pgx.Row
}

type pgxTag struct {
	pgconn.CommandTag
}

func (r pgxQueryRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return ProcessError(err)
	}
	return err
}

func ProcessError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrNoRows)
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrViolatesForeignKey)
		}
	}
	return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrUnknown)
}
