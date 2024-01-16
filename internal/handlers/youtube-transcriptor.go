package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrieldillenburg/post-flow/internal/services"
)

func YouTubeTranscriptorHandler(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var query struct {
		SearchTerm string `json:"searchTerm"`
	}
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	videos, err := services.GetYouTubeVideos(services.Query(query.SearchTerm))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching videos: %v", err), http.StatusInternalServerError)
		return
	}

	// Process videos and get the summary
	summary, err := services.ProcessVideosWorkflow(videos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing videos: %v", err), http.StatusInternalServerError)
		return
	}

	// Send the summary as the response
	fmt.Fprintf(w, "Summary: %s", summary)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
