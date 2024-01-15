package youtube_videos

type YouTubeResponse struct {
	Kind          string       `json:"kind"`
	Etag          string       `json:"etag"`
	NextPageToken string       `json:"nextPageToken"`
	RegionCode    string       `json:"regionCode"`
	PageInfo      PageInfo     `json:"pageInfo"`
	Items         []SearchItem `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type SearchItem struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type ID struct {
	Kind       string `json:"kind"`
	VideoID    string `json:"videoId,omitempty"`
	PlaylistID string `json:"playlistId,omitempty"`
}

type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	PublishTime          string     `json:"publishTime"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
