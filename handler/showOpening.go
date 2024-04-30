package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// letra maiuscula porque queremos exporta as funcoes do package handler
func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}
