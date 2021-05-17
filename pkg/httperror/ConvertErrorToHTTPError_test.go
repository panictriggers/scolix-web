package httperror

import (
	"errors"
	"testing"
)

func TestConvertErrorToHTTPError(t *testing.T) {
	err := errors.New("unknown error")
	e := ConvertErrorToHTTPError(err)
	expectedError := HTTPError{
		Error: true,
		Message: "unknown error",
		HTTPStatusCode: 500,
	}

	if e != expectedError {
		t.Failed()
		t.Errorf("Received output: \n%+v\n", e)
		t.Errorf("Expected output: \n%+v\n", expectedError)
	}

	err = errors.New("bad request")
	e = ConvertErrorToHTTPError(err, 400)
	expectedError = HTTPError{
		Error: true,
		Message: "bad request",
		HTTPStatusCode: 400,
	}

	if e != expectedError {
		t.Failed()
		t.Errorf("Received output: \n%+v\n", e)
		t.Errorf("Expected output: \n%+v\n", expectedError)
	}
}
