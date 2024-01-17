
# YouTube Transcription and Summarization Service

This project provides a service that searches for YouTube videos based on a query, downloads the audio, transcribes the audio using a third-party API, and then summarizes the transcriptions using ChatGPT API. It's built using Go and integrates various external services for transcription and summarization.

## Features

- **YouTube Video Search**: Search for videos on YouTube using a specific query.
- **Audio Extraction**: Download and extract audio from the YouTube videos.
- **Transcription**: Send the audio to a transcription service and receive text transcriptions.
- **Summarization**: Summarize the transcriptions using ChatGPT API.

## Setup and Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/GabrielDillenburg/post-flow/tree/main
   ```

2. **Set Environment Variables**:
   - `YOUTUBE_API_KEY`: Your API key for the YouTube Data API.
   - `TRANSCRIPTION_API_URL`: The URL for the transcription service.
   - `TRANSCRIPTION_API_TOKEN`: The Bearer token for the transcription service.
   - `CHAT_GPT_API_URL`: The URL for the ChatGPT API.
3. **Install Dependencies** (if any):

   ```bash
   go get ./...
   ```

4. **Run the Service**:

   ```bash
   go run main.go 
   ```

   or

   ```bash
    podman-compose up
   ```

## Usage

Send a POST request to `/transcribe` endpoint with a JSON body containing the search query. For example:

```json
{
  "query": "your search term"
}
```

The service will search YouTube videos, download and transcribe the audio, and return a summarized response.

## Project Structure

- `server.go`: Initializes the HTTP server and defines endpoints.
- `youtube-transcriptor.go`: Handles the logic for fetching YouTube videos.
- `downloader.go`: Manages the downloading and extraction of audio from YouTube videos.
- `processor.go`: Orchestrates the processing workflow, including transcription and summarization.
- `summarizator.go`: Handles the communication with the ChatGPT API for summarization.
- `transcriptor.go`: Manages the transcription requests to the third-party API.

## Contributing

Contributions to the project are welcome! Please create a pull request with your proposed changes.

# License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
