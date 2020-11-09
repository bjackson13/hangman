package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/game"
)

/*RegisterGameRoutes register game endpoints*/
func RegisterGameRoutes(router *gin.Engine) {
	gameGroup := router.Group("/game") 
	gameGroup.Use(AuthMiddleware()) 
	{
		gameGroup.GET("/", getGame)
	}
}

func getGame(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		c.HTML(http.StatusOK, "game.html",gin.H{
			"title": "Game",
			"userID": authedUser.UserID,
			"user":	authedUser.Username,
			"isWordCreator": game.WordCreatorID == authedUser.UserID,
			"gameID": game.GameID,
			"pendingGuess": game.PendingGuess,
			"wordCreatorID": game.WordCreatorID,
			"guessingUserID": game.GuessingUserID,
		})
		return
	}
	
	c.Redirect(http.StatusFound, "/")
}