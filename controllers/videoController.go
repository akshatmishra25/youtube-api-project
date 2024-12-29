package controllers

import (
    "context"
    "net/http"
    "strconv"

    "youtube-fam/database"
    "youtube-fam/models"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetVideos(c *gin.Context) {
    videoCollection := db.GetCollection("videos")

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    skip := (page - 1) * limit

    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"publishedAt", -1}})
    findOptions.SetSkip(int64(skip))
    findOptions.SetLimit(int64(limit))

    cursor, err := videoCollection.Find(context.TODO(), bson.M{}, findOptions)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(context.TODO())	

    var videos []models.Video
    if err := cursor.All(context.TODO(), &videos); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, videos)
}