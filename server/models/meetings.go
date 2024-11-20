package models

import (
	"time"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
)

type Meeting struct {
	USER_ID             int    `binding:"required"`
	MEETING_ID          string `binding:"required"`
	MEETING_TITLE       string `binding:"required"`
	MEETING_NOTES       string
	MEETING_STARTTIME   time.Time `binding:"required"`
	MEETING_ENDTIME     time.Time
	MEETING_DESCRIPTION string `binding:"required"`
	MEETING_TYPE        string `binding:"required"`
}

func (meeting *Meeting) Save() error {
	insertCommand := `INSERT INTO MEETINGS VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	preparedStmt, err := database.AppDatabase.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = preparedStmt.Exec(
		meeting.USER_ID,
		meeting.MEETING_ID,
		meeting.MEETING_TITLE,
		meeting.MEETING_NOTES,
		meeting.MEETING_STARTTIME,
		meeting.MEETING_ENDTIME,
		meeting.MEETING_DESCRIPTION,
		meeting.MEETING_TYPE,
	)

	if err != nil {
		return err
	}

	return nil
}
