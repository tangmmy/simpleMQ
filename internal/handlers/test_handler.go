package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
