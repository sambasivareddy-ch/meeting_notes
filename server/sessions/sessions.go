package sessions

import (
	"crypto/sha256"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/sambasivareddy-ch/meeting_notes_app/server/utils"
)

var Store = sessions.NewCookieStore([]byte(utils.GenerateSecretKey()))

func GenerateSessionId(access_token string) string {
	hasher := sha256.New()
	hasher.Write([]byte(access_token))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash
}
