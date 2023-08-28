package {{ cookiecutter.entity_name }}

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
}
