package controllers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"time"
	authService "github.com/bjackson13/hangman/services/authentication"
	"github.com/bjackson13/hangman/models/user"
	"bytes"
	"net/http/httptest"
	"net/http"
)

/*this test is really gross looking but I needed to verify my middleware could authenticate
	 a user prior to calling a handler function and I could retrieve the user from the context*/
func TestAuthMiddleWare(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := authService.GenerateSessionToken(*testUser)
	
	/*Lots of setup for testing the middleware.
		There's probably a better way to do this but it works
	*/
	gin.SetMode(gin.TestMode)
	testFunc := AuthMiddleWare()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	cookie := &http.Cookie{}
	cookie.Name = "hjt"
	cookie.Value = token
	cookie.MaxAge = 86400
	cookie.Secure = false
	cookie.HttpOnly = true

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.AddCookie(cookie)

	//set client ip and user agent TODO
	c.Request.Header.Add("X-Forwarded-For", "192.168.1.1")
	c.Request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)")

	//Run middleware with context
	testFunc(c)

	//verify results
	authedUser := c.MustGet("authorized-user").(*user.User)

	if authedUser.Username != "auth" || authedUser.UserID != 101 {
		t.Errorf("Wrong user info retrieved")
	}
}