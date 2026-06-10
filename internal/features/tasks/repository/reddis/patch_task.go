package tasks_reddis_repository

import (
	"context"
	"restapi/internal/core/domain"
)

func (r *TasksRepository) PatchTask(ctx context.Context, taskId int, taskPatch domain.Task) (domain.Task, error) {

	taskDomain, err := r.mainRepository.PatchTask(ctx, taskId, taskPatch)
	if err != nil {
		return domain.Task{}, err
	}
	key := GetKeyValue(taskId)
	r.CacheInfo(ctx, key, taskDomain)
	r.DeleteCache(ctx, taskDomain.AuthorId, nil)
	return taskDomain, nil
}
