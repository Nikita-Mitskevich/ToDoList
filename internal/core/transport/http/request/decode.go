package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"

	"github.com/go-playground/validator/v10"
)

type validatable interface {
	Validate() error
}

var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	v, ok := dest.(validatable)

	if ok {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("validate json: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	} else {
		if err := requestValidator.Struct(dest); err != nil {
			return fmt.Errorf("validate json: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}
	return nil
}
