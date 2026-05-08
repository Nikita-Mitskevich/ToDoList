package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Name        string
	Description string
}

func (t *TaskDTO) Validate() error {
	if t.Name == "" {
		return errors.New("Имя не должно быть пустой строкой")
	}
	if t.Description == "" {
		return errors.New("Описание не должно быть пустой строкой")
	}
	return nil
}

type ErrDTO struct {
	Message string
	Time    time.Time
}

func (e *ErrDTO) ToJSON() string {
	textBytes, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	s := string(textBytes)
	return s
}

type CompletedDTO struct {
	Completed *bool
}
