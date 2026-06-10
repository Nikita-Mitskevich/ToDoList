package tasks_reddis_repository

import "context"

func (r *TasksRepository) DeleteTask(ctx context.Context, taskId int) error {
	key := GetKeyValue(taskId)
	taskDomain, ok := r.GetFromCache(ctx, key)
	if !ok {
		var err error
		taskDomain, err = r.mainRepository.GetTask(ctx, taskId)
		if err != nil {
			return err
		}
	}
	r.DeleteCache(ctx, taskDomain.AuthorId, &taskId)
	return r.mainRepository.DeleteTask(ctx, taskId)
}
