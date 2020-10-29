package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authService "github.com/bjackson13/hangman/services/authentication"
)

/*AuthMiddleWare verifies user is authticated before passing along to handler functions.*/
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("hjt"); err == nil {
			authorized := false
			user, err := authService.VerifyAndProcessToken(cookie.Value, c.ClientIP(), c.GetHeader("User-Agent"))
			if err == nil {
				authorized = true
			}

			if authorized {
				// If a valid user is returned
				c.Set("authorized-user", user)
				c.Next()
				return
			}
		}
		// if unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Unauthorized",
        })
        c.Abort()
        return
	}
}