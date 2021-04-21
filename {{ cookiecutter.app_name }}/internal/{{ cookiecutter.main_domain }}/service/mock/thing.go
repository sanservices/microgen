package mock

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// GetThing is mock
func (m Mock) GetThing(ctx context.Context, id uint) (*entity.Thing, error) {
	args := m.Mock.Called(ctx, id)
	thing := args.Get(0).(*entity.Thing)
	return thing, nil
}
