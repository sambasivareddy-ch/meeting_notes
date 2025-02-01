package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/models"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/sessions"
)

func GetUserEmailAddress(ctx *gin.Context) {
	value, isExists := ctx.Get("SessionInfo")
	if !isExists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Session Info doesn't exists",
		})
		return
	}

	sessionInfo, ok := value.(sessions.UserSessionInfo)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to assert UserSessionInfo",
		})
		return
	}

	userEmailAddr, err := models.GetUserEmailAddress(sessionInfo.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Useremail doesn't found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"email": userEmailAddr,
	})
}
