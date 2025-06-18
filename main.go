package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/FriendlyJJC/api_server/apiv1"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", http.HandlerFunc(apiv1.APIHandleV1)))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("Something went terribly wrong", "error", err)
	}
}
