package routes

import (
    "youtube-fam/controllers"

    "github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
    router.GET("/videos", controllers.GetVideos) 
}