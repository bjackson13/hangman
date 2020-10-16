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
	controllers.RegisterAuthRoutes(router)
	router.Run()
}