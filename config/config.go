package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)


// getEnv retrieves the value of the environment variable named by the key.

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
type Config struct {
    Port         string
    MongoDBURI   string
    DatabaseName string
    YouTubeAPIKey string
    SearchQuery  string
}

var AppConfig Config

func LoadConfig() {

    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found. Using system environment variables.")
    }

    AppConfig = Config{
        Port:         getEnv("PORT", "8080"),
        MongoDBURI:   getEnv("MONGODB_URI", ""),
        DatabaseName: getEnv("DATABASE_NAME", "campaignDB"),
        YouTubeAPIKey: getEnv("YOUTUBE_API_KEY", ""),
        SearchQuery:  getEnv("SEARCH_QUERY", "football"),
    }

    if AppConfig.MongoDBURI == "" {
        log.Fatal("MONGODB_URI is not set in the environment")
    }
    if AppConfig.YouTubeAPIKey == "" {
        log.Fatal("YOUTUBE_API_KEY is not set in the environment")
    }
}