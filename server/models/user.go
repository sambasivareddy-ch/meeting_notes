package models

import (
	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
)

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func (user UserInfo) SaveUser(accessToken string) error {
	insertCommand := `INSERT INTO USERS VALUES (?, ?, ?, ?)`

	preparedStatement, err := database.AppDatabase.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = preparedStatement.Exec(
		user.Id,
		user.Name,
		user.Email,
		accessToken,
	)

	if err != nil {
		return err
	}

	// Successfully Saved
	return nil
}

func (usr UserInfo) IsUserAlreadyExists() (bool, error) {
	searchQuery := `SELECT count(*) FROM USERS WHERE USER_ID = ?`
	var count int

	preparedStmt, err := database.AppDatabase.Prepare(searchQuery)
	if err != nil {
		return false, err
	}

	err = preparedStmt.QueryRow(usr.Id).Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (usr UserInfo) UpdateUsersAccessToken(newAccessToken string) error {
	updateQuery := `UPDATE USERS SET ACCESS_TOKEN = ? WHERE USER_ID = ?`

	preparedStmt, err := database.AppDatabase.Prepare(updateQuery)
	if err != nil {
		return err
	}

	_, err = preparedStmt.Exec(newAccessToken, usr.Id)
	if err != nil {
		return err
	}

	return nil
}