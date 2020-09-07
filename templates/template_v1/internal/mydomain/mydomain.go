package mydomain

import (
	"context"
	"goproposal/internal/models"
)

// Service declares and summerizes the functionality a
// service in the containing package will implement
type Service interface {
	AddPerson(name string, age int32) error
	Person(id int64) (*models.Person, error)
	Persons() ([]models.Person, error)
}

// Repository declares and summerizes the functionality a
// repository in the containing package will implement
type Repository interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	GetPerson(id int64) (*models.Person, error)
	GetAllPersons() ([]models.Person, error)
	SavePerson(name string, age int32) error
}
