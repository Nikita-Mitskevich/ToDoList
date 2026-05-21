package domain

import "time"

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
	return NewTask(UninitializedID, UninitializedVersion, name, description, false, time.Now(), &time.Time{}, authorId)
}


