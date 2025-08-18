package common

// ErrorResponse represents a standard API error response.
type ErrorResponse struct {
	// Error description
	Message string `json:"message"`
}

// Error implements the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}
