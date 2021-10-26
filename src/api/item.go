package api

import (
	"encoding/json"
	"github/robotxt/iie-app/src/logging"
	"github/robotxt/iie-app/src/service"
	"net/http"
)

type ItemData struct {
	Name        string
	Description string
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

	authUser := r.Context().Value(UserCtxKey("authUser")).(service.UserType)
	logging.Info("authuser UID: ", authUser.UID)
	logging.Info("authuser Email: ", authUser.Email)

	respondJSON(w, http.StatusOK, &registrationResponse{
		Message: "Hello World!",
	})
}
