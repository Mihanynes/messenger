package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
	"schedule/utils/token"
)

type CreateChatInput struct {
	UserId uint `json:"user_id"`
}

func CreateChat(c *gin.Context) {
	var input CreateChatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := CurrentUser(c)

	chat, err := (&models.Chat{FirstUserID: user.ID, SecondUserID: input.UserId}).FindChat()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}

func GetAllMessagesFromChat(c *gin.Context) {
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allMessages, err := chat.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)

	c.JSON(http.StatusOK, gin.H{"all_messages": allMessages, "user_id": user_id})
}
