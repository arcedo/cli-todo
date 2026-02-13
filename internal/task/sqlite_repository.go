package task

import (
	"context"

	"gorm.io/gorm"
)

type SqliteRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) Repository {
	return &SqliteRepository{db}
}

func (r *SqliteRepository) Create(ctx context.Context, task *[]Task) ([]int, error) {
	return []int{0}, nil
}

func (r *SqliteRepository) Delete(ctx context.Context, task *[]Task) error {
	return nil
}

func (r *SqliteRepository) Get(ctx context.Context, ids []int, params []string) ([]Task, error) {
	return []Task{{}}, nil
}

func (r *SqliteRepository) Complete(ctx context.Context, ids []int) error {
	return nil
}
