package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ChatGPTRequest struct {
	Prompt string `json:"prompt"`
	// Add other fields as required by ChatGPT API, e.g., temperature, max tokens, etc.
}

func SendToSummarizationAPI(transcriptions []TranscriptionResult) (string, error) {
	// Construct the prompt for ChatGPT
	prompt := "Please summarize the following transcriptions:\n"
	for _, t := range transcriptions {
		prompt += fmt.Sprintf("Video ID %s: %s\n", t.VideoID, t.Transcription)
	}

	// Prepare ChatGPT request payload
	gptRequest := ChatGPTRequest{
		Prompt: prompt,
		// Set other fields if necessary
	}

	jsonData, err := json.Marshal(gptRequest)
	if err != nil {
		return "", fmt.Errorf("error marshalling request data: %w", err)
	}

	// Prepare the HTTP request to ChatGPT API
	chatGPTURL := os.Getenv("CHAT_GPT_API_URL") // Ensure this environment variable is set
	if chatGPTURL == "" {
		return "", fmt.Errorf("CHAT_GPT_API_URL environment variable is not set")
	}

	req, err := http.NewRequest("POST", chatGPTURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Include necessary headers, such as Authorization if required
	req.Header.Set("Content-Type", "application/json")
	// Example: req.Header.Set("Authorization", "Bearer YOUR_TOKEN_HERE")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to ChatGPT API: %w", err)
	}
	defer resp.Body.Close()

	// Read and return the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(responseBody), nil
}
