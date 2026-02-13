package main

import (
	"log"
	"os"

	"arcedo/cli-todo/internal/db"
	"arcedo/cli-todo/internal/task"
)

func main() {
	database, err := db.ConnectSqlite("cli-todo.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := task.NewSqliteRepository(database)
	handler := task.NewHandler(repo)

	if err := handler.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
