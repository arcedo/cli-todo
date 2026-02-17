package db

import (
	"arcedo/cli-todo/internal/task"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&task.Task{},
	)
}
