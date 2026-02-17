// Package task contains all the logic related with the tasks
package task

import (
	"errors"
	"strings"
	"time"
)

type Task struct {
	// We could use gorm.Model that adds to the model
	// the ID as we have it and the fields CreatedAt, UpdatedAt and DeletedAt
	ID          uint       `gorm:"primary_key"`
	Description string     `gorm:"not null"`
	CompletedAt *time.Time `sql:"index"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	DeletedAt   *time.Time `sql:"index"`
}

var ErrEmptyDescription = errors.New("task description cannot be empty")

func (t Task) validate() error {
	if strings.TrimSpace(t.Description) == "" {
		return ErrEmptyDescription
	}
	return nil
}

type ListOption string

const (
	ListAll         ListOption = "all"
	ListUncompleted ListOption = "uncompleted"
	ListCompleted   ListOption = "completed"
)
