package main

import (
	"log"

	"youtube-fam/config"
	"youtube-fam/database"
	"youtube-fam/routes"
	"youtube-fam/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())

	routes.AddRoutes(router)
	go services.FetchLatestVideos() // goroutine to fetch videos every 10 second
	log.Printf("Starting server on port %s", config.AppConfig.Port)
	router.Run(":" + config.AppConfig.Port)
}
