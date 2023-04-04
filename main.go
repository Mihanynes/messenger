package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"schedule/controllers"
	"schedule/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	r.Use(cors.AllowAll())
	r.Static("/static", "./static")

	public := r.Group("/")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	users := r.Group("/users")
	users.POST("/findUsers", controllers.FindUsers)

	chats := r.Group("/chats")
	chats.POST("/create", controllers.CreateChat)

	messages := r.Group("/messages")
	messages.POST("/create", controllers.CreateMessage)
	messages.PATCH("/update", controllers.UpdateMessage)

	r.Run(":8080")

}
