package tasks_reddis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"
	core_redis_pool "restapi/internal/core/repository/redis/pool"

	"go.uber.org/zap"
)

func (r *TasksRepository) GetTasks(ctx context.Context, userId *int, limit *int, offset *int) ([]domain.Task, error) {
	log := core_logger.FromContext(ctx)
	key := GetTasksKeyValue(userId)
	options := TasksListField(limit, offset)
	val, err := r.reddis.HGet(ctx, key, options).Bytes()
	if err != nil {
		if !errors.Is(err, core_redis_pool.ErrNotFound) {
			log.Error("hget task list", zap.Error(err))
		}
	} else {
		var taskModels TaskModels
		if err := json.Unmarshal(val, &taskModels); err != nil {
			log.Error("deserialize cached task list", zap.Error(err))
		} else {
			taskDomains := TaskDomainsFromModels(taskModels)
			return taskDomains, nil
		}
	}

	tasksDomains, err := r.mainRepository.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	tasksModels := TaskDomainsToModels(tasksDomains)
	bytes, err := json.Marshal(tasksModels)
	if err != nil {
		log.Error("serialize cached task list", zap.Error(err))
	}
	if err := r.reddis.HSet(ctx, key, options, bytes).Err(); err != nil {
		log.Error("hset task list", zap.Error(err))
	}
	return tasksDomains, nil

}
