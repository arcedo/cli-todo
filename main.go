package main

import (
	"log"
	"os"

	"arcedo/cli-todo/internal/cli"
	"arcedo/cli-todo/internal/db"
	"arcedo/cli-todo/internal/task"
)

func main() {
	database, err := db.ConnectSqlite("cli-todo.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := task.NewSqliteRepository(database)
	service := task.NewService(repo)
	cli.Run(os.Args, service)
}
