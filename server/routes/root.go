package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/middleware"
)

func RegisterRootRoute(server *gin.Engine) {
	// Authentication Group
	authenticateGroup := server.Group("/auth")
	authenticateGroup.Use(middleware.GetAccessTokenUsingCode)
	authenticateGroup.GET("/callback", CompleteGoogleAuthentication)
}
