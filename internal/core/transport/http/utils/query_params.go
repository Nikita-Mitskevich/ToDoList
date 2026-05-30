package core_http_utils

import (
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"
	"strconv"
	"time"
)

func GetQueryParamInt(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("param=%s by key=%s not a valid integer: %v: %w", param, key, err, core_errors.ErrInvalidArgument)
	}

	return &val, nil
}

func GetQueryParamDate(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	example := "2006-01-02"
	val, err := time.Parse(example, param)
	if err != nil {
		return nil, fmt.Errorf("param=%s by key=%s not a valid date: %v: %w", param, key, err, core_errors.ErrInvalidArgument)
	}

	return &val, nil
}
