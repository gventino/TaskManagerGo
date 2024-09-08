package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func StartDB() {
	// Opening db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}

	// Creating users table:
	statement, err := db.Prepare(`
	CREATE TABLE users (
		username TEXT PRIMARY KEY
	);
	`)
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec()
	defer statement.Close()

	// Creating tasks table:
	statementTasks, errTasks := db.Prepare(`
	CREATE TABLE tasks (
		id INTEGER PRIMARY KEY,
		name TEXT,
		desc TEXT,
		difficulty TEXT,
		deadline TEXT,
		is_done INTEGER,
		owner TEXT,
		FOREIGN KEY(owner) REFERENCES users(username)
	);
	`)
	if errTasks != nil {
		fmt.Println(err)
	}
	statementTasks.Exec()
	defer statementTasks.Close()

	// Closing db:
	defer db.Close()
}
