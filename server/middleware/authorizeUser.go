package middleware

import (
	"encoding/json"
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

	session_info, err := sessions.RedisClient.Get(sessions.RedisContext, user_session_id).Result()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized Session",
		})
		return
	}

	if session_info == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized Session",
		})
		return
	}

	var retrievedSessionInfo sessions.UserSessionInfo
	err = json.Unmarshal([]byte(session_info), &retrievedSessionInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error happened",
		})
		return
	}

	ctx.Set("authenticated", true)
	ctx.Set("SessionInfo", retrievedSessionInfo)
}
