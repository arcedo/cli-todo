// Package task contains all the logic related with the tasks
package task

import (
	"gorm.io/datatypes"
)

type Task struct {
	// We could use gorm.Model that adds to the model
	// the ID as we have it and the fields CreatedAt, UpdatedAt and DeletedAt
	ID          uint `gorm:"primary_key"`
	Description string
	CompletedAt datatypes.Date `sql:"index"`
	CreatedAt   datatypes.Date
	DeletedAt   datatypes.Date `sql:"index"`
}
