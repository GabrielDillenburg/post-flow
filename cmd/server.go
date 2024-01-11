package server

import (
	"fmt"
	"net/http"
)

// InitializeServer sets up and returns an HTTP server.
func InitializeServer() *http.Server {
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		// extract user input from the request and pass it to the ChatGPT adapter.

		fmt.Fprintf(w, "Generated Response: %s", "response")
	})

	server := &http.Server{Addr: ":8080"}
	return server
}
