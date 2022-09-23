package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/nemca/ipd/internal/config"
	"github.com/nemca/ipd/internal/handlers"
	"github.com/nemca/ipd/internal/logger"
	"github.com/nemca/ipd/internal/middleware"
	"github.com/spf13/viper"
)

var (
	version string = "unknown"
	build   string = "unknown"
)

func main() {
	cfg, err := config.Init(version, build)
	if err != nil {
		log.Fatalf("error occurred while reading config: %s\n", err.Error())
	}
	// if version flag is defined shows info and exit
	if viper.GetBool("version") {
		fmt.Printf("ipd version %s\ncommit:\t%s\n", version, build)
		return
	}

	logger := logger.NewLogger(cfg.Log.Type, cfg.Log.Level)

	router := mux.NewRouter()

	logMiddleware := middleware.NewLogRequestMiddleware(logger)
	rootHandler := handlers.NewRootHandler(cfg)

	router.Use(logMiddleware.Use)
	router.HandleFunc("/", rootHandler.GetIP).Methods(http.MethodGet)

	logger.Infof("listening on %s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort),
		Handler: router,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Errorf("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-idleConnectionsClosed

	logger.Info("stopped")
}
