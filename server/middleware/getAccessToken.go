package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

func GetAccessTokenUsingCode(ctx *gin.Context) {
	code := ctx.Query("code")

	// No code received from the request
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing authorization code",
		})
		return
	}

	// Append all the required information of client
	reqData := url.Values{}
	reqData.Add("client_secret", utils.Client_Secret)
	reqData.Add("client_id", utils.Client_ID)
	reqData.Add("redirect_uri", utils.Redirect_Uri)
	reqData.Add("code", code)
	reqData.Add("grant_type", "authorization_code")

	// Now create a request to Google OAuth URI for generating the token
	tokenRequest, err := http.NewRequest("POST", utils.Google_OAuth_Token_Uri, bytes.NewBufferString(reqData.Encode()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create token request",
		})
		return
	}

	tokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Client HTTP to make request
	client := &http.Client{}

	// Now do the request to get the token
	response, err := client.Do(tokenRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending token request",
		})
		return
	}
	defer response.Body.Close()

	//Now parse the response and store in OAuthResponse Struct to pass to next
	var oAuthResponse utils.OAuthResponse
	if err = json.NewDecoder(response.Body).Decode(&oAuthResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parsing token response",
		})
		return
	}

	ctx.Set("oAuthResponse", oAuthResponse)
}
