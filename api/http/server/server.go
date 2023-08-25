package server

import (
	"elastic-search/internal/http-handlers/news"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
    router := gin.Default()
    router.GET("/news", news.List)
    router.POST("/news", news.Create)
    router.PATCH("/news/:id", news.Update)
    router.GET("/news/:id", news.Detail)
    router.GET("/news/search", news.Search)
    log.Fatalln(router.Run(":2000"))
}
