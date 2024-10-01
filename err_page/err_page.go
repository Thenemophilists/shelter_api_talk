package errPage

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Return404Page(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"error": "Page not found",
	})
}