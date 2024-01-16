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
}

func ProcessVideosWorkflow(videos []DomainVideo) (string, error) {
	var wg sync.WaitGroup
	results := make([]TranscriptionResult, 0, len(videos))
	resultsChan := make(chan TranscriptionResult, len(videos))

	for _, video := range videos {
		wg.Add(1)

		go func(video DomainVideo) {
			defer wg.Done()

			audioFile, err := audioDownloader(video.ID)
			if err != nil {
				log.Printf("Failed to download/extract audio for video %s: %v", video.ID, err)
				return
			}

			transcription, err := TranscriptionAPI(audioFile)
			if err != nil {
				log.Printf("Failed to download/extract audio for video %s: %v", video.ID, err)
				return
			}

			resultsChan <- TranscriptionResult{VideoID: video.ID, Transcription: transcription}

			//delete the audio file. To do: save this on a database.
			if err := os.Remove(audioFile); err != nil {
				log.Printf("Failed to delete audio file %s: %v", audioFile, err)
			}

		}(video)
	}

	// routine to collect all results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		results = append(results, result)
	}

	summary, err := SendToSummarizationAPI(results)
	if err != nil {
		log.Printf("error to summarize the videos")
		return "", fmt.Errorf("error summarizing trascriptions: %w", err)
	}

	return summary, nil
}

func SendToSummarizationAPI(results []TranscriptionResult) (string, error) {
	panic("unimplemented")
}

func TranscriptionAPI(audioFile string) (string, error) {
	panic("unimplemented")
}
