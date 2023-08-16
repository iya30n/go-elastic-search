package news

import (
	"elastic-search/internal/models/news"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	newsItem := news.News{}
	if err := newsItem.Find(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"news": newsItem})
}
