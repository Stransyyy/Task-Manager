package dynamo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Database"))

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected")

	return db, err
}
