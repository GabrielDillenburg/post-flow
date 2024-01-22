package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func TranscriptionAPI(audioFilePath string) (string, error) {
	// Read the audio file into memory
	audioData, err := os.ReadFile(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading audio file: %w", err)
	}

	apiURL := os.Getenv("TRANSCRIPTION_API_URL")
	bearerToken := os.Getenv("TRANSCRIPTION_API_TOKEN")

	if apiURL == "" || bearerToken == "" {
		return "", fmt.Errorf("environment variables for API URL and/or Token are not set")
	}

	// Append query parameter to the API URL (if necessary)
	// apiURL += "/?max_new_tokens=1000"
	payload := bytes.NewReader(audioData)

	// Create a request to the transcription API
	req, err := http.NewRequest("POST", apiURL, payload)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Content-Type", "audio/flac")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Add("Accept", "application/json")

	// Log the request for debugging
	// to do: create an aspect strucuture to handle logs globally
	log.Printf("Sending request to API URL: %s", apiURL)
	log.Printf("Request headers: %+v", req.Header)
	log.Printf("Size of audio file payload: %d bytes", len(audioData))
	log.Printf("First 20 bytes of audio data: %x", audioData[:20])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to transcription API: %w", err)
	}
	defer resp.Body.Close()

	// Log the response for debugging
	// to do: create an aspect strucuture to handle logs globally
	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf("Transcription API returned non-OK status: %s, Response Body: %s", resp.Status, string(responseBody))
		return "", fmt.Errorf("transcription API returned non-OK status: %s", resp.Status)
	}

	// Read and return the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
