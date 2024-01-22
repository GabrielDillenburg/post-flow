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

	tempFilePath := fmt.Sprintf("/tmp/%s_audio.mp3", videoID) // Unique file name for each download

	log.Printf("Downloading audio to temporary file: %s", tempFilePath)

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	commandString := fmt.Sprintf("yt-dlp -k -x --audio-format mp3 -o '%s' '%s'", tempFilePath, videoURL)
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

	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		log.Printf("Downloaded audio file does not exist: %s", tempFilePath)
		return "", fmt.Errorf("downloaded audio file does not exist: %s", tempFilePath)
	}

	log.Printf("Audio downloaded successfully to: %s", tempFilePath)

	return tempFilePath, nil
}
