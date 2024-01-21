package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func TranscriptionAPI(audioFilePath string) (string, error) {
	// Open the audio file
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening audio file: %w", err)
	}
	defer audioFile.Close()

	apiURL := os.Getenv("TRANSCRIPTION_API_URL")
	bearerToken := os.Getenv("TRANSCRIPTION_API_TOKEN")

	if apiURL == "" || bearerToken == "" {
		return "", fmt.Errorf("environment variables for API URL and/or Token are not set")
	}

	// Parse and prepare API URL
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return "", fmt.Errorf("error parsing API URL: %w", err)
	}
	query := parsedURL.Query()
	query.Set("max_new_tokens", "1000") // Adjust parameter as needed
	parsedURL.RawQuery = query.Encode()

	// Create a request to the transcription API
	req, err := http.NewRequest("POST", parsedURL.String(), audioFile)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "audio/mpeg")
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to transcription API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("transcription API returned non-OK status: %s", resp.Status)
	}

	// Read and return the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
