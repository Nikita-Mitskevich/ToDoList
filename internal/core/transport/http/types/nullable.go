package core_http_types

import (
	"encoding/json"
	"restapi/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true
	if string(b) == "null" {
		n.Value = nil
		return nil
	}

	var value T
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	n.Value = &value
	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{Value: n.Value, Set: n.Set}
}
