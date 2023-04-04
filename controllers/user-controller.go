package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedule/models"
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
