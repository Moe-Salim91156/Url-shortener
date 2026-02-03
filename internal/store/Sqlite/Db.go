package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlDb() (*sql.DB, error) {
	// we first open up a DB connection , and with its pointer passed to evertyhing
	// this pointer holds the connection of the DB
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	//create the tables for user, session, url , paste,

	// users table , should be first created , because foreign keeys will point to it
	usersTable := `CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	username TEXT NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	created_at TEXT NOT NULL
	);`
	sessionsTable := `
    CREATE TABLE IF NOT EXISTS sessions (
        session_id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        expires_at TEXT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	urlsTable := `
    CREATE TABLE IF NOT EXISTS urls (
        short_code TEXT PRIMARY KEY,
        long_url TEXT NOT NULL,
        owner_id TEXT NOT NULL,
        creation_time TEXT NOT NULL,
        FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	pastesTable := `
    CREATE TABLE IF NOT EXISTS pastes (
        short_code TEXT PRIMARY KEY,
        content TEXT NOT NULL,
        title TEXT,
        owner_id TEXT NOT NULL,
        creation_time TEXT NOT NULL,
        FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	if _, err := db.Exec(usersTable); err != nil {
		return err
	}
	if _, err := db.Exec(urlsTable); err != nil {
		return err
	}
	if _, err := db.Exec(sessionsTable); err != nil {
		return err
	}
	if _, err := db.Exec(pastesTable); err != nil {
		return err
	}
	return nil
}
