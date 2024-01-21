package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ChatGPTRequest struct {
	Prompt string `json:"prompt"`
	// Add other fields as required by ChatGPT API, e.g., temperature, max tokens, etc.
}

func SendToSummarizationAPI(transcriptions []TranscriptionResult) (string, error) {
	/* Construct the prompt for ChatGPT
	Todo: see if is the correct structure, becase we need only one instruction,
	one request to the gpt API with all video transcriptions.
	*/
	prompt := "Please summarize the following transcriptions:\n"
	for _, t := range transcriptions {
		prompt += fmt.Sprintf("Video ID %s: %s\n", t.VideoID, t.Transcription)
	}

	gptRequest := ChatGPTRequest{
		Prompt: prompt,
		// Set other fields if necessary. Todo: see the openAI docs
	}

	jsonData, err := json.Marshal(gptRequest)
	if err != nil {
		return "", fmt.Errorf("error marshalling request data: %w", err)
	}

	chatGPTURL := os.Getenv("CHAT_GPT_API_URL")
	gptApiToken := os.Getenv("GPT_API_TOKEN")

	if chatGPTURL == "" {
		return "", fmt.Errorf("CHAT_GPT_API_URL environment variable is not set")
	}

	req, err := http.NewRequest("POST", chatGPTURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	log.Printf("GPT API Token: %s", gptApiToken)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+gptApiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to ChatGPT API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ChatGPT API returned non-OK status: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(responseBody), nil
}
