package web_transport

import (
	"net/http"
	core_logger "restapi/internal/core/logger"
	core_http_response "restapi/internal/core/transport/http/response"
)

func (h *HTTPWebHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	html, err := h.webService.GetFile()
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get html file")
	}

	responseHandler.HTMLResponse(html)
}
