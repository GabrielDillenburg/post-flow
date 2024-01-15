package handlers

import (
	"net/http"
)

func YouTubeTranscriptorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		http.Error(w, "GET method is not Allowed", http.StatusMethodNotAllowed)
		return

	default:
		http.Error(w, "Method is not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
