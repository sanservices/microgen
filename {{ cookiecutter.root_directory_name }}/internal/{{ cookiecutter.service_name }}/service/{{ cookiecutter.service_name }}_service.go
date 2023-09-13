package service

import (
	"context"
	{% if cookiecutter.use_cache != 'n' %}"fmt"{% endif %}

	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/entity"
)

// GetThing gives you a thing from storage
func (s *Service) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	{% if cookiecutter.use_cache != 'n' %}
	var u *entity.User
	s.redis.Get(ctx, fmt.Sprintf("id::%d", id), u)
	{% endif %}

	//TODO: Implement business logic
	{% if cookiecutter.use_database != 'n' %}
	return s.repo.GetUser(ctx, id)
	{% else %}
	return u, nil
	{% endif %}
}

func (s *Service) CreateUser(ctx context.Context, thing *entity.User) error {
	{% if cookiecutter.use_cache != 'n' %}
	var u *entity.User
	defer s.redis.Set(ctx, thing.ID, u)
	{% endif %}

	//TODO: Implement business logic
	{% if cookiecutter.use_database != 'n' %}
	return s.repo.SaveUser(ctx, thing)
	{% else %}
	return nil
	{% endif %}
}
