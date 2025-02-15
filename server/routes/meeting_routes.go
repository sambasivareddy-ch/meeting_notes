package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

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

	getEventsRequest, err := http.NewRequest("GET", os.Getenv("GoogleCalenderEventsApi"), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create get events request",
		})
		return
	}

	getEventsRequest.Header.Add("Authorization", "Bearer "+sessionInfo.AccessToken)
	client := &http.Client{}
	response, err := client.Do(getEventsRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get events from Google API",
		})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read events response",
		})
		return
	}

	var meetingsList models.MeetingsList
	if err = json.Unmarshal(body, &meetingsList); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse events response",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meetings": meetingsList,
	})
}
