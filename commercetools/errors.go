package commercetools

import "fmt"

type RequestError interface {
	// The status code of the HTTP response.
	StatusCode() int

	Message() string

	Error() string
}

type requestError struct {
	message    string
	statusCode int
}

func (e requestError) Message() string {
	return e.message
}

func (e requestError) StatusCode() int {
	return e.statusCode
}

func (e requestError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode(), e.Message())
}

type ApplicationError struct {
	Message string
	Errors  []struct {
		Code    string
		Message string
	}
}

func (e ApplicationError) Error() string {
	return e.Message
}

func newRequestError(msg string, statusCode int) RequestError {
	return requestError{
		message:    msg,
		statusCode: statusCode,
	}
}
