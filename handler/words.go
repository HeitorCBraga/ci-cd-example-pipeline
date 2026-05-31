package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type wordsRequest struct {
	Text string `json:"text" binding:"required"`
}

type wordsResponse struct {
	Words []string `json:"words"`
}

func Words(c *gin.Context) {
	var req wordsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wordsResponse{Words: strings.Fields(req.Text)})
}
