package cterrors

import "fmt"

// A requestError wraps a request or service error.
//
// Composed of baseError for code, message, and original error.
type requestError struct {
	message    string
	statusCode int
}

// Code returns the short phrase depicting the classification of the error.
func (r requestError) StatusCode() int {
	return r.statusCode
}

// Message returns the error details message.
func (r requestError) Message() string {
	return r.message
}

// Error returns the string representation of the error.
// Satisfies the error interface.
func (r requestError) Error() string {
	return fmt.Sprintf("%d: %s", r.StatusCode(), r.Message())
}

// newRequestError returns a wrapped error with additional information for
// request status code, and service requestID.
//
// Should be used to wrap all request which involve service requests. Even if
// the request failed without a service response, but had an HTTP status code
// that may be meaningful.
//
// Also wraps original errors via the baseError.

func newRequestError(msg string, statusCode int) *requestError {
	return &requestError{
		message:    msg,
		statusCode: statusCode,
	}
}
