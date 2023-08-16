package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Create(c *gin.Context) {
	params := Params{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	db := mysql.Connect()

	news := &news.NewsModel{
		Title:   params.Title,
		Content: params.Content,
	}

	db.Create(news)

	c.JSON(http.StatusCreated, gin.H{"message": "news created.", "news": news})
}
