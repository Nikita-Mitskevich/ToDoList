package domain

import (
	"fmt"
	core_errors "restapi/internal/core/errors"
	"time"
)

type Task struct {
	ID      int
	Version int

	Name        string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorId int
}

func NewTask(id int, version int, name string,
	description *string, completed bool,
	createdAt time.Time, completedAt *time.Time,
	authorId int) Task {
	return Task{ID: id, Version: version,
		Name: name, Description: description,
		Completed: completed, CreatedAt: createdAt,
		CompletedAt: completedAt, AuthorId: authorId}
}

func NewTaskUninitialized(name string, description *string, authorId int) Task {
	return NewTask(UninitializedID, UninitializedVersion, name, description, false, time.Now(), nil, authorId)
}

func (t *Task) Validate() error {
	if len([]rune(t.Name)) < 1 || len([]rune(t.Name)) > 100 {
		return fmt.Errorf("task name must be between 1 and 100: %w", core_errors.ErrInvalidArgument)
	}

	if t.Description != nil {
		if len([]rune(*t.Description)) < 1 || len([]rune(*t.Description)) > 100 {
			return fmt.Errorf("task name must be between 1 and 100: %w", core_errors.ErrInvalidArgument)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("time completed cant be null if task is completed: %w", core_errors.ErrInvalidArgument)
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("time completed cant be later then time created: %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("time completed must be null if completed is false: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type TaskPatch struct {
	Name        Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func (p *TaskPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf("Title cant be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("Completed value cant be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (t *Task) ApplyPatch(taskPatch TaskPatch) error {
	if err := taskPatch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", core_errors.ErrInvalidArgument)
	}
	temp := *t
	if taskPatch.Name.Set {
		temp.Name = *taskPatch.Name.Value
	}
	if taskPatch.Description.Set {
		temp.Description = taskPatch.Description.Value
	}
	if taskPatch.Completed.Set {
		temp.Completed = *taskPatch.Completed.Value
		if !*taskPatch.Completed.Value {
			temp.CompletedAt = nil
		} else {
			now := time.Now()
			temp.CompletedAt = &now
		}
	}
	if err := temp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}
	*t = temp
	return nil
}

func (t Task) CompletionDuration() *time.Duration {
	if t.CompletedAt == nil || !t.Completed || t.CompletedAt.Before(t.CreatedAt) {
		return nil
	}
	time := t.CompletedAt.Sub(t.CreatedAt)
	return &time
}
