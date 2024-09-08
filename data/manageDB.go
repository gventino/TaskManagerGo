package data

import (
	"database/sql"
	"fmt"
	"main/models"

	_ "github.com/mattn/go-sqlite3"
)

// inserts a user users table
func InsertUser(username string) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM users WHERE username=?", username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// If user not registered yet
	if !rows.Next() {
		// Insert user into db:
		statement, err := db.Prepare(`
		INSERT INTO users (username)
		VALUES (?);
		`)
		if err != nil {
			fmt.Println(err)
		}
		statement.Exec(username)
		defer statement.Close()
		return true
	}

	return false
}

// inserts a task tasks table
func InsertTask(name, desc, difficulty, deadline, owner string) int64 {
	// Open db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Insert user into db:
	result, err := db.Exec(`
	INSERT INTO tasks (name, desc, difficulty, deadline, is_done, owner)
	VALUES (?, ?, ?, ?, ?, ?);
	`, name, desc, difficulty, deadline, 0, owner)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return id
}

// verify if the user exists
func VerifyLogin(username string) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM users WHERE username=?", username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	return rows.Next()
}

// mark a task as done
func CompleteTask(id int) bool {
	// Open db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Insert user into db:
	statement, err := db.Prepare(`
	UPDATE tasks SET is_done = 1
	WHERE id = ?;
	`)
	if err != nil {
		return false
	}
	statement.Exec(id)
	defer statement.Close()

	return true
}

// get tasks from a user
func GetTasks(username string) []models.Task {
	// Open db:
	db, err := sql.Open("sqlite3", "taskManager.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Get the tasks from the tasks table:
	var tasks []models.Task

	rows, err := db.Query(`
	SELECT id, name, desc, difficulty, deadline 
		FROM tasks
		WHERE owner = ? AND is_done = 0;
	`, username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Desc, &task.Difficulty, &task.Deadline); err != nil {
			fmt.Println(err)
			return tasks
		}
		tasks = append(tasks, task)
	}
	rows.Close()
	return tasks
}
