package services

import (
	"fmt"
	"os"
	"os/exec"
)

func audioDownloader(videoID string) (string, error) {
	tempFile, err := os.CreateTemp("", "audio_*.mp3")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	tempFilePath := tempFile.Name()
	tempFile.Close()

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", tempFilePath, videoURL)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("yt-dlp command failed: %w", err)
	}

	return tempFilePath, nil
}
