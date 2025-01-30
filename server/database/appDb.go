package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// App level Database
var AppDatabase *sql.DB

// InitDB() initialize the SQLite3 database & Create DB tables
func InitDB() error {
	var err error

	AppDatabase, err = sql.Open("sqlite3", "app_database.db")

	if err != nil {
		log.Fatalf("failed to connect the database")
	}

	createDatabaseTables()

	// Checks if DB connection is alive else returns err
	return AppDatabase.Ping()
}

// createDatabaseTables() creates a user table
func createDatabaseTables() {
	var prepared_statment *sql.Stmt
	var err error

	// Users table
	user_table_create_command := `CREATE TABLE IF NOT EXISTS USERS(
		USER_ID INTEGER PRIMARY KEY AUTOINCREMENT,
		USER_NAME VARCHAR(30) NOT NULL,
		EMAIL_ADDRESS TEXT NOT NULL
	)`

	// Meetings table
	meetings_table_create_command := `CREATE TABLE IF NOT EXISTS MEETINGS(
		USER_ID INT NOT NULL,
		MEETING_ID INT NOT NULL,
		MEETING_TITLE TEXT NOT NULL,
		MEETING_NOTES TEXT,
		MEETING_STARTTIME TIMESTAMP NOT NULL,
		MEETING_ENDTIME TIMESTAMP,
		MEETING_DESCRIPTION TEXT,
		MEETING_TYPE VARCHAR(30),
		PRIMARY KEY (MEETING_ID),
		FOREIGN KEY (USER_ID) REFERENCES USERS (USER_ID)
	)`

	// Proceeding with User Table creation
	prepared_statment, err = AppDatabase.Prepare(user_table_create_command)
	if err != nil {
		panic("unable to prepare the create users command")
	}
	if _, err = prepared_statment.Exec(); err != nil {
		panic("unable to create the users table")
	}

	// Proceeding with Meetings Table creation
	prepared_statment, err = AppDatabase.Prepare(meetings_table_create_command)
	if err != nil {
		panic("unable to prepare the create meetings command")
	}
	if _, err = prepared_statment.Exec(); err != nil {
		panic("unable to create the meetings table")
	}

	// DONE
}
