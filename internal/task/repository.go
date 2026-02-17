package task

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, tasks []Task) error
	Delete(ctx context.Context, ids []int) (int, error)
	Get(ctx context.Context, ids []int, filter ListFilter) ([]Task, error)
	Complete(ctx context.Context, ids []int) (int, error)
}
