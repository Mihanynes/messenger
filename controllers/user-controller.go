package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
	"schedule/utils/token"
)

type FindUsersInput struct {
	Username string `json:"username" binding:"required"`
}

func FindUsers(c *gin.Context) {
	var input FindUsersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := models.FindUser(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetAllUserChats(c *gin.Context) {

	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chats, err := user.GetAllUserChats()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_chats": chats})
}

func GetUserLastMessages(c *gin.Context) {

	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var chatIcons []models.ChatIcon
	ChatIcons, err := user.GetLastMessages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_last_messages": ChatIcons})
}

func CurrentUser(c *gin.Context) (models.User, error) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		return models.User{}, nil
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.User{}, nil
	}

	return u, err
}
