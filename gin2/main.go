package main

import (
	"fmt"
	"net/http"

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
	router.LoadHTMLGlob("index.html")

	//Initialize the routes
	initializeRoutes()

	//Start serving the application
	router.Run()
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present

func render(c *gin.Context, data gin.H, templateName string) {
	accept := c.Request.Header.Get("Accept")
	fmt.Printf("accept : %s", accept)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
