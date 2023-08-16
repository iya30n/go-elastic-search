package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	newsValidation "elastic-search/internal/validations/news"
)

func Create(c *gin.Context) {
	params := newsValidation.CreateValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	db := mysql.Connect()

	news := news.News{
		Title:   params.Title,
		Content: params.Content,
	}

	db.Create(&news)

	c.JSON(http.StatusCreated, gin.H{"message": "news created.", "news": news})
}
