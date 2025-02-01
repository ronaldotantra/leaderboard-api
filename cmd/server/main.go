package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/ronaldotantra/leaderboard-api/config"
	"github.com/ronaldotantra/leaderboard-api/internal/app"
	"github.com/ronaldotantra/leaderboard-api/internal/logger"
	"github.com/ronaldotantra/leaderboard-api/internal/logger/logrus"
)

func main() {
	var err error
	godotenv.Load(".env")
	config.Init()

	level := logger.Debug
	if config.IsProductionEnvironment() {
		level = logger.Info
	}

	// Initialize of Logger
	logConfig := &logger.Configuration{
		ConsoleJSONFormat: false,
		ConsoleLevel:      level,
	}

	logger.SetRepository(logrus.NewLogrusLogger(logConfig))

	if config.SentryDSN == "" {
		logger.Infof("sentry dsn empty, skipping sentry initialization")
	} else {
		err = sentry.Init(sentry.ClientOptions{
			Dsn:         config.SentryDSN,
			Environment: config.Environment,
		})
		if err != nil {
			logger.Errorf("error initializing sentry %v\n", err)
		}
		logger.Infof("sentry initialized to %s\n", config.SentryDSN)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var appServer *http.Server
	go func() {
		appContainer, err := app.SetupApp(context.Background(), app.ApplicationParams{
			StorageParams: app.StoragesParams{
				DBConnString: config.DatabaseConnectionString,
			},
		})
		if err != nil {
			panic(fmt.Sprintf("error setting up storages - %v\n", err))
		}
		strg := appContainer.Storages
		services := appContainer.Services
		handlers := setupHandlers(strg, services)
		appServer = &http.Server{
			Addr:    fmt.Sprintf(":%s", config.Port),
			Handler: buildAPIRoutes(services, handlers),
		}

		logger.Infof("starting app server on %s\n", appServer.Addr)
		if err := appServer.ListenAndServe(); err != nil {
			logger.Errorf("error listening %v\n", err)
		}
	}()

	<-done
	shutdownServers(context.Background(), appServer)
}

func shutdownServers(ctx context.Context, servers ...*http.Server) {
	wg := &sync.WaitGroup{}
	for _, server := range servers {
		wg.Add(1)
		go func(c context.Context, srv *http.Server) {
			defer wg.Done()
			ctxTimeout, cancel := context.WithTimeout(c, 15*time.Second)
			if err := srv.Shutdown(ctxTimeout); err != nil {
				logger.Errorf("error shutting down server - %v \n", err)
			}
			cancel()
		}(ctx, server)
	}
	wg.Wait()
	logger.Infof("servers shut down gracefully")
}
