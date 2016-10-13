package handlers

import (
	"io"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	is_alive := `{"is_alive": "true"}`
	io.WriteString(w, is_alive)
}
