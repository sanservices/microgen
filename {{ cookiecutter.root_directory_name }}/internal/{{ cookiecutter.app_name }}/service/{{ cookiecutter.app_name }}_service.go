package service

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/entity"
)

// GetThing gives you a thing from storage
func (s *Service) GetUser(ctx context.Context, id uint) (*entity.User, error) {

	//TODO: Implement business logic

	return s.repo.GetUser(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, thing *entity.User) error {

	//TODO: Implement business logic
	return s.repo.SaveUser(ctx, thing)
}
