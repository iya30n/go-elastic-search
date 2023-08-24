package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/internal/services/elasticsearch"
	"elastic-search/pkg/database/mysql"
	"fmt"
	"net/http"

	newsValidation "elastic-search/internal/validations/news"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(c *gin.Context) {
	params := newsValidation.UpdateValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	newsItem := news.News{}
	if err := newsItem.Find(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	db := mysql.Connect()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := newsItem.Update(news.News{
			Title:   params.Title,
			Content: params.Content,
		})

		if err != nil {
			return err
		}

		if _, err := elasticsearch.Index("go_news_index", newsItem); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "news updated", "news": newsItem})
}
