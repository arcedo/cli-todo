// Package db contains the configuration of the database
package db

import (
	"arcedo/cli-todo/internal/task"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(task.Task{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
