package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=joao2304 dbname=event_booking_database sslmode=disable")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date_time TIMESTAMP NOT NULL,
        user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
    )
    `
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

	createRegristrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
        id SERIAL PRIMARY KEY,
        event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (users) REFERENCES users(id),
    )
	`
	_, err = DB.Exec(createRegristrationTable)
	if err != nil {
		panic("could not create registration table")
	}
}
