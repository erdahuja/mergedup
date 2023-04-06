package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	"mergedup/app/services/app-api/handlers"
	"mergedup/app/services/debug-api"
	"mergedup/foundation/logger"

	"mergedup/business/sys/database"
	"mergedup/foundation/config"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	_ "github.com/spf13/viper/remote"
)

const build = "dev"

// @title Cart Service
// @version 1.0
// @description This is a generic cart service with RBAC.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /v1
func main() {
	log, err := logger.New("CART-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	opt := maxprocs.Logger(log.Infof)
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}

	fmt.Printf("Go version: %s\n", runtime.Version())

	cfg, err := config.LoadConfig(build, ".")
	if err != nil {
		return err
	}

	// =========================================================================
	// Start Database

	log.Infow("startup", "status", "initializing database support", "host", cfg.DBHost)
	db, err := setupDB(cfg)
	if err != nil {
		log.Errorw("main: unable to connect to database")
		return err
	}

	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", cfg.DBHost)
		db.Close()
	}()

	log.Infow("main:", "config", cfg)

	// =========================================================================
	// Start Debug Service

	log.Infow("startup", "status", "debug v1 router started", "host", cfg.DebugHost)

	go func() {
		if err := http.ListenAndServe(cfg.DebugHost, debug.Mux(build, log, db)); err != nil {
			log.Fatal("shutdown", "status", "debug v1 router closed", cfg.DebugHost, "ERROR", err)
		}
	}()

	// =========================================================================
	// Start API Service

	log.Infow("startup", "status", "initializing V1 API support")

	return runAPI(cfg, db, log)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "    ", "\t")
	return string(s)
}

func setupDB(cfg config.Configurations) (*sqlx.DB, error) {
	db, err := database.Open(cfg.GetDBConfig())
	if err != nil {
		return db, errors.Wrap(err, "connecting to db")
	}

	return db, nil
}

func runAPI(cfg config.Configurations, db *sqlx.DB, log *zap.SugaredLogger) error {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown: shutdown,
		Log:      log,
		DB:       db,
	})

	api := http.Server{
		Addr:         cfg.APIHost,
		Handler:      apiMux,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Infow("startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	return waitForCompletion(serverErrors, shutdown, &api, time.Duration(cfg.ShutdownTimeout)*time.Second)
}

func waitForCompletion(serverErrors chan error, shutdown chan os.Signal, api *http.Server, timeout time.Duration) error {
	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %+v : Start shutdown", sig)
		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}
	return nil
}
