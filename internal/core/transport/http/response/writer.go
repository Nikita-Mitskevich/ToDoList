package core_http_response

import "net/http"

var (
	StartCode int = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func CreateResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, statusCode: StartCode}
}

func (h *ResponseWriter) WriteHeader(statusCode int) {
	h.ResponseWriter.WriteHeader(statusCode)
	h.statusCode = statusCode
}

func (h *ResponseWriter) GetStatusCode() int {
	if h.statusCode == StartCode {
		panic("no status code set")
	}
	return h.statusCode
}
