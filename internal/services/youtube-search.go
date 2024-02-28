package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	youtube_videos "github.com/gabrieldillenburg/post-flow/internal/models"
)

type Query string
type DomainVideo struct {
	ID          string
	Description string
	Title       string
}

func GetYouTubeVideos(q Query) ([]DomainVideo, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("YOUTUBE_API_KEY environment variable is not set")
	}

	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?part=snippet&type=video&videoDuration=short&maxResults=1&q=%s&key=%s", q, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from YouTube API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("YouTube API returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var ytResponse youtube_videos.Video
	err = json.Unmarshal(body, &ytResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	var videos []DomainVideo
	for _, item := range ytResponse.Items {
		if item.ID.Kind == "youtube#video" {
			videos = append(videos, DomainVideo{
				ID:          item.ID.VideoID,
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
			})
		}
	}

	return videos, nil
}
