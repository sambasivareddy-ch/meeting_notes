package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

func GetAccessTokenUsingCode(ctx *gin.Context) {
	var err error

	code := ctx.Query("code")
	client_id := utils.Client_ID
	client_secret := utils.Client_Secret
	redirect_uri := utils.Redirect_Uri

	reqData := url.Values{}
	reqData.Add("client_secret", client_secret)
	reqData.Add("client_id", client_id)
	reqData.Add("redirect_uri", redirect_uri)
	reqData.Add("code", code)
	reqData.Add("grant_type", "authorization_code")

	var tokenRequest *http.Request
	if tokenRequest, err = http.NewRequest(
		"POST",
		utils.Google_OAuth_Token_Uri,
		bytes.NewBufferString(reqData.Encode()),
	); err != nil {
		log.Fatalf("unable to create the token request")
	}

	tokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	var response *http.Response
	if response, err = client.Do(tokenRequest); err != nil {
		log.Fatalf("Error while sending token request to google token uri")
	}

	defer response.Body.Close()

	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Unable to read the response body")
		return
	}

	var oAuthResponse utils.OAuthResponse
	fmt.Println(json.Unmarshal(body, &oAuthResponse))

	// Step 4: Parse the response
	// var oAuthResponse utils.OAuthResponse
	// if err = json.NewDecoder(response.Body).Decode(&oAuthResponse); err != nil {
	// 	fmt.Println("Error parsing response:", err)
	// 	return
	// }

	ctx.Set("oAuthResponse", oAuthResponse)
}
