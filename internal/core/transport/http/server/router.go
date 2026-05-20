package server

import (
	"fmt"
	"net/http"
	"restapi/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("/v1")
	ApiVersion2 = ApiVersion("/v2")
	ApiVersion3 = ApiVersion("/v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []middleware.Middleware
}

func NewAPIVersionRouter(v ApiVersion, middleware ...middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{apiVersion: v, ServeMux: http.NewServeMux(), middleware: middleware}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.WithMiddleware())
	}
}


func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return middleware.ChainMiddleware(r, r.middleware...)
}