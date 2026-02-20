package main

import (
	"context"
	"log"
	"os"

	"arcedo/cli-todo/internal/cli"
	"arcedo/cli-todo/internal/db"
	"arcedo/cli-todo/internal/task"
)

func main() {
	ctx := context.Background()
	database, err := db.ConnectSqlite("cli-todo.db")
	if err != nil {
		log.Fatalf("failed to connect SQLite: %v", err)
	}
	if err = db.Migrate(database); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
	repo := task.NewSqliteRepository(database)
	service := task.NewService(repo)

	cli := cli.New(service, os.Stdout, os.Stderr)
	cli.Run(ctx, os.Args)
}
