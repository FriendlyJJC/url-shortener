package main

import (
	"net/http"

	"github.com/FriendlyJJC/api_server/apiv1"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", http.HandlerFunc(apiv1.APIHandleV1)))

	http.ListenAndServe(":8080", mux)
}
