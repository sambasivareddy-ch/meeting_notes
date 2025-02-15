package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/models"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/sessions"
)

func GetUserMeetingsRoute(ctx *gin.Context) {
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

	meetingList, err := models.GetMeetingsList(sessionInfo.UserId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert into meetings table",
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
		"email":    userEmailAddr,
		"meetings": meetingList,
	})
}

func UpdateMeetingNotesWithMeetingIdRoute(ctx *gin.Context) {
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

	meetingId := ctx.Param("meetingId")

	fmt.Println("Meeting ID: ", meetingId)
	meetingNotes := ctx.PostForm("meetingNotes")

	err := models.UpdateMeetingNotesWithMeetingId(meetingId, sessionInfo.UserId, meetingNotes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update meeting notes",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Meeting notes updated successfully",
	})
}
