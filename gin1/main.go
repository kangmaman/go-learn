package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	//set the router as the default one provided by Gin
	router = gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	//process the templates at the start so that they don't have to loaded from the disk again.
	//This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	//Initialize the routes
	initializeRoutes()

	//Start serving the application
	router.Run()
}
