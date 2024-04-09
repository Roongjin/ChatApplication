package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Raise500Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   err,
	})
	c.Abort()
}
