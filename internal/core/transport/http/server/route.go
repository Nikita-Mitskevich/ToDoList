package server

import (
	"net/http"
	"restapi/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []middleware.Middleware
}

func NewRoute(Method string, Path string,
	Handler http.HandlerFunc) Route {
	return Route{Method: Method,
		Path:    Path,
		Handler: Handler}
}

func (r *Route) WithMiddleware() http.Handler {
	return middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
