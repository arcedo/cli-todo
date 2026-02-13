// Package task contains all the logic related with the tasks
package task

import (
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	CompletedAt time.Time `sql:"index" json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `sql:"index" json:"deleted_at"`
}
