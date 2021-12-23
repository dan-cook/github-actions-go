package handler

import (
	"fmt"
	"net/http"
)

var _ http.Handler = &PingHandler{}

type PingHandler struct{}

// Ping returns a Pong!
func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Pong!")
}
