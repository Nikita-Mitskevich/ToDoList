package statistics_repository

import (
	core_postgres_pool "restapi/internal/core/repository/postgres/pool"
)

type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsReposiory(pool core_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{pool: pool}
}
