package routes

import (
	"net/http"

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

	ctx.SetCookie("session_id", body.Token, 24*60*60, "/", "meeting-notes-phi.vercel.app", true, true)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "session created",
	})
}
