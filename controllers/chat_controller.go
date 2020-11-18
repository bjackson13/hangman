package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/bjackson13/hangman/services/chat"
	"github.com/bjackson13/hangman/models/user"
	"net/http"
	"time"
)

/*RegisterChatRoutes register routes for chat on router*/
func RegisterChatRoutes(router *gin.Engine) {
	chatGroup := router.Group("/chat")
	chatGroup.Use(AuthMiddleware(), GetLobbyOrGame()) 
	{
		chatGroup.GET("/", getAllMessages)
		chatGroup.POST("/", addMessage)
		chatGroup.GET("/since/:time", getMessagesSince)
	} 
}

/*getAllMessages return all chat messages*/
func getAllMessages(c *gin.Context) {
	chatID := c.MustGet("chatID").(int)
	authedUser := c.MustGet("authorized-user").(*user.User)
	chatService := chat.NewService()

	userChat, err := chatService.GetAllMessages(chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not check messages, please try again",
		})
		return
	}

	userChat.SortMessages()
	c.HTML(http.StatusOK, "chat_messages", gin.H {
		"messages": userChat.Messages,
		"username": authedUser.Username,
	})	
}

func getMessagesSince(c *gin.Context) {
	chatID := c.MustGet("chatID").(int)
	authedUser := c.MustGet("authorized-user").(*user.User)
	chatService := chat.NewService()

	timestamp := c.Param("time")
	userChat, err := chatService.GetMessagesSince(timestamp, chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not check messages, please try again",
		})
		return
	}

	
	c.HTML(http.StatusOK, "chat_messages", gin.H {
		"messages": userChat.Messages,
		"username": authedUser.Username,
	})	
}

func addMessage(c *gin.Context) {
	chatID := c.MustGet("chatID").(int)
	authedUser := c.MustGet("authorized-user").(*user.User)
	chatService := chat.NewService()
	messageText := c.PostForm("message")
	err := chatService.AddMessage(chatID, authedUser.UserID, time.Now().Unix(), messageText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not add messages, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":	"Message added",
	})
}
