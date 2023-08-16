package news

import (
	"elastic-search/internal/models/news"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	news := &news.News{}
	newsList, err := news.All()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"news": newsList})
}

