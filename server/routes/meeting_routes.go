package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/models"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/sessions"
)

type NotesBody struct {
	Notes string `json:"notes"`
}

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

	var submittedNotes NotesBody
	err := ctx.BindJSON(&submittedNotes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err = models.UpdateMeetingNotesWithMeetingId(meetingId, sessionInfo.UserId, submittedNotes.Notes)
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

func GetNotesForMeetingIdRoute(ctx *gin.Context) {
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

	meetingNotes, err := models.GetMeetingNotesWithMeetingId(meetingId, sessionInfo.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get meeting notes",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"notes": meetingNotes,
	})
}

func ReloadMeetingsRoute(ctx *gin.Context) {
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

	eventsResponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read events response",
		})
		return
	}

	var meetingsList models.MeetingsList
	if err = json.Unmarshal(eventsResponseBody, &meetingsList); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse events response",
		})
		return
	}

	_, err = models.InsertIntoMeetingsTable(meetingsList, sessionInfo.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert into meetings table",
		})
		return
	}

	reloadedMeetings, err := models.GetMeetingsList(sessionInfo.UserId)

	ctx.JSON(http.StatusOK, gin.H{
		"meetings": reloadedMeetings,
		"message":  "Meetings reloaded successfully",
	})
}
