package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBody struct {
	Token string `json:"token"`
}

func SetCookieRoute(ctx *gin.Context) {
	var body TokenBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request, pass token",
		})
		return
	}

	// ctx.SetCookie("session_id", body.Token, 24*60*60, "/", "meeting-notes-phi.vercel.app", true, true)

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "session_id",
		Value: body.Token,
		Path: "/",
		Expires: time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status": "session created",
	})
}
