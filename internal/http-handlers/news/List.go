package news

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	news := &news.News{}
	db := mysql.Connect()
	db.Find(&news)

	c.JSON(http.StatusOK, news)
}

