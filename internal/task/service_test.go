package task_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"arcedo/cli-todo/internal/task"
)

// mocks
type mockRepository struct {
	createFunc   func(ctx context.Context, tasks []task.Task) error
	deleteFunc   func(ctx context.Context, ids []int) (int, error)
	getFunc      func(ctx context.Context, ids []int, filter task.ListFilter) ([]task.Task, error)
	completeFunc func(ctx context.Context, ids []int) (int, error)
}

func (m *mockRepository) Create(ctx context.Context, tasks []task.Task) error {
	return m.createFunc(ctx, tasks)
}

func (m *mockRepository) Delete(ctx context.Context, ids []int) (int, error) {
	return m.deleteFunc(ctx, ids)
}

func (m *mockRepository) Get(ctx context.Context, ids []int, filter task.ListFilter) ([]task.Task, error) {
	return m.getFunc(ctx, ids, filter)
}

func (m *mockRepository) Complete(ctx context.Context, ids []int) (int, error) {
	return m.completeFunc(ctx, ids)
}

// actual tests
func TestService_Create(t *testing.T) {
	ctx := context.Background()

	validTask := "Do homework"
	emptyTask := ""

	mock := &mockRepository{
		createFunc: func(ctx context.Context, tasks []task.Task) error {
			return nil
		},
	}
	svc := task.NewService(mock)

	t.Run("all valid tasks", func(t *testing.T) {
		tasks, err := svc.Create(ctx, []string{validTask})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tasks) != 1 || tasks[0].Description != validTask {
			t.Fatalf("unexpected tasks returned: %+v", tasks)
		}
	})

	t.Run("some invalid tasks", func(t *testing.T) {
		tasks, err := svc.Create(ctx, []string{validTask, emptyTask})
		if err == nil {
			t.Fatal("expected validation error, got nil")
		}
		if !strings.Contains(err.Error(), "validation errors") {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tasks) != 1 || tasks[0].Description != validTask {
			t.Fatalf("expected only valid task returned, got: %+v", tasks)
		}
	})
}

func TestService_Delete(t *testing.T) {
	ctx := context.Background()
	mock := &mockRepository{
		deleteFunc: func(ctx context.Context, ids []int) (int, error) {
			if len(ids) == 0 {
				return 0, errors.New("no IDs provided")
			}
			return len(ids), nil
		},
	}
	svc := task.NewService(mock)

	t.Run("successful delete", func(t *testing.T) {
		got, err := svc.Delete(ctx, []int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 2 {
			t.Fatalf("expected 2 affected, got %d", got)
		}
	})

	t.Run("delete error", func(t *testing.T) {
		_, err := svc.Delete(ctx, []int{})
		if err == nil || !strings.Contains(err.Error(), "failed to delete tasks") {
			t.Fatalf("expected wrapped delete error, got: %v", err)
		}
	})
}

func TestService_List(t *testing.T) {
	ctx := context.Background()
	mock := &mockRepository{
		getFunc: func(ctx context.Context, ids []int, filter task.ListFilter) ([]task.Task, error) {
			if len(ids) == 0 {
				return nil, errors.New("no tasks found")
			}
			return []task.Task{
				{Description: "Task1"},
				{Description: "Task2"},
			}, nil
		},
	}
	svc := task.NewService(mock)

	t.Run("successful list", func(t *testing.T) {
		tasks, err := svc.List(ctx, []int{1, 2}, task.All)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tasks) != 2 {
			t.Fatalf("expected 2 tasks, got %d", len(tasks))
		}
	})

	t.Run("list error", func(t *testing.T) {
		_, err := svc.List(ctx, []int{}, task.All)
		if err == nil || !strings.Contains(err.Error(), "failed to list tasks") {
			t.Fatalf("expected wrapped list error, got: %v", err)
		}
	})
}

func TestService_Complete(t *testing.T) {
	ctx := context.Background()
	mock := &mockRepository{
		completeFunc: func(ctx context.Context, ids []int) (int, error) {
			if len(ids) == 0 {
				return 0, errors.New("no tasks to complete")
			}
			return len(ids), nil
		},
	}
	svc := task.NewService(mock)

	t.Run("successful complete", func(t *testing.T) {
		got, err := svc.Complete(ctx, []int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 2 {
			t.Fatalf("expected 2 completed, got %d", got)
		}
	})

	t.Run("complete error", func(t *testing.T) {
		_, err := svc.Complete(ctx, []int{})
		if err == nil || !strings.Contains(err.Error(), "failed to complete tasks") {
			t.Fatalf("expected wrapped complete error, got: %v", err)
		}
	})
}
