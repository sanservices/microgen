package v1

import (
	"goproposal/internal/mydomain"

	"github.com/go-playground/validator"
	"github.com/san-services/apicore"
	"github.com/san-services/apicore/apisettings"
)

// Handler is the api handler for version 1.
type Handler struct {
	conf     *apisettings.Settings
	service  mydomain.Service
	validate *validator.Validate
}

func NewHandler(settings *apisettings.Settings, service mydomain.Service) *Handler {
	validate := apicore.ValidationRules()

	return &Handler{
		conf:     settings,
		service:  service,
		validate: validate,
	}
}
