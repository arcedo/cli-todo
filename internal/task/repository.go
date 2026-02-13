package task

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, task *[]Task) ([]int, error)
	Delete(ctx context.Context, task *[]Task) error
	Get(ctx context.Context, ids []int, params []string) ([]Task, error)
	Complete(ctx context.Context, ids []int) error
}
