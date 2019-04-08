package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 8080, "Port to serve the app")

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, fmt.Sprintf("static%s", r.URL.Path))
	}
}

func mailingListHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/mailing-success.html")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	fmt.Printf("Running on port %d...", *port)

	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/mailing", mailingListHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
