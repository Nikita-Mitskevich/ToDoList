package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"restapi/internal/core/domain"
	"restapi/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statisticsService StatisticsService,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
