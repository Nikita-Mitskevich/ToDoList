package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{httpHandlers: httpHandler}
}

func (s *HTTPServer) StartServer() error {
	mux := mux.NewRouter()
	mux.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)

	mux.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	mux.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleMarkAsCompletedTask)

	mux.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.httpHandlers.HandleGetAllNotCompletedTask)

	mux.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTask)

	mux.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)

	if err := http.ListenAndServe(":9091", mux); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return err
		}
	}
	return nil

}
