package {{ cookiecutter.service_name }}

import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/entity"
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
	{% if cookiecutter.use_database == 'y' %}
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	SaveUser(ctx context.Context, thing *entity.User) error
	{% endif %}
}

{% if cookiecutter.use_cache == 'y' %}
// Cache is the contract interface for the caching layer,
// responsible for storing and retrieving data from a cache store.
// For example:
//
//	type Cache interface {
//	    Set(ctx context.Context, key string, v interface{}) error
//	    Get(ctx context.Context, key string, v interface{}) error
//	    Delete(ctx context.Context, key string) error
//	}
type Cache interface {
	Set(ctx context.Context, key string, v interface{}) error
	Get(ctx context.Context, key string, v interface{}) error
	Delete(ctx context.Context, key string) error
}
{% endif %}