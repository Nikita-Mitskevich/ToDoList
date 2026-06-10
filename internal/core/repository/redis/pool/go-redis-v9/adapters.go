package core_goredisv9_pool

import (
	"errors"
	core_redis_pool "restapi/internal/core/repository/redis/pool"

	"github.com/redis/go-redis/v9"
)

type StringCmd struct {
	*redis.StringCmd
}

type IntCmd struct {
	*redis.IntCmd
}

type StatusCmd struct {
	*redis.StatusCmd
}



func (s StringCmd) Bytes() ([]byte, error) {
	data, err := s.StringCmd.Bytes()
	if err != nil {
		return nil, mapError(err)
	}

	return data, nil
}

func mapError(err error) error {
	if errors.Is(err, redis.Nil) {
		return core_redis_pool.ErrNotFound
	}

	return err
}
