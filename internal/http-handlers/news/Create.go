package news

import (
	"elastic-search/internal/models/news"
	"fmt"
	"net/http"

	newsValidation "elastic-search/internal/validations/news"
	"github.com/gin-gonic/gin"
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
	news.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "news created.", "news": news})
}
