package mock

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/entity"
)

// Mock is just mock repo with mock instance from testify
type Mock struct {
	Mock *mock.Mock
}

// NewWithExpectations creates new mock with pre-setup expectations
func NewWithExpectations() *Mock {
	m := &Mock{&mock.Mock{}}

	ctx := context.TODO()
	errCtx := context.WithValue(ctx, "wanterror", true)
	dummyErr := errors.New("let's say error")

	var fakeThing *entity.ThingRec
	gofakeit.Struct(&fakeThing)

	m.Mock.On("GetThing", errCtx, 42).
		Return(nil, dummyErr)

	m.Mock.On("GetThing", mock.Anything, 42).
		Return(fakeThing, nil)

	return m
}
