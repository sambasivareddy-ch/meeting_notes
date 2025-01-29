package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err == nil {
		fmt.Println(prettyJSON.String())
	}

	ctx.Redirect(http.StatusFound, "/my-meetings")
}
