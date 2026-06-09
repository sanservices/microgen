package dto

// Generic messages

// Generic success
type genericSuccessResp struct {
	Data interface{}
}

// Generic error
type genericErrorResp struct {
	Body string `json:"error"`
}

// Bad request
type badRequestResp struct {
	Body string `json:"error"`
}

// Internal server error
type serverErrorResp struct {
	Body string `json:"body"`
}
