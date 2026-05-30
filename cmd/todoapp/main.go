package main

import (
	"context"
	"fmt"
	"os/signal"
	core_logger "restapi/internal/core/logger"
	core_pgx_pool "restapi/internal/core/repository/postgres/pool/pgx"
	"restapi/internal/core/transport/http/middleware"
	"restapi/internal/core/transport/http/server"
	statistics_repository "restapi/internal/features/statistics/repository"
	statistics_service "restapi/internal/features/statistics/service"
	statistics_transport_http "restapi/internal/features/statistics/transport/http"
	tasks_postgres_repository "restapi/internal/features/tasks/repository/postgres"
	tasks_service "restapi/internal/features/tasks/service"
	tasks_transport "restapi/internal/features/tasks/transport/http"
	users_postgres_repository "restapi/internal/features/users/repository/postgres"
	users_service "restapi/internal/features/users/service"
	user "restapi/internal/features/users/transport/http"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	log, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("не удалось получить логгер")
	}
	defer log.Close()

	log.Debug("application time zone", zap.Any("zone", timeZone))

	log.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewConnectionPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	log.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUserService(usersRepository)
	usersTransport := user.NewUsersHTTPHandler(usersService)

	log.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransport := tasks_transport.NewTasksHTTPHandler(tasksService)

	log.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_repository.NewStatisticsReposiory(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransport := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	log.Debug("initializing HTTP server")
	httpServer := server.NewHTTPServer(server.NewConfigMust(), log, middleware.RequestId(), middleware.Logger(log),
		middleware.Trace(), middleware.Panic())

	apiRouter := server.NewAPIVersionRouter(server.ApiVersion1)
	apiRouter.RegisterRoutes(usersTransport.Routes()...)
	apiRouter.RegisterRoutes(tasksTransport.Routes()...)
	apiRouter.RegisterRoutes(statisticsTransport.Routes()...)

	httpServer.RegisterAPIRoutes(apiRouter)
	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run main: ", zap.Error(err))
	}
}
