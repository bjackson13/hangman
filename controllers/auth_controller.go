package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authService "github.com/bjackson13/hangman/services/authentication"
)

/*RegisterAuthRoutes registering the paths for the authentication service*/
func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth") 
	{
		auth.POST("/login", login)
	}
}

func login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("pass")
	user, err := authService.AuthenticateUserLogin(username, password, c.ClientIP(), c.GetHeader("User-Agent"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "unauthorized",
		})
	} else {
		token := authService.GenerateSessionToken(*user)
	
		c.SetCookie("hjt", token, 86400, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status":  "authorized",
			"user":	user.Username,
		})
	}
}