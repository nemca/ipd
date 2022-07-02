package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	forwarderForHeader string = "X-Forwarder-For"
	listenPort         int    = 8080
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	log.Printf("listening on %v\n", listenPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), logRequest(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	forwarderFor := r.Header.Get(forwarderForHeader)
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
