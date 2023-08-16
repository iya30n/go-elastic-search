package news

import (
	"github.com/gin-gonic/gin"
    "net/http"
)

func List(ctx *gin.Context) {
	data := map[string]interface{}{
		"result": "the list of the news",
	}

	ctx.JSONP(http.StatusOK, data)
}

