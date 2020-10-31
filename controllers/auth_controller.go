package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authService "github.com/bjackson13/hangman/services/authentication"
	"github.com/bjackson13/hangman/models/user"
)

/*RegisterAuthRoutes registering the paths for the authentication service*/
func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth") 
	{
		auth.POST("/login", login)
		auth.POST("/validateLogin", AuthMiddleware(), validate)
	}
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
		token := authService.GenerateSessionToken(*user)
	
		c.SetCookie("hjt", token, 86400, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "lobby.html",gin.H{
			"user":	user.Username,
		})
	}
}

/*validate small function to validate user logins. Primarily used for testing*/
func validate(c *gin.Context) {
	u := c.MustGet("authorized-user").(*user.User)
	message := "Hello " + u.Username
	c.JSON(http.StatusOK, gin.H{
		"message":  message,
	})
}