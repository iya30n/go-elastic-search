package news

import (
	"net/http"

	"elastic-search/internal/services/elasticsearch"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	title, content := c.Query("title"), c.Query("content")

	searchQuery := make(map[string]string)

	if len(title) > 0 {
		searchQuery["title"] = title
	}

	if len(content) > 0 {
		searchQuery["content"] = content
	}

	res, err := elasticsearch.Search("go_news_index", searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}
