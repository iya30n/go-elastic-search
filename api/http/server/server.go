package server

import (
	"elastic-search/internal/http-handlers/news"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
    router := gin.Default()
    router.GET("/news", news.List)
    router.POST("/news", news.Create)
    log.Fatalln(router.Run(":2000"))
}
