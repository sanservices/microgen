package mocked

import (
	"context"
	"goproposal/internal/models"

	"github.com/san-services/apicore/apisettings"
)

// InMemory struct
type Mocked struct {
	data *store
}

type store struct {
	persons []models.Person
}

// New returns a new InMemory struct
func New(ctx context.Context, db apisettings.Database) (*Mocked, error) {
	new := &Mocked{}
	err := new.Connect(ctx)
	return new, err
}

// Connect inmemory implementation
func (m *Mocked) Connect(ctx context.Context) error {
	m.data = &store{}
	return nil
}

// Close inmemory implementation
func (m Mocked) Close(ctx context.Context) error {
	return nil
}
