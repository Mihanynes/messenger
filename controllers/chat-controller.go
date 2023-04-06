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
	chat, err := (&input).FindChat()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}

func GetAllMessagesFromChat(c *gin.Context) {
	var input models.Chat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chat, err0 := (&input).FindChat()
	if err0 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err0.Error()})
		return
	}
	allMessages, err := chat.GetAllMessages()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"all_messages": allMessages})
}
