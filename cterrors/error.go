package cterrors

// A RequestError is an interface to extract request failure information from
// an Error such as the request ID of the failed request returned by a service.
// RequestErrors may not always have a requestID value if the request failed
// prior to reaching the service such as a connection error.
//
// Example:
//
//     output, err := svc.GetByID(id)
//     if err != nil {
//         if reqerr, ok := err.(RequestError); ok {
//             log.Println("Request failed", reqerr.StatusCode(), reqerr.Message())
//         } else {
//             log.Println("Error:", err.Error())
//         }
//     }
type RequestError interface {
	// The status code of the HTTP response.
	StatusCode() int

	Message() string

	Error() string
}

// NewRequestError returns a new request error wrapper for the given Error
// provided.
func NewRequestError(msg string, statusCode int) RequestError {
	return newRequestError(msg, statusCode)
}
