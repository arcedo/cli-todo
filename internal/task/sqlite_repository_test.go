package task_test

import (
	"context"
	"testing"

	"arcedo/cli-todo/internal/db"
	"arcedo/cli-todo/internal/task"

	"gorm.io/gorm"
)

func setupRepository(t *testing.T) (*gorm.DB, task.Repository) {
	database, err := db.ConnectSqlite(":memory:")
	if err != nil {
		t.Fatalf("failed to connect to sqlite: %v", err)
	}

	if err := db.Migrate(database); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}
	return database, task.NewSqliteRepository(database)
}

func seedTasks(t *testing.T, repo task.Repository, descriptions ...string) []task.Task {
	ctx := context.Background()
	tasks := make([]task.Task, len(descriptions))
	for i, desc := range descriptions {
		tasks[i] = task.Task{Description: desc}
	}

	if err := repo.Create(ctx, tasks); err != nil {
		t.Fatalf("failed to seed tasks: %v", err)
	}

	return tasks
}

func TestSqliteRepository_Create(t *testing.T) {
	db, repo := setupRepository(t)

	t.Run("creates tasks successfully", func(t *testing.T) {
		tasks := seedTasks(t, repo, "T1", "T2")

		for _, tk := range tasks {
			if tk.ID == 0 {
				t.Errorf("expected task %q ID not to be 0", tk.Description)
			}
		}

		var count int64
		db.Model(&task.Task{}).Count(&count)
		if count != 2 {
			t.Errorf("expected 2 tasks in DB, got %d", count)
		}
	})
}

func TestSqliteRepository_Delete(t *testing.T) {
	_, repo := setupRepository(t)
	ctx := context.Background()
	tasks := seedTasks(t, repo, "Task 1", "Task 2", "Task 3")

	t.Run("delete existing tasks", func(t *testing.T) {
		ids := []int{int(tasks[0].ID), int(tasks[1].ID)}
		deleted, err := repo.Delete(ctx, ids)
		if err != nil {
			t.Fatalf("failed to delete tasks: %v", err)
		}
		if deleted != len(ids) {
			t.Errorf("expected %d tasks deleted, got %d", len(ids), deleted)
		}
	})

	t.Run("delete non-existent task", func(t *testing.T) {
		deleted, err := repo.Delete(ctx, []int{999})
		if err != nil {
			t.Fatalf("unexpected error deleting non-existent task: %v", err)
		}
		if deleted != 0 {
			t.Errorf("expected 0 tasks deleted, got %d", deleted)
		}
	})

	t.Run("delete all remaining tasks", func(t *testing.T) {
		all, _ := repo.Get(ctx, nil, task.All)
		var ids []int
		for _, tk := range all {
			ids = append(ids, int(tk.ID))
		}
		deleted, _ := repo.Delete(ctx, ids)
		if deleted != len(ids) {
			t.Errorf("expected %d tasks deleted, got %d", len(ids), deleted)
		}
	})
}

func TestSqliteRepository_Complete(t *testing.T) {
	_, repo := setupRepository(t)
	ctx := context.Background()
	tasks := seedTasks(t, repo, "Task A", "Task B")

	t.Run("complete single task", func(t *testing.T) {
		rows, err := repo.Complete(ctx, []int{int(tasks[0].ID)})
		if err != nil {
			t.Fatalf("failed to complete task: %v", err)
		}
		if rows != 1 {
			t.Errorf("expected 1 row updated, got %d", rows)
		}
	})

	t.Run("complete already completed task", func(t *testing.T) {
		rows, err := repo.Complete(ctx, []int{int(tasks[0].ID)})
		if err != nil {
			t.Fatalf("unexpected error completing task: %v", err)
		}
		if rows != 0 {
			t.Errorf("expected 0 rows updated for already completed task, got %d", rows)
		}
	})

	t.Run("complete multiple tasks including one already completed", func(t *testing.T) {
		rows, err := repo.Complete(ctx, []int{int(tasks[0].ID), int(tasks[1].ID)})
		if err != nil {
			t.Fatalf("failed to complete tasks: %v", err)
		}
		if rows != 1 {
			t.Errorf("expected 1 row updated (Task B), got %d", rows)
		}
	})
}

func TestSqliteRepository_Get(t *testing.T) {
	_, repo := setupRepository(t)
	ctx := context.Background()
	tasks := seedTasks(t, repo, "Task X", "Task Y", "Task Z")

	// complete Task X
	_, err := repo.Complete(ctx, []int{int(tasks[0].ID)})
	if err != nil {
		t.Fatalf("failed to complete task")
	}

	t.Run("get all tasks", func(t *testing.T) {
		all, err := repo.Get(ctx, nil, task.All)
		if err != nil {
			t.Fatalf("failed to get all tasks: %v", err)
		}
		if len(all) != 3 {
			t.Errorf("expected 3 tasks, got %d", len(all))
		}
	})

	t.Run("get completed tasks", func(t *testing.T) {
		completed, err := repo.Get(ctx, nil, task.Completed)
		if err != nil {
			t.Fatalf("failed to get completed tasks: %v", err)
		}
		if len(completed) != 1 {
			t.Errorf("expected 1 completed task, got %d", len(completed))
		}
		if completed[0].ID != tasks[0].ID {
			t.Errorf("completed task ID mismatch, expected %d, got %d", tasks[0].ID, completed[0].ID)
		}
	})

	t.Run("get uncompleted tasks", func(t *testing.T) {
		uncompleted, err := repo.Get(ctx, nil, task.Uncompleted)
		if err != nil {
			t.Fatalf("failed to get uncompleted tasks: %v", err)
		}
		if len(uncompleted) != 2 {
			t.Errorf("expected 2 uncompleted tasks, got %d", len(uncompleted))
		}
	})

	t.Run("get by specific IDs", func(t *testing.T) {
		specific, err := repo.Get(ctx, []int{int(tasks[1].ID)}, task.All)
		if err != nil {
			t.Fatalf("failed to get specific task: %v", err)
		}
		if len(specific) != 1 || specific[0].ID != tasks[1].ID {
			t.Errorf("expected task ID %d, got %v", tasks[1].ID, specific)
		}
	})
}
