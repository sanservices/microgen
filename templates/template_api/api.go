package api_version

import (
	"github.com/go-playground/validator"
	"github.com/san-services/apicore"
	"github.com/san-services/apicore/apisettings"
)

// Handler is the api handler for version api_version.
type Handler struct {
	conf *apisettings.Settings
	// service  mydomain.Service
	validate *validator.Validate
}

func NewHandler(
	settings *apisettings.Settings,
	// service mydomain.Service // -- Add this with the correct service reference
) *Handler {

	validate := apicore.ValidationRules()

	return &Handler{
		conf: settings,
		// service:  service,
		validate: validate,
	}
}
