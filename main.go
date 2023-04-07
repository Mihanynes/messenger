package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"schedule/controllers"
	"schedule/middlewares"
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

	protected := r.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentU)

	users := r.Group("/users")
	users.POST("/findUsers", controllers.FindUsers)
	users.GET("/user_chats", controllers.GetAllUserChats)
	users.GET("/user_last_messages", controllers.GetUserLastMessages)

	chats := r.Group("/chats")
	chats.POST("/create", controllers.CreateChat)
	chats.POST("/all_messages", controllers.GetAllMessagesFromChat)

	messages := r.Group("/messages")
	messages.POST("/create", controllers.CreateMessage)
	messages.PATCH("/update", controllers.UpdateMessage)

	r.Run(":8080")

}
