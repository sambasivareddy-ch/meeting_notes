package models

import (
	"time"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
)

type Session struct {
	USER_ID        int       `binding:"required"`
	SESSION_ID     string    `binding:"required"`
	ACCESS_TOKEN   string    `binding:"required"`
	REFRESH_TOKEN  string    `binding:"required"`
	CREATED_TIME   time.Time `binding:"required"`
	LAST_USED_TIME time.Time `binding:"required"`
	EXPIRES_IN     int       `binding:"required"`
}

func (session *Session) Save() error {
	insertCommand := `INSERT INTO SESSIONS VALUES (?, ?, ?, ?, ?, ?, ?)`

	preparedStmt, err := database.AppDatabase.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = preparedStmt.Exec(
		session.USER_ID,
		session.SESSION_ID,
		session.ACCESS_TOKEN,
		session.REFRESH_TOKEN,
		session.CREATED_TIME,
		session.LAST_USED_TIME,
		session.EXPIRES_IN,
	)

	if err != nil {
		return err
	}

	return nil
}
