package api

import (
	"net/http"
)

type baseResponse struct {
	Message string
}

// BaseAPI
func (api *ApiV1) BaseAPI(w http.ResponseWriter, r *http.Request) {

	respondJSON(w, http.StatusOK, &baseResponse{
		Message: "Hello World!",
	})
}
