package main

// APIError is a custom error type for the APIServer
type APIError struct {
	ErrorMessage string `json:"error"`
	Code         int    `json:"code"`
}

// Error returns a string of the ErrorMessage
func (e *APIError) Error() string {
	return e.ErrorMessage
}
