package httpresponses

import "scolix/pkg/httperror"

type Response struct {
	Error 	httperror.HTTPError 	`json:"error,omitempty"`
	Status 	int 					`json:"-"`
	Body 	interface{} 			`json:"body,omitempty"`
}