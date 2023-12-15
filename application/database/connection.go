package mysql

import (
	"database/sql"
	"fmt"
	"os"

	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db      *sql.DB
	retries int
}

func (db Storage) GetAll() ([]*task.Task, error) {
	rows, err := db.db.Query("SELECT taskID, title, description, due_date, completed FROM storage")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task

	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (db Storage) Store(t *task.Task) error {

	_, err := db.db.Exec("INSERT INTO storage (title, description, due_date, completed) VALUES (?, ?, ?, ?)", t.Title, t.Description, t.DueDate, t.Completed)
	return err
}

func (db Storage) MarkCompleted(id int) error {
	_, err := db.db.Exec("UPDATE `storage` SET completed = true WHERE taskID = ?", id)
	return err
}

func (db Storage) Delete(id int) error {
	_, err := db.db.Exec("DELETE FROM storage WHERE taskID = ?", id)
	return err
}

func (db Storage) Get(id int) (*task.Task, error) {
	var t task.Task
	err := db.db.QueryRow("SELECT taskID, title, description, due_date, completed FROM storage WHERE taskID = ?", id).Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (db Storage) Edit(t *task.Task) error {
	_, err := db.db.Exec("UPDATE `storage` SET title=?, description=?, due_date=?, completed=? WHERE taskID=?", t.Title, t.Description, t.DueDate, t.Completed, t.ID)
	return err
}

func (d *Storage) Open() error {

	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Database"))

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d Storage) Close() error {
	return d.db.Close()
}

func New() *Storage {
	return &Storage{}
}
