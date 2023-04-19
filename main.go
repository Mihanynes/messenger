package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
	"schedule/controllers"
	"schedule/middlewares"
	"schedule/models"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func main() {

	models.ConnectDataBase()

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.Static("/static", "./static")

	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	public := router.Group("/")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentU)

	users := router.Group("/users")
	users.POST("/findUsers", controllers.FindUsers)
	users.GET("/user_chats", controllers.GetAllUserChats)
	users.GET("/{user_last_messages}", controllers.GetUserLastMessages)

	chats := router.Group("/chats")
	chats.POST("/create", controllers.CreateChat)
	chats.POST("/all_messages", controllers.GetAllMessagesFromChat)

	messages := router.Group("/messages")
	messages.POST("/create", controllers.CreateMessage)
	messages.PATCH("/update", controllers.UpdateMessage)

	router.Run(":8080")

}
