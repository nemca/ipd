package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/nemca/ipd/internal/config"
)

var (
	cfg     *config.Config
	err     error
	version string = "unknown"
	build   string = "unknown"
)

func main() {
	cfg, err = config.Init(version, build)
	if err != nil {
		log.Fatalf("error occurred while reading config: %s\n", err.Error())
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	log.Printf("listening on %s:%s\n", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HTTP.ListenAddress, cfg.HTTP.ListenPort), logRequest(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	forwarderFor := r.Header.Get(cfg.HTTP.ForwardedHeader)
	if forwarderFor != "" {
		fmt.Fprintf(w, forwarderFor)
		return
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, host)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
