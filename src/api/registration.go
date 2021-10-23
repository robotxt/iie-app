package api

import (
	"encoding/json"
	"net/http"

	"github/robotxt/iie-app/src/logging"
	"github/robotxt/iie-app/src/service"
)

type registrationResponse struct {
	Message string
}

type registrationData struct {
	Email    string
	Password string
}

// RegistrationApi
func (api *ApiV1) RegistrationApi(w http.ResponseWriter, r *http.Request) {
	rd := &registrationData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// var userServices service.UserService
	thisUser := service.UserType{}
	thisUser.Email = rd.Email
	thisUser.Password = rd.Password

	existingUser, err := thisUser.GetUserByEmail(api.Ctx)
	if existingUser != nil {
		respondJSON(w, http.StatusBadRequest, &registrationResponse{
			Message: "User already exist.",
		})
		return
	}

	// update to new hashed password
	newPassword := thisUser.HashPassword()
	thisUser.Password = string(newPassword)

	newUser, err := thisUser.CreateFirebaseUser(api.Ctx)

	thisUser.UID = newUser.UID
	thisUser.CreateUserProfile(api.Ctx)

	if err != nil {
		respondJSON(w, http.StatusBadRequest, &registrationResponse{
			Message: "User already exist.",
		})
		return
	}

	logging.Info("new user is created: ", newUser)

	respondJSON(w, http.StatusOK, &registrationResponse{
		Message: "Hello World!",
	})
}
