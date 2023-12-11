package main

import (
	"fmt"

	db "github.com/Stransyyy/Task-Manager/mysql"
	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("run.env")

	db, err := db.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("")

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	task.Task_option()
}
