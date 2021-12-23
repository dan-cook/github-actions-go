package main

import (
	"log"
	"net/http"

	"github-actions-go/internal/handler"
)

func main() {
	pingHandler := handler.NewPingHandler()

	mux := http.NewServeMux()
	mux.Handle("/ping", pingHandler)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}
