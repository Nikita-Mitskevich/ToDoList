package server

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRoute(Method string, Path string,
	Handler http.HandlerFunc) Route {
	return Route{Method: Method,
		Path:    Path,
		Handler: Handler}
}

//6.10.55
