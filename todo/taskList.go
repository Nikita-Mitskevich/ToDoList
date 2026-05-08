package todo

import (
	"errors"
	"restapi/errors_tec"
	"restapi/repository"
	"sync"
	"time"
)

type Task struct {
	Name        string
	IsCompleted bool
	Description string
	StartedAt   time.Time
	EndedAt     time.Time
}

type Tasks struct {
	mtx      sync.Mutex
	database *repository.Database_handler
}

func CreateNewList() (*Tasks, error) {
	d, err := repository.Create_database_handler()
	if err != nil {
		return nil, err
	}
	return &Tasks{database: d}, nil
}

func CreateNewTask(name string, description string) *Task {
	return &Task{
		Name:        name,
		Description: description,
		IsCompleted: false,
		StartedAt:   time.Now(),
		EndedAt:     time.Time{},
	}
}

func (t *Tasks) CreateTask(task Task) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if err, _ := t.database.FoundByName_database(task.Name); !errors.Is(errors_tec.ErrTaskDoesntExist, err) {
		return err
	}
	err := t.database.CreateNewTask_database(repository.Todolist_model{Name: task.Name, IsCompleted: task.IsCompleted, Description: task.Description, StartedAt: task.StartedAt, EndedAt: &task.EndedAt})
	return err
}

func (t *Tasks) DeleteTask(name string) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if err, _ := t.database.FoundByName_database(name); err != nil {
		return err
	}
	err := t.database.DeleteByName_database(name)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) MarkTask(name string, status bool) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	err, task := t.database.FoundByName_database(name)
	if err != nil {
		return err
	}
	task.IsCompleted = status
	if status {
		now := time.Now()
		task.EndedAt = &now
	} else {
		task.EndedAt = nil
	}
	err = t.database.MarkTask_database(task)
	return err
}

func (t *Tasks) GetAllNotCompleted() []repository.Todolist_model {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	tmp, _ := t.database.GetAllNotCompleted_database()
	return tmp
}

func (t *Tasks) GetAll() []repository.Todolist_model {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	tmp, _ := t.database.GetAll_database()
	return tmp
}

func (t *Tasks) GetTask(name string) (*Task, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	err, task := t.database.FoundByName_database(name)
	if err != nil {
		return nil, err
	}

	var endedAt time.Time
	if task.EndedAt != nil {
		endedAt = *task.EndedAt
	}
	retTask := Task{Name: task.Name, IsCompleted: task.IsCompleted, Description: task.Description, StartedAt: task.StartedAt, EndedAt: endedAt}
	return &retTask, nil
}
