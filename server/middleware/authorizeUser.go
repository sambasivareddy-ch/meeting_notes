package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/sessions"
)

func AuthorizeUser(ctx *gin.Context) {
	user_session_id, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}

	session_id, err := sessions.RedisClient.Get(sessions.RedisContext, "session_id").Result()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized Session",
		})
		return
	}

	if user_session_id != session_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized Session",
		})
		return
	}

	ctx.Set("authenticated", true)
}
