package core_redis_pool

import "errors"

var ErrNotFound = errors.New("В кеше не нашлось ни одной подходящей записи")
