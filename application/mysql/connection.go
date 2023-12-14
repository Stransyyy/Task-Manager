package mysql

import (
	"database/sql"

	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db      *sql.DB
	retries int
}

func (db Storage) GetAll() ([]*task.Task, error) {
	rows, err := db.db.Query("SELECT id, title, description, due_date, completed FROM tasks")
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
	_, err := db.db.Exec("INSERT INTO tasks (title, description, due_date, completed) VALUES (?, ?, ?, ?)", t.Title, t.Description, t.DueDate, t.Completed)
	return err
}

func (db Storage) MarkCompleted(id int) error {
	_, err := db.db.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
	return err
}

func (db Storage) Delete(id int) error {
	_, err := db.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (db Storage) Get(id int) (*task.Task, error) {
	var t task.Task
	err := db.db.QueryRow("SELECT id, title, description, due_date, completed FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (d *Storage) Open() error {
	db, err := sql.Open("mysql", "root:password@tcp(")
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
