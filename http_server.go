package main

import (
	"fmt"
	"net/http"
)

// InitializeServer sets up and returns an HTTP server.
func InitializeServer() *http.Server {
	chatAdapter := NewChatGPTAdapter()

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		// extract user input from the request and pass it to the ChatGPT adapter.
		response, err := chatAdapter.GenerateResponse("User's input")
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Generated Response: %s", response)
	})

	server := &http.Server{Addr: ":8080"}
	return server
}
