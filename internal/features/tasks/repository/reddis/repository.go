package tasks_reddis_repository

import (
	core_redis_pool "restapi/internal/core/repository/redis/pool"
	tasks_postgres_repository "restapi/internal/features/tasks/repository/postgres"
)

type TasksRepository struct {
	reddis         core_redis_pool.Pool
	mainRepository *tasks_postgres_repository.TasksRepository
}

func NewTasksRepository(reddis core_redis_pool.Pool, main *tasks_postgres_repository.TasksRepository) *TasksRepository {
	return &TasksRepository{reddis: reddis, mainRepository: main}
}
