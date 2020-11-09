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
		gameGroup.POST("/makeGuess", makeGuess)
		gameGroup.GET("/guess", checkGuess)
		gameGroup.GET("/guess/deny", denyGuess)
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

func makeGuess(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		guess := c.PostForm("guess")
		err := gameService.MakeGuess(game.GameID, authedUser.UserID, guess)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success":	"Guess submitted!",
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not make a guess, please try again",
	})
}

func checkGuess(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		guess, err := gameService.CheckGuesses(game.GameID, authedUser.UserID)
		if err == nil {
			c.HTML(http.StatusOK, "pending_guess", gin.H{
				"guess":	guess,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"empty":	"no guesses",
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not check guesses, please try again",
	})
}

func denyGuess(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		err := gameService.DenyGuess(game.GameID)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success":	"Guess denied/removed",
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not deny guess, please try again",
	})
}