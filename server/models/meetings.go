package models

import (
	"fmt"
	"time"

	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
)

type Timings struct {
	DateTime time.Time `json:"dateTime"`
	TimeZone string    `json:"timeZone"`
}

type Organizer struct {
	Email string `json:"email"`
	Self  bool   `json:"self"`
}

type Meeting struct {
	Id                string    `json:"id"`
	Meeting_Title     string    `json:"summary"`
	Meeting_StartTime Timings   `json:"start"`
	Meeting_EndTime   Timings   `json:"end"`
	Meeting_Link      string    `json:"hangoutLink"`
	Meeting_Organizer Organizer `json:"organizer"`
}

type MeetingsList struct {
	Meetings []Meeting `json:"items"`
}

type MeetingFromTable struct {
	Id                string    `json:"id"`
	Meeting_Title     string    `json:"summary"`
	Meeting_StartTime Timings   `json:"start"`
	Meeting_EndTime   Timings   `json:"end"`
	Meeting_Notes     string    `json:"notes"`
	Meeting_Link      string    `json:"hangoutLink"`
	Meeting_Organizer Organizer `json:"organizer"`
}

type MeetingsListFromTable struct {
	Meetings []MeetingFromTable `json:"items"`
}

func (meeting *Meeting) Save(user_id string) error {
	insertCommand := `INSERT INTO MEETINGS (USER_ID, MEETING_ID, MEETING_TITLE, MEETING_STARTTIME, MEETING_LINK, MEETING_ORGANIZER) VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING`

	preparedStmt, err := database.AppDatabase.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = preparedStmt.Exec(
		user_id,
		meeting.Id,
		meeting.Meeting_Title,
		meeting.Meeting_StartTime.DateTime,
		meeting.Meeting_Link,
		meeting.Meeting_Organizer.Email,
	)

	if err != nil {
		return err
	}

	return nil
}

func InsertIntoMeetingsTable(meetingsList MeetingsList, user_id string) (MeetingsList, error) {
	var updateMeetingList MeetingsList
	for _, meeting := range meetingsList.Meetings {
		if meeting.Meeting_Link != "" {
			err := meeting.Save(user_id)
			updateMeetingList.Meetings = append(updateMeetingList.Meetings, meeting)
			if err != nil {
				fmt.Println("Error while saving meeting: ", err)
				return updateMeetingList, err
			}
		}
	}

	return updateMeetingList, nil
}

func GetMeetingsList(user_id string) (MeetingsListFromTable, error) {
	var meetingsList MeetingsListFromTable
	selectCommand := `SELECT MEETING_ID, MEETING_TITLE, MEETING_STARTTIME, COALESCE(MEETING_NOTES, '') AS MEETING_NOTES, MEETING_LINK, MEETING_ORGANIZER FROM MEETINGS WHERE USER_ID = ?`

	rows, err := database.AppDatabase.Query(selectCommand, user_id)
	if err != nil {
		return meetingsList, err
	}

	for rows.Next() {
		var meeting MeetingFromTable
		err = rows.Scan(
			&meeting.Id,
			&meeting.Meeting_Title,
			&meeting.Meeting_StartTime.DateTime,
			&meeting.Meeting_Notes,
			&meeting.Meeting_Link,
			&meeting.Meeting_Organizer.Email,
		)
		if err != nil {
			return meetingsList, err
		}

		meetingsList.Meetings = append(meetingsList.Meetings, meeting)
	}

	return meetingsList, nil
}

func UpdateMeetingNotesWithMeetingId(meetingId string, user_id string, meetingNotes string) error {
	updateCommand := `UPDATE MEETINGS SET MEETING_NOTES = ? WHERE MEETING_ID = ? AND USER_ID = ?`

	preparedStmt, err := database.AppDatabase.Prepare(updateCommand)
	if err != nil {
		return err
	}

	_, err = preparedStmt.Exec(meetingNotes, meetingId, user_id)
	if err != nil {
		return err
	}

	return nil
}
