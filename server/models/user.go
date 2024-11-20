package models

import (
	"github.com/sambasivareddy-ch/meeting_notes_app/server/database"
)

type User struct {
	USER_ID       int
	USER_NAME     string `binding:"required"`
	EMAIL_ADDRESS string `binding:"required"`
}

func (user *User) Save() error {
	insertCommand := "INSERT INTO USERS (USER_NAME, EMAIL_ADDRESS) VALUES (?, ?)"

	preparedStmt, err := database.AppDatabase.Prepare(insertCommand)

	if err != nil {
		return err
	}

	defer preparedStmt.Close()

	if _, err = preparedStmt.Exec(user.USER_NAME, user.EMAIL_ADDRESS); err != nil {
		return err
	}

	return nil
}
