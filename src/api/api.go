package api

import (
	"context"
	"github/robotxt/iie-app/src/middleware"
	"net/http"

	repo "github/robotxt/iie-app/src/repo/firebase"

	"github.com/gorilla/mux"
)

type ApiV1 struct {
	Router   *mux.Router
	Ctx      context.Context
	Firebase *repo.FirebaseApp
}

func (a *ApiV1) apiv1Handler(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	secure := a.Router.PathPrefix("/api/v1").Subrouter()
	secure.HandleFunc(path, f).Methods(method)
}

func (a *ApiV1) SetRouters() {
	a.apiv1Handler("GET", "/", middleware.HandleRequest(a.BaseAPI))
	a.apiv1Handler("POST", "/registration", middleware.HandleRequest(a.RegistrationApi))

}
