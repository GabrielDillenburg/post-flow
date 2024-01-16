package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func TranscriptionAPI(audioFilePath string) (string, error) {
	// Read the audio file
	audioData, err := os.ReadFile(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading audio file: %w", err)
	}

	// Create a request to the transcription API
	req, err := http.NewRequest("POST", "http://transcription.api/endpoint", bytes.NewBuffer(audioData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "audio/mpeg") // Set the appropriate content type

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to transcription API: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil

}
