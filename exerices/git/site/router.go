package site

import "github.com/gin-gonic/gin"

var router *gin.Engine

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
}

// MainRouter
func MainRoute() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}
