package service

import (
	"context"
	{% if cookiecutter.use_cache != 'n' %}"fmt"{% endif %}
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.module_name }}-proto/pb"
)

// GetUser gives you a thing from storage
func (s *Service) GetUser(ctx context.Context, id uint32) (*pb.GetUserResponse, error) {
	//TODO: Implement business logic
	{% if cookiecutter.use_database == 'y' %}

	user, err := s.repo.GetUser(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		UserID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Age: user.Age,
		Email:user.Email,
	}, nil

	{% else %}
	return u, nil
	{% endif %}
}
