package dto

import "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/entity"

// Examples with validation
type GetUserRQ struct {
	Query struct {
		ID    *int    `query:"id" validate:"omitempty,numeric"`
		Email *string `query:"email" validate:"omitempty,email"`
		Name  *string `query:"firstname" validate:"omitempty"`
	} `query:"data"`
}

// Responses can be used like this or you can return the entity directly. Body can be used for POST and PUT requests as arguments
type GetUserRS struct {
	Body struct {
		Users []userExample `json:"users"`
	} `json:"data"`
}

// This struct should be an entity instead of an internal dto struct
type userExample struct {
	ID        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

// disclaimer: there are non-ok json tag values for some of the *RQ structs, because they are in path or somewhere else, so it is not actual json
// swagger:parameters getThingRQ
type getThingRQ struct {
	// in: query
	ID string `json:"id"`
}

// swagger:response getThingRS
type getThingRS struct {
	// in: body
	Thing *entity.User `json:"thing"`
}
