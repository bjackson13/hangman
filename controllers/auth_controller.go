package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authService "github.com/bjackson13/hangman/services/authentication"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/game"
	"github.com/bjackson13/hangman/services/lobby"
)

/*RegisterAuthRoutes registering the paths for the authentication service*/
func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth") 
	{
		auth.POST("/login", login)
		auth.GET("/login", getLoginPage)
		auth.POST("/validateLogin", AuthMiddleware(), validate)
		auth.GET("/validateLogin", AuthMiddleware(), validate)
		auth.GET("/logout", AuthMiddleware(), logout)
		auth.POST("/logout", AuthMiddleware(), logout)
	}
}

func getLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

/*login controller handles our user login. will generate and return session cookie if valid, otherwise return 400*/
func login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("pass")
	user, err := authService.AuthenticateUserLogin(username, password, c.ClientIP(), c.GetHeader("User-Agent"))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error":  "Unauthorized, please log in to continue",
		})
	} else {

		go func() {
			lobbyService := lobby.NewService()
			lobbyService.AddUser(user.UserID)
		}()

		token := authService.GenerateSessionToken(*user)
	
		c.SetCookie("hjt", token, 86400, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "lobby.html",gin.H{
			"user":	user.Username,
		})
	}
}

/*logout clear out hjt cookie effectivley invalidating users session*/
func logout(c *gin.Context) {
	u := c.MustGet("authorized-user").(*user.User)
	c.SetCookie("hjt", "", 0, "/", "localhost", false, true)
	lobbyService := lobby.NewService()
	gameService := games.NewService()

	go func() {
		lobbyService.RemoveUser(u.UserID)
		game := gameService.GetUserGame(u.UserID)
		if game != nil {
			gameService.EndGame(game.GameID)
		}
	}()
}

/*validate small function to validate user logins. Primarily used for testing*/
func validate(c *gin.Context) {
	u := c.MustGet("authorized-user").(*user.User)
	message := "Hello " + u.Username
	c.JSON(http.StatusOK, gin.H{
		"message":  message,
	})
}