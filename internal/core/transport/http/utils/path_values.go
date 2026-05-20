package core_http_utils

import (
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"
	"strconv"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key=%s in path values: %w", key, core_errors.ErrInvalidArgument)
	}
	value, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("path value=%s by key %s not a valid integer: %v: %w", pathValue, key, err, core_errors.ErrInvalidArgument)
	}
	return value, nil
}
