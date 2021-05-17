package httpHandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"scolix/pkg/httperror"
	"scolix/pkg/httpresponses"
)

func Index(w http.ResponseWriter, r *http.Request) {
	res := httpresponses.Response{
		Status: http.StatusOK,
		Error: httperror.HTTPError{Error: false},
		Body:  httpresponses.IndexReponse{Message: "Hello World"},
	}

	w.WriteHeader(res.Status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		res = httpresponses.Response{
			Status: http.StatusInternalServerError,
			Error: httperror.ConvertErrorToHTTPError(err),
		}
		_ = json.NewEncoder(w).Encode(res)
	}
}