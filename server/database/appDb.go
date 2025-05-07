package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "samba-17793"
// 	password = ""
// 	dbname   = "postgres"
// )

// App level Database
var AppDatabase *sql.DB

// InitDB() initialize the SQLite3 database & Create DB tables
func InitDB() error {
	var err error

	dsn := "postgresql://postgres:" + os.Getenv("PGPASSWORD") + "@caboose.proxy.rlwy.net:48713/railway"
	AppDatabase, err = sql.Open("postgres", dsn)

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
		USER_ID TEXT PRIMARY KEY,
		USER_NAME TEXT NOT NULL,
		EMAIL_ADDRESS TEXT NOT NULL,
		ACCESS_TOKEN TEXT NOT NULL
	)`

	// Meetings table
	meetings_table_create_command := `CREATE TABLE IF NOT EXISTS MEETINGS(
		USER_ID TEXT NOT NULL,
		MEETING_ID TEXT NOT NULL,
		MEETING_TITLE TEXT NOT NULL,
		MEETING_NOTES TEXT,
		MEETING_STARTTIME TIMESTAMP NOT NULL,
		MEETING_LINK TEXT NOT NULL,
		MEETING_ORGANIZER TEXT NOT NULL,
		PRIMARY KEY (MEETING_ID),
		FOREIGN KEY (USER_ID) REFERENCES USERS (USER_ID)
	)`

	// Proceeding with User Table creation
	prepared_statment, err = AppDatabase.Prepare(user_table_create_command)
	if err != nil {
		fmt.Print(err)
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
