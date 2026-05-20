package core_pgx_pool

import (
	"errors"
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
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrNoRows
		}
	}
	return err
}
