package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1></h1><h4>serving: <code>%s</code></h4>", r.URL.Path)
	slog.Info("served", slog.String("host", r.Host))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)

	Body := "The current time is:"

	fmt.Fprintf(w, "<h1 align='center'>%s</h1>", Body)
	fmt.Fprintf(w, "<h3 align='center'>%s</h3><br/>", t)
	fmt.Fprintf(w, "<h4>serving: <code>%s</code><br/>", r.URL.Path)
	slog.Info("served", slog.String("time", r.Host))
}

func main() {
	var cr customRouter

	server := &http.Server{
		Addr:           net.JoinHostPort("0.0.0.0", "3000"),
		Handler:        cr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	log.Fatal(server.ListenAndServe())
}

var routes []string

type customRouter struct{}

func (customRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	slog.Info("served", slog.String("path", r.URL.Path))
}
