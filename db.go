package main

import (
	"database/sql"
	"fmt"
)

const schema = `
	CREATE TABLE IF NOT EXISTS journal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		created_at DATE NOT NULL,
		deleted_at DATE
	);

	CREATE TABLE IF NOT EXISTS entry (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(128) NOT NULL,
		body TEXT,
		mood VARCHAR(64),
		created_at DATE NOT NULL,
		deleted_at DATE,
		journal_id INTEGER NOT NULL REFERENCES journal(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		entry_id INTEGER NOT NULL REFERENCES entry(id) ON DELETE CASCADE
	);
`

// OpenDB opens database connection and creates tables from schema
func OpenDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./journal.db")
	if err != nil {
		return
	}

	_, err = db.Exec(schema)
	if err != nil {
		return
	}

	return
}

// OpenTestDB opens memory database connection and creates tables from schema
func OpenTestDB(name string) (db *sql.DB, err error) {
	file := fmt.Sprintf("file:%v.db?cache=shared&mode=memory", name)
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		return
	}

	_, err = db.Exec(schema)
	if err != nil {
		return
	}

	return
}
