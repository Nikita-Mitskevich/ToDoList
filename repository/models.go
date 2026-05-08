package repository

import "time"

type Todolist_model struct {
	Name        string
	IsCompleted bool
	Description string
	StartedAt   time.Time
	EndedAt     *time.Time
}
