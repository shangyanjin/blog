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
	r := gin.Default()

	// Initialize templates
	initTemplate(r)

	// Initialize routes
	initRouter(r)

	// Start server
	runServer(r)
}

// initConfig initializes configuration
func initConfig() {
	config.Init()
}

// initTemplate initializes templates and static files
func initTemplate(r *gin.Engine) {
	// Get the default template directory from the configuration
	defaultTemplate := config.GetString("template.dir", "template")
	// Get the default theme from the configuration
	defaultTheme := defaultTemplate + "/" + config.GetString("template.theme", "default")

	//release mode: load templates with root and subdir
	//tmplFiles := template.New("")
	//if _, err := tmplFiles.ParseGlob(defaultTheme + "/*.html"); err != nil {
	//	logrus.Error("Error loading root templates:", err)
	//}
	//if _, err := tmplFiles.ParseGlob(defaultTheme + "/**/*.html"); err != nil {
	//	logrus.Error("Error loading subdir templates:", err)
	//}
	//r.SetHTMLTemplate(tmplFiles)

	//debug mode: load templates from disk every time and only support subdir
	r.LoadHTMLGlob(defaultTheme + "/**/*.html")

	// Serve static files from configured theme
	r.Static("/static", defaultTheme+"/static")
}

// initRouter initializes routes
func initRouter(r *gin.Engine) {
	router.SetupRouter(r)
}

// runServer starts the server
func runServer(r *gin.Engine) {
	r.Run(config.GetString("server.port", ":8080"))
}
