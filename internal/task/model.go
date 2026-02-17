// Package task contains all the logic related with the tasks
package task

import (
	"errors"
	"fmt"
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
		return fmt.Errorf("task '%s': %w", t.Description, ErrEmptyDescription)
	}
	return nil
}

type ListFilter string

const (
	All         ListFilter = "all"
	Uncompleted ListFilter = "uncompleted"
	Completed   ListFilter = "completed"
)

// Maybe in a future we add order by value commands
/*
type ListOrderValue string

const (
	CreatedAt   ListOrderValue = "created"
	DeletedAt   ListOrderValue = "deleted"
	CompletedAt ListOrderValue = "completed"
)*/
