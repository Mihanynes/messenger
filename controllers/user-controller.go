package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "io/ioutil"
	"log"
	"net/http"
	"schedule/models"
	"schedule/utils/token"
	"strings"
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

func GetUser(c *gin.Context) {
	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SavePhoto(c *gin.Context) {
	formFile, err := c.FormFile("photo")

	if err != nil {
		log.Println("image upload error --> ", err)
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(formFile.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	c.SaveUploadedFile(formFile, "D:\\sites\\messanger\\messenger\\static\\avatars\\"+image)

	user, err := CurrentUser(c)

	if err != nil {
		return
	}

	user.UpdatePhoto("/avatars/" + image)

	c.JSON(http.StatusOK, gin.H{"image_src": "/avatars/" + image})

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

	ChatIcons, err := user.GetLastMessages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_last_messages": ChatIcons, "user_id": user.ID})
}

func CurrentUser(c *gin.Context) (models.User, error) {

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		return models.User{}, nil
	}

	u, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.User{}, nil
	}

	return u, err
}
