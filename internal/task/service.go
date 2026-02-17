package task

import "context"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Create(ctx context.Context, desc []string) error {
	tasks := []Task{}
	for _, d := range desc {
		t := Task{
			Description: d,
		}
		if err := t.validate(); err != nil {
			return err
		}
		tasks = append(tasks, t)
	}

	if err := s.r.Create(ctx, tasks); err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, ids []int) error {
	return nil
}

func (s *Service) List(ctx context.Context, ids []int, filter, order *string) error {
	return nil
}

func (s *Service) Complete(ctx context.Context, ids []int) error {
	return nil
}
