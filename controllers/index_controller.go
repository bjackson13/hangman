package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/game"
	"github.com/bjackson13/hangman/services/lobby"
)

/*RegisterIndexRoutes register base pages*/
func RegisterIndexRoutes(router *gin.Engine) {
	index := router.Group("/") 
	index.Use(AuthMiddleware()) 
	{
		index.GET("/", directToGameOrLobby)
	}
}

func directToGameOrLobby(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()
	lobbyService := lobby.NewService()

	userGame := gameService.GetUserGame(authedUser.UserID)
	if userGame == nil {
		lobbyChan := make(chan []user.User)
		/*Go add the user to the lobby and get all uses from the lobby*/
		go func() {
			lobbyService.AddUser(authedUser.UserID)
			uArr, _ := lobbyService.GetLobbyUsers()
			lobbyChan <- uArr
		}()

		users := <- lobbyChan
		c.HTML(http.StatusOK, "lobby.html",gin.H{
			"title": "Lobby",
			"user":	authedUser.Username,
			"LobbyUsers": users,
		})
	}
	c.Redirect(http.StatusFound, "/game/")
}
