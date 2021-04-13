package mock

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"

	"github.com/stretchr/testify/mock"
)

// Mock represents mock service
type Mock struct {
	Mock *mock.Mock
}

// NewWithExpectations creates new mock with pre-setup expectations
func NewWithExpectations() *Mock {
	m := &Mock{&mock.Mock{}}

	ctx := context.TODO()
	errCtx := context.WithValue(ctx, "wanterror", true)
	dummyErr := errors.New("let's say error")

	var fakeThing *entity.Thing
	gofakeit.Struct(&fakeThing)

	m.Mock.On("GetThing", errCtx, 42).
		Return(nil, dummyErr)

	m.Mock.On("GetThing", mock.Anything, 42).
		Return(fakeThing, nil)

	return m
}
