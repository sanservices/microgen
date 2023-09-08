package {{ cookiecutter.app_name }}

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/entity"
)

// Service is the contract for the service layer, 
// responsible for handling business logic and orchestrating actions.
// For example:
//
//	type Service interface {
//		Entity(ctx context.Context, e *entity.Entity) error
//  	RegisterEntity(ctx context.Context, e *entity.Entity) error
// 		UpateEntity(ctx context.Context, e *entity.Entity) error
//	}
type Service interface {
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	CreateUser(ctx context.Context, thing *entity.User) error
}

// Repository is the contract interface for the repository layer, 
// responsible for data storage and retrieval.
// For example:
//
//	type Repository interface {
//		Find(ctx context.Context, e *entity.Entity) error
//  	RegisterEntity(ctx context.Context, e *entity.Entity) error
// 		UpateEntity(ctx context.Context, e *entity.Entity) error
//	}
type Repository interface {
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	SaveUser(ctx context.Context, thing *entity.User) error
}

{% if cookiecutter.use_cache == 'y' %}
type Cache interface {
	Set(ctx context.Context, key string, v interface{}) error
	Get(ctx context.Context, key string, v interface{}) error
	Delete(ctx context.Context, key string) error
}
{% endif %}