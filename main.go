// main.go

package main

import (
	"octopus.com/gin-test/configurations"
	"octopus.com/gin-test/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	configurations.Load()
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("view/templates/*")

	// Initialize the routes
	routes.Register(router)

	// Start serving the application
	router.Run(configurations.Config.Port)
}
