package {{ cookiecutter.main_domain }}

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// Service declares and summarizes the functionality a
// service in the containing package will implement
type Service interface {
	GetThing(ctx context.Context, id uint) (*entity.Thing, error)
}

// Repository declares and summarizes the functionality a
// repository in the containing package will implement
type Repository interface {
	GetThing(ctx context.Context, id uint) (*entity.ThingRec, error)
}
