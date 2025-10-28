package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=cay password=491014 dbname=apidb host=localhost port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	// Tes koneksi
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Could not ping database: %v", err))
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
	);`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		user_id INTEGER REFERENCES users(id)
	);`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		event_id INTEGER REFERENCES events(id),
		user_id INTEGER REFERENCES users(id)
	);`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic(err)
	}
}
