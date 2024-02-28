package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	helper "github.com/gabrieldillenburg/post-flow/internal/helpers/logs"
)

func TranscriptionAPI(audioFilePath string) (string, error) {
	// Read the audio file data
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening audio file: %w", err)
	}
	defer audioFile.Close()

	audioData, err := io.ReadAll(audioFile)
	if err != nil {
		return "", fmt.Errorf("error reading audio data: %w", err)
	}

	// Prepare API details
	apiURL := os.Getenv("TRANSCRIPTION_API_URL")
	bearerToken := os.Getenv("TRANSCRIPTION_API_TOKEN")
	model := os.Getenv("TRANSCRIPTOR_API_MODEL")

	if apiURL == "" || bearerToken == "" || model == "" {
		return "", fmt.Errorf("environment variables for API URL, Token, or Model are not set")
	}

	// Create a multipart writer
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Add "file" part with audio data
	filename := filepath.Base(audioFilePath)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("error creating file part: %w", err)
	}
	_, err = part.Write(audioData)
	if err != nil {
		return "", fmt.Errorf("error writing audio data to part: %w", err)
	}

	// Add "model" part with retrieved model value
	part, err = writer.CreateFormField("model")
	if err != nil {
		return "", fmt.Errorf("error creating model part: %w", err)
	}
	_, err = part.Write([]byte(model))
	if err != nil {
		return "", fmt.Errorf("error writing model value to part: %w", err)
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("error closing multipart writer: %w", err)
	}

	// Create a request with multipart format
	req, err := http.NewRequest("POST", apiURL, &buffer)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set the content type header with boundary
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Add("Accept", "application/json")

	helper.PrintMultipartRequest(req)
	helper.PrintRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to transcription API: %w", err)
	}
	defer resp.Body.Close()

	// Log the response for debugging
	// to do: create an aspect strucuture to handle logs globally
	if resp.StatusCode != http.StatusOK {
		helper.PrintResponse(resp)
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
