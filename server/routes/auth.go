package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/models"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/sessions"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

func CompleteGoogleAuthentication(ctx *gin.Context) {
	value, exists := ctx.Get("oAuthResponse")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to authorize the gmail",
		})
		return
	}

	oAuthResponse, ok := value.(utils.OAuthResponse)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid OAuth response",
		})
		return
	}

	// Verified the Token, now create a request for concerned user profile
	newProfileRequest, err := http.NewRequest("GET", os.Getenv("UserProfileUri"), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create profile request",
		})
		return
	}

	newProfileRequest.Header.Add("Authorization", "Bearer "+oAuthResponse.AccessToken)

	// Now request for the profile
	client := &http.Client{}
	response, err := client.Do(newProfileRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to get profile from Google API",
		})
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read profile response",
		})
		return
	}

	var userInfo models.UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to unmarshal profile response",
		})
		return
	}

	isUserExists, err := userInfo.IsUserAlreadyExists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to fetch the user",
		})
		return
	}

	if !isUserExists {
		// Save the user info in the Users table
		err = userInfo.SaveUser(oAuthResponse.AccessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create the new user",
			})
			return
		}
	} else {
		// Just Update the access token
		err = userInfo.UpdateUsersAccessToken(oAuthResponse.AccessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update the access token",
			})
			return
		}
	}

	newSessionId := sessions.GenerateSessionId(oAuthResponse.AccessToken)
	newSessionInfo := &sessions.UserSessionInfo{
		UserId:      userInfo.Id,
		AccessToken: oAuthResponse.AccessToken,
	}

	sessionJsonObject, err := json.Marshal(newSessionInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create session info",
		})
		return
	}

	_, err = sessions.RedisClient.Set(sessions.RedisContext, newSessionId, sessionJsonObject, 24*time.Hour).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create a session",
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

	getEventsRequest.Header.Add("Authorization", "Bearer "+newSessionInfo.AccessToken)
	response, err = client.Do(getEventsRequest)
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

	_, err = models.InsertIntoMeetingsTable(meetingsList, newSessionInfo.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert into meetings table",
		})
		return
	}

	ctx.SetCookie("session_id", newSessionId, 24*60*60, "/", "meeting-notes-phi.vercel.app", true, true)
	ctx.Redirect(http.StatusFound, "https://meeting-notes-phi.vercel.app/")
}

func LogoutRoute(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Session not found",
		})
		return
	}

	_, err = sessions.RedisClient.Del(sessions.RedisContext, sessionId).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete the session",
		})
		return
	}

	ctx.SetCookie("session_id", "", -1, "/", "meeting-notes-phi.vercel.app", true, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Logged out successfully",
		"isLoggedOut": true,
	})
}

func LoginStatusRoute(ctx *gin.Context) {
	_, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"isLoggedIn": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
	})
}
