package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/routes"
	// "github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

// Base home route
func baseHomeRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server saying 'Hello, World!'",
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://meeting-notes-phi.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS method (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize the Server
	httpServer := gin.Default()

	// Load the environment variables
	// utils.LoadEnv()

	httpServer.Use(CORSMiddleware())

	// On app startup (initial) initialize the Database & create tables
	if err := database.InitDB(); err != nil {
		log.Fatal(err.Error())
	}

	httpServer.GET("/", baseHomeRouteHandler)

	routes.RegisterRootRoute(httpServer)

	// Running the server at "8080" port
	if err := httpServer.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
