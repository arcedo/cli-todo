package task

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type SqliteRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) Repository {
	return &SqliteRepository{db}
}

func (r *SqliteRepository) Create(ctx context.Context, tasks []Task) error {
	return r.db.WithContext(ctx).Create(&tasks).Error
}

func (r *SqliteRepository) Delete(ctx context.Context, ids []int) (rowsDeleted int, err error) {
	rowsDeleted, err = gorm.G[Task](r.db).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Update(ctx, "deleted_at", time.Now())
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func (r *SqliteRepository) Get(ctx context.Context, ids []int, filter ListFilter) (tasks []Task, err error) {
	db := r.db.WithContext(ctx)
	switch filter {
	case IDs:
		db = db.Where("id IN ?", ids)
	case All:
		db = db.Where("deleted_at IS NULL")
	case Completed:
		db = db.Where("completed_at IS NOT NULL AND deleted_at IS NULL")
	case Uncompleted:
		db = db.Where("completed_at IS NULL AND deleted_at IS NULL")
	case Removed:
		db = db.Where("deleted_at IS NOT NULL")
	}

	if err = db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *SqliteRepository) Complete(ctx context.Context, ids []int) (rowsCompleted int, err error) {
	db := gorm.G[Task](r.db)

	rowsCompleted, err = db.
		Where("id IN ? AND completed_at IS NULL", ids).
		Update(ctx, "completed_at", time.Now())
	if err != nil {
		return 0, err
	}

	return rowsCompleted, nil
}
