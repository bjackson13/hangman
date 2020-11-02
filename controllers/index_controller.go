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
		index.GET("/lobby", getLobby)
	}
}

func directToGameOrLobby(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()
	lobbyService := lobby.NewService()

	userGame := gameService.GetUserGame(authedUser.UserID)
	if userGame == nil {
		go func() {
			lobbyService.AddUser(authedUser.UserID)
		}()
		users, _ := lobbyService.GetLobbyUsers()
		c.HTML(http.StatusOK, "lobby.html",gin.H{
			"user":	authedUser.Username,
			"LobbyUsers": users,
		})
	}

	c.HTML(http.StatusOK, "game.html",gin.H{
		"user":	authedUser.Username,
		//need to add more down the line
	})

}

func getLobby(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()
	go func() {
		if !lobbyService.UserIsInLobby(authedUser.UserID){
			lobbyService.AddUser(authedUser.UserID)
		}
	}()
	users, _ := lobbyService.GetLobbyUsers()
	c.HTML(http.StatusOK, "lobby.html",gin.H{
		"user":	authedUser.Username,
		"LobbyUsers": users,
	})
}

