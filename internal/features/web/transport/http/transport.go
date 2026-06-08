package web_transport

import (
	"restapi/internal/core/transport/http/server"
)

type HTTPWebHandler struct {
	webService WebService
}

type WebService interface {
	GetFile() ([]byte, error)
}

func NewHTTPWebHandler(w WebService) *HTTPWebHandler {
	return &HTTPWebHandler{webService: w}
}

func (w *HTTPWebHandler) GetRoutes() []server.Route {
	return []server.Route{
		{
			Path:    "/",
			Handler: w.GetFile,
		},
	}
}
