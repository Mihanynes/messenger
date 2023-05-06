package main

import (
	"encoding/json"
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

type UserConn struct {
	UserId uint `json:"user_id"`
}

type ClientsChats struct {
	CompanionID uint            `json:"companion_id"`
	Connection  *websocket.Conn `json:"connection"`
}

var clients = make(map[uint]*websocket.Conn)
var clientsChats = make(map[uint]ClientsChats)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	_, message, _ := connection.ReadMessage()
	var input UserConn
	json.Unmarshal(message, &input)
	clients[input.UserId] = connection
}

type MessageSend struct {
	UserId  uint           `json:"user_id"`
	Message models.Message `json:"message"`
}

func wsSendMessage(w http.ResponseWriter, r *http.Request) {
	connection, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	_, message, _ := connection.ReadMessage()
	var input MessageSend
	json.Unmarshal(message, &input)
	if val, ok := clients[input.UserId]; ok {
		obj, _ := json.Marshal(input.Message)
		val.WriteMessage(websocket.TextMessage, obj)
	}
	val, ok := clientsChats[input.UserId]
	fmt.Println(val)
	if ok && val.CompanionID == input.Message.SenderID {
		obj, _ := json.Marshal(input.Message)
		fmt.Println(obj)
		val.Connection.WriteMessage(websocket.TextMessage, obj)
	}
}

type GetMessage struct {
	UserID      uint `json:"user_id"`
	CompanionId uint `json:"companion_id"`
}

func wsGetMessage(w http.ResponseWriter, r *http.Request) {
	connection, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	_, message, _ := connection.ReadMessage()
	var input GetMessage
	json.Unmarshal(message, &input)
	clientsChats[input.UserID] = ClientsChats{CompanionID: input.CompanionId, Connection: connection}
}

func main() {

	models.ConnectDataBase()

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.Static("/static", "./static")

	router.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	router.GET("/ws/sendMessage", func(c *gin.Context) {
		wsSendMessage(c.Writer, c.Request)
	})

	router.GET("/ws/getMessage", func(c *gin.Context) {
		wsGetMessage(c.Writer, c.Request)
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
	users.GET("/user_last_messages", controllers.GetUserLastMessages)
	users.GET("/getUser", controllers.GetUser)
	users.POST("/savePhoto", controllers.SavePhoto)

	chats := router.Group("/chats")
	chats.POST("/create", controllers.CreateChat)
	chats.POST("/all_messages", controllers.GetAllMessagesFromChat)

	messages := router.Group("/messages")
	messages.POST("/create", controllers.CreateMessage)
	messages.PATCH("/update", controllers.UpdateMessage)

	router.Run(":8080")

}
