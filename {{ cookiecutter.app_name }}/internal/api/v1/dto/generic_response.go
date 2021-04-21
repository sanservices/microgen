package dto

// Generic messages

// Generic success
// swagger:model genericSuccessRS
type genericSuccessResp struct {
	// in: body
	Data interface{}
}

// Generic error
// swagger:model genericErrorRS
type genericErrorResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Bad request
// swagger:model badRequestRS
type badRequestResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Internal server error
// swagger:model serverErrorRS
type serverErrorResp struct {
	// Error message
	// in: body
	Body string `json:"body"`
}
