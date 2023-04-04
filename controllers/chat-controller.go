package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
)

func CreateChat(c *gin.Context) {
	var input models.Chat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chat, err := (&input).Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}
