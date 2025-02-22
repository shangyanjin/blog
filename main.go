package main

import (
	"blog/config"
	"blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	initConfig()

	// Create gin engine
	app := gin.Default()

	// Initialize templates
	initTemplate(app)

	// Initialize static files
	initStatic(app)

	// Initialize routes
	initRouter(app)

	// Start server
	runServer(app)
}

// initConfig initializes configuration
func initConfig() {
	config.Init()
}

// initTemplate initializes templates
func initTemplate(app *gin.Engine) {
	app.LoadHTMLGlob("template/**/*")
}

// initStatic initializes static files
func initStatic(app *gin.Engine) {
	app.Static("/static", "./static")
}

// initRouter initializes routes
func initRouter(app *gin.Engine) {
	router.SetupRouter(app)
}

// runServer starts the server
func runServer(app *gin.Engine) {
	app.Run(config.GetString("server.port", ":8080"))
}
