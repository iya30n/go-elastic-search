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

func Create(c *gin.Context) {
	params := newsValidation.CreateValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	news := news.News{
		Title:   params.Title,
		Content: params.Content,
	}

	db := mysql.Connect()
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := news.Save(); err != nil {
			return err
		}

		if _, err := elasticsearch.Index("go_news_index", news); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "news created.", "news": news})
}
