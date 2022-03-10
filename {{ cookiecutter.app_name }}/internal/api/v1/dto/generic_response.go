package dto

// Generic messages

// Generic success
// swagger:response genericSuccessRS
type genericSuccessResp struct {
	// in: body
	Data interface{}
}

// Generic error
// swagger:response genericErrorRS
type genericErrorResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Bad request
// swagger:response badRequestRS
type badRequestResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Internal server error
// swagger:response serverErrorRS
type serverErrorResp struct {
	// Error message
	// in: body
	Body string `json:"body"`
}
