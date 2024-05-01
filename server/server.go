package server

import (
	"fmt"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "la")
}
