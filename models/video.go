package models

import "time"

type Video struct {
    ID          string    `json:"id" bson:"_id,omitempty"`
    Title       string    `json:"title" bson:"title"`
    Description string    `json:"description" bson:"description"`
    PublishedAt time.Time `json:"publishedAt" bson:"publishedAt"`
    Thumbnails  []string  `json:"thumbnails" bson:"thumbnails"`
}