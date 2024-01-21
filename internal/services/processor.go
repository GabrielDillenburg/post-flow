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
	var audioFiles []string // Store paths of audio files

	for _, video := range videos {
		wg.Add(1)

		go func(video DomainVideo) {
			defer wg.Done()

			audioFile, err := audioDownloader(video.ID)
			if err != nil {
				log.Printf("Failed to download/extract audio for video %s: %v", video.ID, err)
				resultsChan <- TranscriptionResult{VideoID: video.ID, Error: err}
				return
			}

			// Add the path of the audio file to the slice
			audioFiles = append(audioFiles, audioFile)

			transcription, err := TranscriptionAPI(audioFile)
			if err != nil {
				log.Printf("Failed to transcribe audio for video %s: %v", video.ID, err)
				resultsChan <- TranscriptionResult{VideoID: video.ID, Error: err}
			} else {
				resultsChan <- TranscriptionResult{VideoID: video.ID, Transcription: transcription}
			}
		}(video)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []TranscriptionResult
	var hasErrors bool
	for result := range resultsChan {
		if result.Error != nil {
			hasErrors = true
			log.Printf("Error in processing video ID %s: %v", result.VideoID, result.Error)
		}
		results = append(results, result)
	}

	if hasErrors {
		return "", fmt.Errorf("errors occurred during video processing, check logs for details")
	}

	summary, err := SendToSummarizationAPI(results)
	if err != nil {
		log.Printf("Error summarizing transcriptions: %v", err)
		return "", fmt.Errorf("error summarizing transcriptions: %w", err)
	}

	// Delete all audio files after all processing is done
	for _, file := range audioFiles {
		if err := os.Remove(file); err != nil {
			log.Printf("Failed to delete audio file %s: %v", file, err)
		}
	}

	return summary, nil
}
