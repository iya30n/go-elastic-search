package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	newsId := c.Param("id")

	params := Params{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	newsItem := news.News{}
	db := mysql.Connect()
	db.Find(&newsItem, "id", newsId)

	db.Model(&newsItem).Updates(news.News{
		Title: params.Title,
		Content: params.Content,
	})

	c.JSON(http.StatusOK, gin.H{"message": "news updated", "news": newsItem})
}
