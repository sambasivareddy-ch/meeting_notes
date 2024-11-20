package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

func CompleteGoogleAuthentication(ctx *gin.Context) {
	// user := models.User{}
	// session := models.Session{}

	value, exists := ctx.Get("oAuthResponse")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to authorize the gmail",
		})
	}

	oAuthResponse := value.(utils.OAuthResponse)

	newProfileRequest, err := http.NewRequest("GET", utils.User_Profile_Uri, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	newProfileRequest.Header.Add("Authorization", "Bearer "+oAuthResponse.AccessToken)

	client := &http.Client{}
	response, err := client.Do(newProfileRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to get profile from google api",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
