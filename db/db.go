package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

type Task struct {
	Title  string `json:"title"`
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func CreateTask(task Task) (bool, error) {
	fmt.Printf("Creating task with title: %s\n", task.Title)
	stmt, err := db.Prepare("INSERT INTO tasks(title, status) VALUES(?, 'todo')")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetTasks() ([]Task, error) {
	query := "SELECT id, title, status FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Status); err != nil {
			fmt.Printf("error when get tasks: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTask(task Task) (bool, error) {
	fmt.Printf("Updating task with title: %s\n", task.Title)
	stmt, err := db.Prepare("UPDATE tasks SET title = ?, status = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Status, task.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func InitDB() *sql.DB {
	var err error
	fmt.Println("Initializing database...")
	db, err = sql.Open("sqlite3", "./tasks.db")

	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, status TEXT DEFAULT 'todo')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database initialized.")
	return db
}
