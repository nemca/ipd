package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nemca/ipd/internal/config"
	"github.com/nemca/ipd/internal/handlers"
	"github.com/nemca/ipd/internal/logger"
	"github.com/nemca/ipd/internal/middleware"
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

	logger := logger.NewLogger(cfg.Log.Type, cfg.Log.Level)

	router := mux.NewRouter()

	logMiddleware := middleware.NewLogRequestMiddleware(logger)
	rootHandler := handlers.NewRootHandler(cfg)

	router.Use(logMiddleware.Use)
	router.HandleFunc("/", rootHandler.GetIP).Methods(http.MethodGet)

	logger.Infof("listening on %s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort), router)
	if err != nil {
		logger.Fatal(err)
	}
}
