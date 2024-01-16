package server

import (
	"net/http"

	handler "github.com/gabrieldillenburg/post-flow/internal/handlers"
)

func InitializeServer() *http.Server {
	http.HandleFunc("/transcribe", handler.YouTubeTranscriptorHandler)

	server := &http.Server{Addr: ":8080"}
	return server
}
