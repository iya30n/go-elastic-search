package news

import (
	"elastic-search/internal/models/news"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	newsValidation "elastic-search/internal/validations/news"
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

	err := newsItem.Update(news.News{
		Title:   params.Title,
		Content: params.Content,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "news updated", "news": newsItem})
}
