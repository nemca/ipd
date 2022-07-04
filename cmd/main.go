package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nemca/ipd/internal/config"
	"github.com/nemca/ipd/internal/handlers"
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

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()
	router.Use(middleware.NewLogRequestMiddleware().Use)

	rootHandler := handlers.NewRootHandler(cfg)

	router.HandleFunc("/", rootHandler.GetIP).Methods(http.MethodGet)

	log.Printf("listening on %s:%s\n", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort), router)
	if err != nil {
		log.Fatal(err)
	}
}
