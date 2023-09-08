package oracle

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/entity"
)

// GetThing stands for getting a thing from db
func (m *Oracle) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	thing := &entity.User{}

	//TODO: Implement mysql calls
	return thing, nil
}

func (m *Oracle) SaveUser(ctx context.Context, thing *entity.User) error {

	//TODO: Implement mysql calls
	return nil
}
