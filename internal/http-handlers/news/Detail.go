package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	newsId := c.Param("id")
	newsItem := news.News{}

	db := mysql.Connect()
	db.Find(&newsItem, "id", newsId)

	c.JSON(http.StatusOK, gin.H{"news": newsItem})
}
