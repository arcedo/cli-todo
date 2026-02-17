package task

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Create(ctx context.Context, desc []string) (tasks []Task, err error) {
	var errs []string
	for _, d := range desc {
		t := Task{Description: d}
		if err := t.validate(); err != nil {
			errs = append(errs, fmt.Sprintf("task '%s': %v", t.Description, err))
			continue
		}
		tasks = append(tasks, t)
	}
	if len(errs) > 0 {
		return tasks, fmt.Errorf("validation errors: %s", strings.Join(errs, "; "))
	}

	if err := s.r.Create(ctx, tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Service) Delete(ctx context.Context, ids []int) (affected int, err error) {
	affected, err = s.r.Delete(ctx, ids)
	if err != nil {
		return 0, fmt.Errorf("failed to delete tasks: %w", err)
	}
	return affected, nil
}

func (s *Service) List(ctx context.Context, ids []int, filter ListFilter) (tasks []Task, err error) {
	tasks, err = s.r.Get(ctx, ids, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}

func (s *Service) Complete(ctx context.Context, ids []int) (affected int, err error) {
	affected, err = s.r.Complete(ctx, ids)
	if err != nil {
		return 0, fmt.Errorf("failed to complete tasks: %w", err)
	}
	return affected, nil
}
