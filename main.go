package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bjackson13/hangman/controllers"
	"github.com/bjackson13/hangman/services/config"
	"log"
)

func main() {
	err := config.LoadEnvVariables()
	if err != nil{
		log.Fatal("Failed to load env variables")
		return
	}

	/*
		Create a Gin router and attach routes to it through controller methods.
	*/
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	controllers.RegisterIndexRoutes(router)
	controllers.RegisterAuthRoutes(router)
	controllers.RegisterLobbyRoutes(router)
	controllers.RegisterGameRoutes(router)
	controllers.RegisterChatRoutes(router)

	router.Run()
}