package tasks_reddis_repository

import (
	"context"
	"restapi/internal/core/domain"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	taskDomain, err := r.mainRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}
	key := GetKeyValue(taskDomain.ID)
	r.CacheInfo(ctx, key, taskDomain)
	r.DeleteCache(ctx, task.AuthorId, nil)
	return taskDomain, nil
}
