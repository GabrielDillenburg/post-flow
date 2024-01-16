package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendToSummarizationAPI(transcriptions []TranscriptionResult) (string, error) {
	// Serialize the transcriptions to JSON
	jsonData, err := json.Marshal(transcriptions)
	if err != nil {
		return "", fmt.Errorf("error marshalling transcriptions: %w", err)
	}

	// Prepare the HTTP request
	req, err := http.NewRequest("POST", "http://summarization.api/endpoint", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to summarization API: %w", err)
	}
	defer resp.Body.Close()

	// Read and return the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(responseBody), nil
}
