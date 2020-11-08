package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/services/lobby"
	"strconv"
)

/*RegisterLobbyRoutes register lobby endpoints*/
func RegisterLobbyRoutes(router *gin.Engine) {
	lobbyGroup := router.Group("/lobby") 
	lobbyGroup.Use(AuthMiddleware()) 
	{
		lobbyGroup.GET("/", getLobby)
		lobbyGroup.GET("/lobbyUsers", getLobbyUsers)
		lobbyGroup.POST("/invite/user/:inviteeID", invitePlayer)
		lobbyGroup.GET("/invite/check", checkInvites)
		lobbyGroup.POST("/invite/accept", acceptInvite)
		lobbyGroup.POST("/invite/deny", denyInvite)
	}
}

func getLobby(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func getLobbyUsers(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()
	statusChan := make(chan bool)
	go func() {
		statusChan <- lobbyService.UserIsInLobby(authedUser.UserID)
	}()

	users, _ := lobbyService.GetLobbyUsers()

	inLobby := <- statusChan
	
	if !inLobby {
		c.JSON(http.StatusFound, gin.H{
			"url":	"/game",
		})
	} else {
			c.HTML(http.StatusOK, "user_cards.html",gin.H{
			"user":	authedUser.Username,
			"LobbyUsers": users,
		})
	}
}

func invitePlayer(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	inviteeID, _ := strconv.Atoi(c.Param("inviteeID"))

	lobbyService := lobby.NewService()
	err := lobbyService.InviteUserToPlay(inviteeID, authedUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not invite user to game, please try again",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success":	"Great success, very nice invite!",
		})
	} 
}

func checkInvites(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()

	inviterName, inviterID, err := lobbyService.CheckInvites(authedUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not check invites, please try again",
		})
		return
	}

	if inviterName != nil {
		c.HTML(http.StatusOK, "invite.html", gin.H{
			"success":	"User has invite",
			"username": inviterName,
			"inviterID": inviterID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"nil":	"No invites found",
		})
	}
}

func acceptInvite(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()

	/*it should be safe to skip the error check here, if one occurs the next function should blow up but be handled*/
	inviterID, _ := strconv.Atoi(c.PostForm("inviterID")) 
	_, err := lobbyService.AcceptInvite(authedUser.UserID, inviterID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not accept invite, please try again",
		})
		return
	}

	c.Redirect(http.StatusFound, "/game")
}

func denyInvite(c *gin.Context) {
	authedUser := c.MustGet("authorized-user").(*user.User)
	lobbyService := lobby.NewService()

	err := lobbyService.DenyInvite(authedUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":	"Could not deny invite, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":	"Invite Denied",
	})
}