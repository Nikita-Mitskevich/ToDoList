package main

import (
	"context"
	"fmt"
	"os/signal"
	core_logger "restapi/internal/core/logger"
	core_pgx_pool "restapi/internal/core/repository/postgres/pool/pgx"
	core_goredisv9_pool "restapi/internal/core/repository/redis/pool/go-redis-v9"
	"restapi/internal/core/transport/http/middleware"
	"restapi/internal/core/transport/http/server"
	statistics_repository "restapi/internal/features/statistics/repository"
	statistics_service "restapi/internal/features/statistics/service"
	statistics_transport_http "restapi/internal/features/statistics/transport/http"
	tasks_postgres_repository "restapi/internal/features/tasks/repository/postgres"
	tasks_reddis_repository "restapi/internal/features/tasks/repository/reddis"
	tasks_service "restapi/internal/features/tasks/service"
	tasks_transport "restapi/internal/features/tasks/transport/http"
	users_postgres_repository "restapi/internal/features/users/repository/postgres"
	users_service "restapi/internal/features/users/service"
	user "restapi/internal/features/users/transport/http"
	web_repository "restapi/internal/features/web/repository/file_system"
	web_service "restapi/internal/features/web/service"
	web_transport "restapi/internal/features/web/transport/http"
	"syscall"
	"time"

	"go.uber.org/zap"

	_ "restapi/docs"
)

var (
	timeZone = time.UTC
)

// @title Golang Todo API
// @version 1.0
// @description Todo application REST-API scheme
// @host 127.0.0.1:5050
// @BasePath /api/v1
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

	log.Debug("initializing redis connection pool")
	cache_pool, err := core_goredisv9_pool.NewRedisv9Pool(ctx, core_goredisv9_pool.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init redis connection pool", zap.Error(err))
	}
	defer cache_pool.Close()

	log.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUserService(usersRepository)
	usersTransport := user.NewUsersHTTPHandler(usersService)

	log.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	reddisRepository := tasks_reddis_repository.NewTasksRepository(cache_pool, tasksRepository)
	tasksService := tasks_service.NewTasksService(reddisRepository)
	tasksTransport := tasks_transport.NewTasksHTTPHandler(tasksService)

	log.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_repository.NewStatisticsReposiory(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransport := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	log.Debug("initializing feature", zap.String("feature", "web"))
	webRepository := web_repository.NewWebRepository()
	webService := web_service.NewWebService(webRepository)
	webTransport := web_transport.NewHTTPWebHandler(webService)

	log.Debug("initializing HTTP server")
	httpServer := server.NewHTTPServer(server.NewConfigMust(), log, middleware.CORS(), middleware.RequestId(), middleware.Logger(log),
		middleware.Trace(), middleware.Panic())

	apiRouter := server.NewAPIVersionRouter(server.ApiVersion1)
	apiRouter.RegisterRoutes(usersTransport.Routes()...)
	apiRouter.RegisterRoutes(tasksTransport.Routes()...)
	apiRouter.RegisterRoutes(statisticsTransport.Routes()...)

	httpServer.RegisterAPIRoutes(apiRouter)
	httpServer.RegisterRoutes(webTransport.GetRoutes()...)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run main: ", zap.Error(err))
	}
}
