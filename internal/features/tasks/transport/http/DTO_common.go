package tasks_transport

import "time"

type TaskDTO struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Name        string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorId int `json:"author_user_id"`
}
