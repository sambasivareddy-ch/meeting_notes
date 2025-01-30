package models

import (
	"errors"
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

func (user UserInfo) SaveUser() error {
	insertCommand := `INSERT INTO USERS VALUES (?, ?, ?)`

	preparedStatement, err := database.AppDatabase.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = preparedStatement.Exec(
		user.Id,
		user.Name,
		user.Email,
	)

	if err != nil {
		return err
	}

	// Successfully Saved
	return nil
}

func (usr UserInfo) IsUserAlreadyExists() (bool, error) {
	searchQuery := `SELECT USER_ID FROM USERS WHERE USER_ID = ?`

	preparedStmt, err := database.AppDatabase.Prepare(searchQuery)
	if err != nil {
		return false, err
	}

	result, err := preparedStmt.Exec(
		usr.Id,
	)
	if err != nil {
		return false, err 
	}

	returnedRowsCount, err := result.RowsAffected()
	if err != nil {
		return false, err 
	}

	if returnedRowsCount != 1 {
		return false, errors.New("duplicate users exists") 
	}

	return true, nil
}