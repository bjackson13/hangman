package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bjackson13/hangman/controllers"
)

func main()  {
	/*
		Create a Gin router and attach routes to it through controller methods.
	*/
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	controllers.RegisterIndexRoutes(router)
	controllers.RegisterAuthRoutes(router)
	controllers.RegisterLobbyRoutes(router)
	router.Run()
}