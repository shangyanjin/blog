package main

import (
	"blog/config"
	"blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize config
	config.Init()

	// Create gin engine
	app := gin.Default()

	// Load templates
	app.LoadHTMLGlob("template/**/*")

	// Static files
	app.Static("/static", "./static")

	// Setup routes
	router.SetupRouter(app)

	// Run server
	app.Run(config.AppConfig.ServerPort)
}
