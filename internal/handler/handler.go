package handler

import (
	"fmt"
	"net/http"
)

var _ http.Handler = &PingHandler{}

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping returns a Pong!
func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	h.serve(w, r)
}

func (h *PingHandler) serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pong!")
}
