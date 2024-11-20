package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/routes"
)

// Base home route
func baseHomeRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server saying 'Hello, World!'",
	})
}

func main() {
	// Initialize the Server
	httpServer := gin.Default()

	// On app startup (initial) initialize the Database & create tables
	if err := database.InitDB(); err != nil {
		log.Fatal(err.Error())
	}

	httpServer.GET("/", baseHomeRouteHandler)

	routes.RegisterRootRoute(httpServer)

	// Running the server at "8080" port
	if err := httpServer.Run(":8080"); err != nil {
		log.Fatalf(err.Error())
	}
}
