package dtos

type PersonResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// person parameters
//
// swagger:parameters person
type GetPersonInput struct {
	// in:path
	// required: true
	ID int64 `json:"id"`
}

// persons parameters
//
// swagger:parameters persons
type GetPersonsInput struct {
}

// addPerson parameters
//
// swagger:parameters addPerson
type AddPersonInput struct {

	// in:body
	// required:true
	Body struct {
		Name string `json:"name" validate:"required,min=4"`
		Age  int32  `json:"age" validate:"required,min=18"`
	}
}

// swagger:response personResponse
type PersonResponse struct {
	// in: body
	Data PersonResp `json:"data"`
}

// swagger:response personsResponse
type PersonsResponse struct {
	// in: body
	Data []struct{ PersonResp } `json:"data"`
}
