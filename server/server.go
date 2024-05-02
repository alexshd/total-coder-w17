package server

import (
	"fmt"
	"net/http"
)

// countHandler for /count route, exepts POST only.
func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "fetching data from sourdce")
	default:
		// TODO: make it move automaticly.
		http.Redirect(w, r, "/count/doc", http.StatusMovedPermanently)

	}
}

func Run() error {
	server := http.NewServeMux()

	server.HandleFunc("/count", countHandler)
	server.HandleFunc("/count/doc", countDocHandler)
	server.HandleFunc("/", mainHandler)

	return http.ListenAndServe(":5000", server)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello to All")
}

func countDocHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I welcome you to the amazing Documentation Page :)")
}
