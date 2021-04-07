package errs

// ServiceError is common error struct for api errors
type ServiceError struct {
	Message    string          `json:"message"`
	Code       string          `json:"code"`
	Properties []PropertyError `json:"properties,omitempty"`
}

// PropertyError is error of property
type PropertyError struct {
	Property string   `json:"property"`
	Messages []string `json:"message"`
}

// New ServiceError
func New(msg, code string) ServiceError {
	return ServiceError{
		Message: msg,
		Code:    code,
	}
}

func (e ServiceError) Error() string {
	return e.Message
}
