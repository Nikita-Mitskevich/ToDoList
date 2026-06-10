package tasks_reddis_repository

import (
	"context"
	"encoding/json"
	"restapi/internal/core/domain"
	core_logger "restapi/internal/core/logger"

	"go.uber.org/zap"
)

func (r *TasksRepository) GetFromCache(ctx context.Context, key string) (domain.Task, bool) {
	log := core_logger.FromContext(ctx)
	val, err := r.reddis.Get(ctx, key).Bytes()
	if err != nil {
		return domain.Task{}, false
	}
	var taskModel TaskModel
	err = json.Unmarshal(val, &taskModel)
	if err != nil {
		log.Error("unmarshal cache", zap.Error(err))
		return domain.Task{}, false
	}

	taskDomain := TaskDomainFromModel(taskModel)

	return taskDomain, true
}

func (r *TasksRepository) CacheInfo(ctx context.Context, key string, task domain.Task) {
	log := core_logger.FromContext(ctx)
	taskModel := TaskDomainToModel(task)
	bytes, err := json.Marshal(taskModel)
	if err != nil {
		log.Error("marshal task model", zap.Error(err))
	}
	if err := r.reddis.Set(ctx, key, bytes, 0).Err(); err != nil {
		log.Error("", zap.Error(err))
	}
}

func (r *TasksRepository) DeleteCache(ctx context.Context, userId int, taskId *int) {
	log := core_logger.FromContext(ctx)
	toDelete := []string{
		GetTasksKeyValue(nil),
		GetTasksKeyValue(&userId),
	}
	if taskId != nil {
		toDelete = append(toDelete, GetKeyValue(*taskId))
	}

	if err := r.reddis.Del(ctx, toDelete...).Err(); err != nil {
		log.Error("delele cache", zap.Error(err))
	}

}
