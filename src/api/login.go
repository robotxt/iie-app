package api

import (
	"encoding/json"
	"net/http"

	"github/robotxt/iie-app/src/service"
)

type loginResponse struct {
	Token string
}

type loginData struct {
	Email    string
	Password string
}

// LoginApi
func (api *ApiV1) LoginApi(w http.ResponseWriter, r *http.Request) {
	rd := &loginData{}
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

	user, err := thisUser.GetUserByEmail(api.Ctx)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, &registrationResponse{
			Message: "User does not exist.",
		})
		return
	}

	thisUser.UID = user.UID
	token, err := thisUser.CreateCustomToken(api.Ctx)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, &registrationResponse{
			Message: "Error: Please try again later",
		})
		return
	}

	respondJSON(w, http.StatusOK, &loginResponse{
		Token: token,
	})
}
