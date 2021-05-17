package httperror

type HTTPError struct {
	Error bool `json:"error"` 		// Error is true if an error occurred, false if there is no error
	Message string `json:"message,omitempty"` // Message contains the error message
	HTTPStatusCode int `json:"-"`	// HTTPStatusCode contains the HTTP status response code
}

func ConvertErrorToHTTPError(err error, statusCode ...int) HTTPError {
	var status int
	if len(statusCode) > 0 {
		status = statusCode[0]
	} else {
		status = 500 // Unknown error, make it a Internal Server Error (500)
	}

	return HTTPError{
		Error: true,
		Message: err.Error(),
		HTTPStatusCode: status,
	}
}