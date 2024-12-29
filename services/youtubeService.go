package services

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"youtube-fam/config"
	"youtube-fam/database"
	"youtube-fam/models"

	"encoding/json"
	"fmt"
)

type YouTubeResponse struct {
    Items []struct {
        ID struct {
            VideoID string `json:"videoId"`
        } `json:"id"`
        Snippet struct {
            Title       string    `json:"title"`
            Description string    `json:"description"`
            PublishedAt time.Time `json:"publishedAt"`
            Thumbnails  struct {
                Default struct {
                    URL string `json:"url"`
                } `json:"default"`
            } `json:"thumbnails"`
        } `json:"snippet"`
    } `json:"items"`
}

func FetchLatestVideos() {
    for {
        fetchAndStoreVideos()
        time.Sleep(10 * time.Second)
    }
}

func fetchAndStoreVideos() {
    apiKey := config.AppConfig.YouTubeAPIKey
    searchQuery := config.AppConfig.SearchQuery
    url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&order=date&q=%s&type=video&key=%s", searchQuery, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error fetching videos: %v", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading response body: %v", err)
        return
    }

    var ytResponse YouTubeResponse
    if err := json.Unmarshal(body, &ytResponse); err != nil {
        log.Printf("Error unmarshalling response: %v", err)
        return
    }

    videoCollection := db.GetCollection("videos")
    for _, item := range ytResponse.Items {
        video := models.Video{
            ID:          item.ID.VideoID,
            Title:       item.Snippet.Title,
            Description: item.Snippet.Description,
            PublishedAt: item.Snippet.PublishedAt,
            Thumbnails:  []string{item.Snippet.Thumbnails.Default.URL},
        }
        _, err := videoCollection.InsertOne(context.TODO(), video)
        if err != nil {
            log.Printf("Error inserting video: %v", err)
        }
    }
}