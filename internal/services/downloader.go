package services

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func audioDownloader(videoID string) (string, error) {
	if videoID == "" {
		return "", errors.New("videoID cannot be empty")
	}

	tempFile, err := os.CreateTemp("", "audio_*.mp3") // Change extension to .webm or a generic one
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			log.Printf("Failed to remove temp file: %s, error: %v", tempFile.Name(), err)
		}
	}()

	tempFilePath := tempFile.Name()
	tempFile.Close()

	log.Printf("Downloading audio to temporary file: %s", tempFilePath)

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	commandString := fmt.Sprintf("yt-dlp -x --audio-format best -o '%s' '%s'", tempFilePath, videoURL)
	cmd := exec.Command("bash", "-c", commandString)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()
	cmd.Dir = "/app"

	if err := cmd.Run(); err != nil {
		log.Printf("yt-dlp stdout: %s", stdout.String())
		log.Printf("yt-dlp stderr: %s", stderr.String())
		return "", fmt.Errorf("yt-dlp command failed: %w. Stderr: %s", err, stderr.String())
	}

	return tempFilePath, nil
}
