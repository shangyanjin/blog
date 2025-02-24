package main

import (
	"blog/config"
	"blog/model"
	"blog/pkg/cache"
	"blog/pkg/xdaemon"
	"blog/router"

	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	initLogger()
	// Initialize configuration
	initConfig()
	// Initialize database
	initDb()
	// Initialize cache
	initCache()
	// Run as daemon if specified
	runDaemon()
	// Run delayed tasks
	runDelay()
	// Initialize and run the Gin router
	runRouter()

}

// initLogger initializes logrus logger with some defaults
func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// initConfig initializes the application configuration
func initConfig() {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Failed to initialize configuration: %v", err)
	}

}

// initCache initializes the cache system
func initCache() {
	if err := cache.InitCache(); err != nil {
		logrus.Fatalf("Failed to initialize cache: %v", err)
	}
}

// initDb initializes the database connection
func initDb() {
	if err := model.InitDB(); err != nil {
		logrus.Fatalf("Failed to initialize database: %v", err)
	}
}

// resetPassword resets admin account credentials
func resetPassword() {
	admin := model.User{
		Id:       1,
		Account:  "admin",           // Default admin account
		Email:    "admin@admin.com", // Default admin email
		Password: "admin",           // Default admin password
		Status:   "1",               // Enable admin account
	}
	// Save will automatically create or update based on ID
	if err := model.DB.Save(&admin).Error; err != nil {
		logrus.Fatalf("Failed to reset admin credentials: %v", err)
	}

	logrus.Infof("Admin credentials reset successfully - Account: %s, Email: %s, Password: %s",
		admin.Account, admin.Email, admin.Password)
}

// runRouter initializes and runs the Gin router
func runRouter() {
	// Get gin mode from config, default to ReleaseMode
	mode := config.GetString("server.mode", gin.ReleaseMode)
	gin.SetMode(mode)

	// Create Gin instance
	r := gin.Default()

	// Initialize router routes
	router.InitRouter(r)

	// Start HTTP server
	httpPort := config.GetInt("server.port", 80)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: r,
	}
	logrus.Infof("Starting HTTP server on port %d in %s mode", httpPort, mode)
	if err := server.ListenAndServe(); err != nil {
		logrus.Errorf("HTTP server failed: %v", err)
	}
}

// runDaemon parses the command line flags and runs the application as a daemon process if specified
func runDaemon() {
	// Default daemon flag
	isDaemon := false
	isReset := false

	// Manually parse command line arguments to handle /d, /debug, /release, and /r
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-d", "/d":
			isDaemon = true
		case "-r", "/r":
			isReset = true
		}
	}

	// Reset password if requested
	if isReset {
		resetPassword()
		os.Exit(0) // Exit after resetting password
	}

	// If the daemon flag is set, run the application as a daemon
	if isDaemon {
		logFile := "daemon.log"
		daemon := xdaemon.NewDaemon(logFile)
		daemon.MaxCount = 2 // Maximum restart attempts
		daemon.Run()
	}
}

// runDelay executes tasks once after a one-minute delay
func runDelay() {
	go func() {
		time.Sleep(10 * time.Second)
		logrus.Info("tasks to run in 10 Second...done")
	}()
}
