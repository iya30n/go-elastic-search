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
    log.Fatalln(router.Run(":2000"))
}
