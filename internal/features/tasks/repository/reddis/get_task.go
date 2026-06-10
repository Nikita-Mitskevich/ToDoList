package tasks_reddis_repository

import (
	"context"
	"restapi/internal/core/domain"
)

func (r *TasksRepository) GetTask(ctx context.Context, taskId int) (domain.Task, error) {
	key := GetKeyValue(taskId)
	taskDomain, ok := r.GetFromCache(ctx, key)
	if ok {
		return taskDomain, nil
	}

	taskDomain, err := r.mainRepository.GetTask(ctx, taskId)
	if err != nil {
		return domain.Task{}, err
	}

	r.CacheInfo(ctx, key, taskDomain)
	return taskDomain, nil
}
