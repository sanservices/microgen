package service

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// GetThing gives you a thing from storage
func (s *Service) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {

	//TODO: Implement business logic

	return s.repo.GetThing(ctx, id)
}

func (s *Service) CreateThing(ctx context.Context, thing *entity.Thing) error {

	//TODO: Implement business logic
	return s.repo.SaveThing(ctx, thing)
}
