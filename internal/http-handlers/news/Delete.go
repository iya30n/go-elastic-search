package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/internal/services/elasticsearch"
	"elastic-search/pkg/database/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context) {
	newsItem := news.News{}
	if err := newsItem.Find(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	db := mysql.Connect()
	db.Transaction(func(tx *gorm.DB) error {
		if _, err := elasticsearch.Delete("go_news_index", newsItem); err != nil {
			return err
		}

		if err := newsItem.Delete(); err != nil {
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, gin.H{"message": "News Deleted!"})
}
