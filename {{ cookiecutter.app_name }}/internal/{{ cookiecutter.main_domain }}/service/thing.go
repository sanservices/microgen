package service

import (
	"context"
	"fmt"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// GetThing gives you a thing from storage
func (s *Service) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	thing, err := s.repo.GetThing(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Thing{
		ID:         thing.ID,
		CategoryID: thing.CategoryID,
		Image:      fmt.Sprintf("https://imageserver.com/%d/%d", thing.ID, thing.CategoryID),
	}, nil
}
