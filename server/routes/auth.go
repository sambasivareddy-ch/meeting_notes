package routes

import (
	"encoding/json"
	"io"
	"net/http"

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
	newProfileRequest, err := http.NewRequest("GET", utils.User_Profile_Uri, nil)
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
		err = userInfo.SaveUser()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create the new user",
			})
			return
		}
	}

	newSessionId := sessions.GenerateSessionId(oAuthResponse.AccessToken)
	_, err = sessions.RedisClient.Set(sessions.RedisContext, "session_id", newSessionId, 24*60*60).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create a session",
		})
		return
	}

	ctx.SetCookie("session_id", newSessionId, 24*60*60, "/", "localhost", false, true)
	ctx.Redirect(http.StatusFound, "http://localhost:3000/my-meetings")
}
