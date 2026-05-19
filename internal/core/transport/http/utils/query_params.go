package core_http_utils

import (
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"
	"strconv"
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
