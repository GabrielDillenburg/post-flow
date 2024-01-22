package services

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type TranscriptionResult struct {
	VideoID       string
	Transcription string
	Error         error
}

func ProcessVideosWorkflow(videos []DomainVideo) (string, error) {
	var wg sync.WaitGroup
	resultsChan := make(chan TranscriptionResult, len(videos))
	var audioFiles []string // Slice to store audio file paths
	var mutex sync.Mutex    // Mutex to protect access to the audioFiles slice

	// First, download all audios
	for _, video := range videos {
		wg.Add(1)

		go func(video DomainVideo) {
			defer wg.Done()

			audioFilePath, err := audioDownloader(video.ID)
			if err != nil {
				log.Printf("Failed to download audio for video %s: %v", video.ID, err)
				resultsChan <- TranscriptionResult{VideoID: video.ID, Error: err}
				return
			}

			mutex.Lock()
			audioFiles = append(audioFiles, audioFilePath) // Safely store the audio file path
			mutex.Unlock()
		}(video)
	}

	// Wait for all downloads to complete
	wg.Wait()

	// process each audio file
	for _, audioFilePath := range audioFiles {
		go func(audioFilePath string) {
			transcription, err := TranscriptionAPI(audioFilePath)
			if err != nil {
				log.Printf("Failed to transcribe audio: %v", err)
				resultsChan <- TranscriptionResult{Error: err}
				return
			}

			resultsChan <- TranscriptionResult{Transcription: transcription, Error: nil}
		}(audioFilePath)
	}

	var results []TranscriptionResult

	for i := 0; i < len(videos); i++ {
		result := <-resultsChan
		if result.Error != nil {
			return "", fmt.Errorf("errors occurred during video processing, check logs for details")
		}
		results = append(results, result)
	}

	summary, err := SendToSummarizationAPI(results)
	if err != nil {
		log.Printf("Error summarizing transcriptions: %v", err)
		return "", fmt.Errorf("error summarizing transcriptions: %w", err)
	}

	// Delete all audio files after all processing is done
	mutex.Lock()
	for _, file := range audioFiles {
		if err := os.Remove(file); err != nil {
			log.Printf("Failed to delete audio file %s: %v", file, err)
		}
	}
	mutex.Unlock()

	return summary, nil
}
