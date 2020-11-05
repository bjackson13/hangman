package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/lobby"
)

/*RegisterLobbyRoutes register lobby endpoints*/
func RegisterLobbyRoutes(router *gin.Engine) {
	lobbyGroup := router.Group("/lobby") 
	lobbyGroup.Use(AuthMiddleware()) 
	{
		lobbyGroup.GET("/", getLobby)
		lobbyGroup.GET("/lobbyUsers", getLobbyUsers)
		lobbyGroup.POST("/lobby/invite/:inviteeID", invitePlayer)
		lobbyGroup.GET("/lobby/invite/check", checkInvites)
		lobbyGroup.POST("/lobby/invite/accept", acceptInvite)
		lobbyGroup.POST("/lobby/invite/deny", denyInvite)
	}
}

func getLobby(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func getLobbyUsers(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()
	users, _ := lobbyService.GetLobbyUsers()
	c.HTML(http.StatusOK, "user_cards.html",gin.H{
		"user":	authedUser.Username,
		"LobbyUsers": users,
	})
}

func invitePlayer(c *gin.Context) {
	inviteeID := := c.Param("inviteeID")
}

func checkInvites(c *gin.Context) {
	
}

func acceptInvite(c *gin.Context) {
	
}

func denyInvite(c *gin.Context) {
	
})