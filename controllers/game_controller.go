package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/game"
	"strconv"
)

/*RegisterGameRoutes register game endpoints*/
func RegisterGameRoutes(router *gin.Engine) {
	gameGroup := router.Group("/game") 
	gameGroup.Use(AuthMiddleware()) 
	{
		gameGroup.GET("/", getGame)
		gameGroup.POST("/makeGuess", makeGuess)
		gameGroup.GET("/guess", checkPendingGuess)
		gameGroup.GET("/guess/deny", denyGuess)
		gameGroup.GET("/guess/incorrect", getIncorrectGuesses)
		gameGroup.POST("/word/create", createWord)
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
		if game.PendingGuess != "" {
			c.JSON(http.StatusOK, gin.H{
				"error":	"Guess Already Pending",
			})
			return
		} else if game.WordID == -1 {
			c.JSON(http.StatusOK, gin.H{
				"error":	"Please wait for word creator to pick word",
			})
			return
		}
		alreadyMadeErr := "Guess already made"
		result := gameService.MakeGuess(game.GameID, authedUser.UserID, game.WordID, guess)
		if result.Error == nil {
			c.JSON(http.StatusOK, gin.H{
				"success":	"Guess submitted!",
			})
			return
		} else if result.Error.Error() == alreadyMadeErr {
			c.JSON(http.StatusOK, gin.H{
				"error":	"Guess Already Made",
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not make a guess, please try again",
	})
}

func checkPendingGuess(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		guess, err := gameService.CheckPendingGuesses(game.GameID, authedUser.UserID)
		if err == nil {
			c.HTML(http.StatusOK, "pending_guess", gin.H{
				"guess":	guess,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"empty":	"no guesses",
		})
		return
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
		result := gameService.DenyGuess(*game)
		if result.Error == nil {
			c.JSON(http.StatusOK, gin.H{
				"success":	"Guess denied/removed",
				"guessesExceeded": result.LimitExceeded,
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not deny guess, please try again",
	})
}

func createWord(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		length,_ := strconv.Atoi(c.PostForm("length"))
		err := gameService.AddWord(game.GameID, length)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success":	"Word created",
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not create word, please try again",
	})
}

func getIncorrectGuesses(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	gameService := games.NewService()

	/*If we have a valid game*/
	if game := gameService.GetUserGame(authedUser.UserID); game != nil {
		if game.WordID == -1 {
			c.JSON(http.StatusOK, gin.H{
				"error":	"Please wait for word creator to pick word",
			})
			return
		}
		guesses, err := gameService.GetIncorrectGuesses(game.WordID)
		if err == nil {
			c.HTML(http.StatusOK, "incorrect_guesses", gin.H{
				"incorrect":	guesses,
			})
			return
		}
	}
	/*If user is not in game or is not the guessing user*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":	"Could not get guesses, please try again",
	})
}