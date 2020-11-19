package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bjackson13/hangman/controllers"
	"github.com/bjackson13/hangman/services/config"
	"net/http/cgi" //for when I try to deploy on solcace
	"log"
	"os"
)

func main() {
	err := config.LoadEnvVariables()
	if err != nil{
		log.Fatal("Failed to load env variables")
		return
	}
	
	isCGI := os.Getenv("CGI_USE")

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

	if isCGI == "true" {
		//For use with Common gateway interface (I am going to attempt to deploy on solace)
		log.Fatal(cgi.Serve(router))
	} else {
		router.Run()
	}
}