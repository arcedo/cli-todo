package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"arcedo/cli-todo/internal/task"
)

// ------------------------
// Mock repository
// ------------------------

type mockRepo struct{}

func (m *mockRepo) Create(ctx context.Context, tasks []task.Task) error {
	return nil
}

func (m *mockRepo) Delete(ctx context.Context, ids []int) (int, error) {
	return len(ids), nil
}

func (m *mockRepo) Get(ctx context.Context, ids []int, filter task.ListFilter) ([]task.Task, error) {
	now := time.Now()
	return []task.Task{
		{ID: 1, Description: "Task 1", CompletedAt: nil},
		{ID: 2, Description: "Task 2", CompletedAt: &now},
	}, nil
}

func (m *mockRepo) Complete(ctx context.Context, ids []int) (int, error) {
	return len(ids), nil
}

// ------------------------
// Error repository (for testing errors)
// ------------------------

type errorRepo struct{}

func (m *errorRepo) Create(ctx context.Context, tasks []task.Task) error {
	return errMock("create failed")
}

func (m *errorRepo) Delete(ctx context.Context, ids []int) (int, error) {
	return 0, errMock("delete failed")
}

func (m *errorRepo) Get(ctx context.Context, ids []int, filter task.ListFilter) ([]task.Task, error) {
	return nil, errMock("list failed")
}

func (m *errorRepo) Complete(ctx context.Context, ids []int) (int, error) {
	return 0, errMock("complete failed")
}

// simple helper for error
type errMock string

func (e errMock) Error() string { return string(e) }

// ------------------------
// Tests
// ------------------------

func newTestCLI(repo task.Repository) (*CLI, *bytes.Buffer, *bytes.Buffer) {
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	svc := task.NewService(repo)
	c := New(svc, out, errOut)
	return c, out, errOut
}

func TestCLI_PrintUsage(t *testing.T) {
	c, out, _ := newTestCLI(&mockRepo{})
	c.printUsage()

	got := out.String()
	if !strings.Contains(got, "Usage") {
		t.Errorf("expected usage printed, got %q", got)
	}
}

func TestCLI_NewCommand(t *testing.T) {
	c, out, _ := newTestCLI(&mockRepo{})

	args := []string{"cli", "new", "My Task"}
	c.Run(context.Background(), args)

	got := out.String()
	if !strings.Contains(got, "Task") {
		t.Errorf("expected task printed, got %q", got)
	}
}

func TestCLI_RemoveCommand(t *testing.T) {
	c, out, _ := newTestCLI(&mockRepo{})

	args := []string{"cli", "remove", "1", "2"}
	c.Run(context.Background(), args)

	got := out.String()
	if !strings.Contains(got, "2 of 2 tasks successfully deleted") {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestCLI_ListCommand(t *testing.T) {
	c, out, _ := newTestCLI(&mockRepo{})

	args := []string{"cli", "list"}
	c.Run(context.Background(), args)

	got := out.String()
	if !strings.Contains(got, "[ ] 1 - Task 1") || !strings.Contains(got, "[x] 2 - Task 2") {
		t.Errorf("unexpected output:\n%s", got)
	}
}

func TestCLI_CompleteCommand(t *testing.T) {
	c, out, _ := newTestCLI(&mockRepo{})

	args := []string{"cli", "complete", "1", "2"}
	c.Run(context.Background(), args)

	got := out.String()
	if !strings.Contains(got, "2 of 2 tasks successfully completed") {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestCLI_InvalidIDs(t *testing.T) {
	c, _, errOut := newTestCLI(&mockRepo{})

	args := []string{"cli", "remove", "notanumber"}
	c.Run(context.Background(), args)

	got := errOut.String()
	if !strings.Contains(got, "failed to parse ID") {
		t.Errorf("expected parse error, got %q", got)
	}
}

func TestCLI_NewCommandError(t *testing.T) {
	c, _, errOut := newTestCLI(&errorRepo{})

	args := []string{"cli", "new", "My Task"}
	c.Run(context.Background(), args)

	got := errOut.String()
	if !strings.Contains(got, "create failed") {
		t.Errorf("expected error printed, got %q", got)
	}
}

func TestCLI_RemoveCommandError(t *testing.T) {
	c, _, errOut := newTestCLI(&errorRepo{})

	args := []string{"cli", "remove", "1"}
	c.Run(context.Background(), args)

	got := errOut.String()
	if !strings.Contains(got, "delete failed") {
		t.Errorf("expected error printed, got %q", got)
	}
}

func TestCLI_ListCommandError(t *testing.T) {
	c, _, errOut := newTestCLI(&errorRepo{})

	args := []string{"cli", "list"}
	c.Run(context.Background(), args)

	got := errOut.String()
	if !strings.Contains(got, "list failed") {
		t.Errorf("expected error printed, got %q", got)
	}
}

func TestCLI_CompleteCommandError(t *testing.T) {
	c, _, errOut := newTestCLI(&errorRepo{})

	args := []string{"cli", "complete", "1"}
	c.Run(context.Background(), args)

	got := errOut.String()
	if !strings.Contains(got, "complete failed") {
		t.Errorf("expected error printed, got %q", got)
	}
}
