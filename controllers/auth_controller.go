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
	user, err := authService.AuthenticateUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "unauthorized",
		})
	} else {
		token := authService.GenerateSessionToken(*user)
	
		c.JSON(http.StatusOK, gin.H{
			"status":  "authorized",
			"hjt": token,
		})
	}
}