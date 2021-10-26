package api

import (
	"context"
	"net/http"

	repo "github/robotxt/iie-app/src/repo/firebase"

	"github.com/gorilla/mux"
)

type ApiV1 struct {
	Router   *mux.Router
	Ctx      context.Context
	Firebase *repo.FirebaseApp
}

func (a *ApiV1) apiv1Handler(method string, path string, f func(w http.ResponseWriter, r *http.Request), public bool) {
	r := a.Router.PathPrefix("/api/v1").Subrouter()
	r.HandleFunc(path, f).Methods(method)

	// API secure authentication middleware
	middleware := Middleware{Ctx: a.Ctx, Firebase: a.Firebase}
	a.Router.Use(middleware.SecureApiRequest(public))
}

func (a *ApiV1) SetRouters() {

	a.apiv1Handler("GET", "/", HandleRequest(a.BaseAPI), true)
	a.apiv1Handler("POST", "/registration", HandleRequest(a.RegistrationApi), true)

	// secure API - Not public using JWT auth
	a.apiv1Handler("POST", "/item", HandleRequest(a.CreateItemApi), false)
}
