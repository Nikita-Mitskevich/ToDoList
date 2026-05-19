package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	core_errors "restapi/internal/core/errors"
	core_logger "restapi/internal/core/logger"

	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	zap *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(zap *core_logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{zap: zap, rw: rw}
}

func (h *HTTPResponseHandler) HandlePanic(p any, msg string) {
	err := fmt.Errorf("Unexpected panic: %v", p)
	h.zap.Error(msg, zap.Error(err))
	h.errorResponse(http.StatusInternalServerError, err, msg)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		StatusCode int
		logFunc    func(msg string, zap ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrConflict):
		{
			StatusCode = http.StatusConflict
			logFunc = h.zap.Warn
		}
	case errors.Is(err, core_errors.ErrInvalidArgument):
		{
			StatusCode = http.StatusBadRequest
			logFunc = h.zap.Warn
		}
	case errors.Is(err, core_errors.ErrNotFound):
		{
			StatusCode = http.StatusNotFound
			logFunc = h.zap.Debug
		}
	default:
		{
			StatusCode = http.StatusInternalServerError
			logFunc = h.zap.Error
		}
	}
	logFunc(msg, zap.Error(err))
	h.errorResponse(StatusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	h.rw.WriteHeader(statusCode)
	resp := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}
	h.JSONResponse(resp, statusCode)
}

func (h *HTTPResponseHandler) JSONResponse(responseBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)
	data, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		h.zap.Error("write HTTP response", zap.Error(err))
		return
	}
	h.rw.Write(data)
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}
