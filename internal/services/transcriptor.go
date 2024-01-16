package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func TranscriptionAPI(audioFilePath string) (string, error) {
	// Read the audio file
	audioData, err := os.ReadFile(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading audio file: %w", err)
	}

	apiURL := os.Getenv("TRANSCRIPTION_API_URL")
	bearerToken := os.Getenv("TRANSCRIPTION_API_TOKEN")

	if apiURL == "" || bearerToken == "" {
		return "", fmt.Errorf("environment variables for API URL and/or Token are not set")
	}

	// Append the max_new_tokens query parameter
	maxNewTokens := "1000"
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return "", fmt.Errorf("error parsing API URL: %w", err)
	}
	query := parsedURL.Query()
	query.Set("max_new_tokens", maxNewTokens)
	parsedURL.RawQuery = query.Encode()
	finalURL := parsedURL.String()

	// Create a request to the transcription API
	req, err := http.NewRequest("POST", finalURL, bytes.NewBuffer(audioData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "audio/mpeg")
	req.Header.Set("Authorization", "Bearer"+bearerToken)

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
