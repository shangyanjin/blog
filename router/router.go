package router

import (
	"blog/config"
	"blog/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter(app *gin.Engine) {
	// Routes accessible without authentication
	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Blog Home",
		})
	})

}

func InitRouter(r *gin.Engine) {
	r.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())

	// setup virtual host by domain
	//r.Use(middleware.HostMiddleware())

	setupAdminRoute(r)
	setupApiRoute(r)
	setupUserRoute(r)
	setupWebRoutes(r)
	setupTemplate(r)
	//setupFile(r)
}

// NotFound handles gin NotFound error
func setupNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "404 Not Found"})
}

// MethodNotAllowed handles gin MethodNotAllowed error
func setupMethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusMethodNotAllowed, "error": "404 StatusMethodNotAllowed"})
}

// AccessForbidden handles Access Forbidden http code
func setupAccessForbidden(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusForbidden, "error": "403 StatusForbidden"})

}

// InternalError handles Internal Server Error http code
func setupInternalError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusInternalServerError, "error": "500 StatusForbidden"})
}

// NoRoute handles responses for undefined routes
func setupNoRoute(c *gin.Context) {
	// Define the paths and their img file directories
	pathMap := map[string]string{
		"/backend": "./data/www/vue/backend",
		"/pc":      "./data/www/vue/pc",
	}
	// Function to serve img files or a default index.html file
	serveStaticFile := func(basePath, baseDir string) {
		// Get the relative path from the request URL
		path := strings.TrimPrefix(c.Request.URL.Path, basePath)
		// Generate the full path to the img file
		fullPath := filepath.Join(baseDir, path)
		// Path to the default index.html file
		indexFile := filepath.Join(baseDir, "index.html")

		// Check if the img file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Static file does not exist, serve index.html
			c.File(indexFile)
			c.Abort()
		} else if err != nil {
			// Handle other errors
			logrus.Errorf("Error checking file %s: %v", fullPath, err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": err, "message": "Internal server error"})
			c.Abort()
		} else {
			// Static file exists, serve the file
			c.File(fullPath)
			c.Abort()
		}
	}

	// Iterate over the path map to find a matching base path
	for basePath, baseDir := range pathMap {
		if strings.HasPrefix(c.Request.URL.Path, basePath) {
			serveStaticFile(basePath, baseDir)
			return
		}
	}

	// If the request path does not match any base path, return a 404 error
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
}

// setupTemplate configures the HTML templates for the application

// initTemplate initializes templates and static files
func setupTemplate(r *gin.Engine) {

	// Setup template functions
	funcMap := GetTemplateFuncMap()
	r.SetFuncMap(funcMap)

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

	// Debug mode: Load templates from disk each time, subdir only
	r.LoadHTMLGlob(defaultTheme + "/**/*.html")

	// Serve static files from configured theme
	r.Static("/static", defaultTheme+"/static")
}
