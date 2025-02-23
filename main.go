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

	// Initialize routes
	initRouter(app)

	// Start server
	runServer(app)
}

// initConfig initializes configuration
func initConfig() {
	config.Init()
}

// initTemplate initializes templates and static files
func initTemplate(app *gin.Engine) {
	// Get the default template directory from the configuration
	defaultTemplate := config.GetString("template.dir", "template")
	// Get the default theme from the configuration
	defaultTheme := defaultTemplate + "/" + config.GetString("template.theme", "default")
	// Load HTML templates from configured theme directory
	app.LoadHTMLGlob(defaultTheme + "/**/*.html")

	// Serve static files from configured theme
	app.Static("/static", defaultTheme+"/static")
}

// initRouter initializes routes
func initRouter(app *gin.Engine) {
	router.SetupRouter(app)
}

// runServer starts the server
func runServer(app *gin.Engine) {
	app.Run(config.GetString("server.port", ":8080"))
}
