package fixture

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// GetThing stands for getting the thing by id
func (f *Fixture) GetThing(ctx context.Context, _ uint) (*entity.ThingRec, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var thing *entity.ThingRec
		gofakeit.Struct(&thing)
		return thing, nil
	}
}
