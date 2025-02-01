package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/routes"
)

// Base home route
func baseHomeRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server saying 'Hello, World!'",
	})
}

func CORSMiddleware() gin.HandlerFunc {
	fmt.Println("Calling CORS middleware")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle OPTIONS method (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // Respond with 204 No Content
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize the Server
	httpServer := gin.Default()

	httpServer.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           24 * 60 * 60,
	}))

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
