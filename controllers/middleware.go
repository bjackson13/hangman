package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authService "github.com/bjackson13/hangman/services/authentication"
	"github.com/bjackson13/hangman/services/game"
	"github.com/bjackson13/hangman/services/lobby"
	"github.com/bjackson13/hangman/models/user"
	"time"
	"log"
)

var lobbyChatID int = -404

/*AuthMiddleware verifies user is authticated before passing along to handler functions.*/
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("hjt"); err == nil {
			authorized := false
			authedUser, err := authService.VerifyAndProcessToken(cookie.Value, c.ClientIP(), c.GetHeader("User-Agent"))
			if err == nil {
				authorized = true
			} else {
				log.Println(err)
			}
			
			if authorized {
				// If a valid user is returned
				c.Set("authorized-user", authedUser)
				c.Next()
				return
			}
		} else {
			log.Println(err)
		} 
		
		// if unauthorized
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error":  "Unauthorized, please log in to continue",
		})
        c.Abort()
        return
	}
}

/*GetLobbyOrGame determine users chatID by finding if they are in a game or in the lobby*/
func GetLobbyOrGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		authedUser := c.MustGet("authorized-user").(*user.User)
		lobbyChan := make(chan int)
		gameChan := make(chan int)

		go func() {
			lobbyService := lobby.NewService()
			if lobbyService.UserIsInLobby(authedUser.UserID) {
				lobbyChan <- lobbyChatID
			}
		}()

		go func() {
			gameService := games.NewService()
			if game := gameService.GetUserGame(authedUser.UserID); game != nil {
				gameChan <- game.GameID
			}
		}()

		select {
			case chatID := <- lobbyChan:
				c.Set("chatID", chatID)
				c.Next()
				return
			case chatID := <- gameChan:
				c.Set("chatID", chatID)
				c.Next()
				return
			case <-time.After(time.Second * 2): //timeout after 2 seconds of waiting
				c.Redirect(http.StatusFound, "/")
		}
	}
}