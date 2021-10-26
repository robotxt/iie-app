package api

import (
	"encoding/json"
	"net/http"
)

type ItemData struct {
	Email    string
	Password string
}

// CreateItemApi
func (api *ApiV1) CreateItemApi(w http.ResponseWriter, r *http.Request) {
	rd := &ItemData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	respondJSON(w, http.StatusOK, &registrationResponse{
		Message: "Hello World!",
	})
}
