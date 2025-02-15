package utils

import (
	"os"
)

func LoadEnv() {
	// Load the environment variables
	os.Setenv("ClientID", "664085272562-5oscsb6e238a5bvu92nbbs3di2ajtdt6.apps.googleusercontent.com")
	os.Setenv("ClientSecret", "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")
	os.Setenv("GoogleOAuthTokenUri", "https://oauth2.googleapis.com/token")
	os.Setenv("RedirectUri", "http://localhost:8080/auth/callback")
	os.Setenv("UserProfileUri", "https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	os.Setenv("GoogleCalenderEventsApi", "https://www.googleapis.com/calendar/v3/calendars/primary/events")
}
