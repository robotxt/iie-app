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

// PublicURLS using default key
var PublicURLS = []string{
	"/api/v1/hello",
	"/api/v1/login",
	"/api/v1/registration",
}

func (a *ApiV1) apiv1Handler(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	r := a.Router.PathPrefix("/api/v1").Subrouter()
	r.HandleFunc(path, f).Methods(method)
}

func (a *ApiV1) SetRouters() {

	// API secure authentication middleware
	middleware := Middleware{Ctx: a.Ctx, Firebase: a.Firebase}
	a.Router.Use(middleware.SecureApiRequest())

	// API endpoints
	a.apiv1Handler("GET", "/hello", HandleRequest(a.BaseAPI))
	a.apiv1Handler("POST", "/registration", HandleRequest(a.RegistrationApi))
	a.apiv1Handler("POST", "/login", HandleRequest(a.LoginApi))

	// secure API - Not public using JWT auth
	a.apiv1Handler("POST", "/item", HandleRequest(a.CreateItemApi))
	a.apiv1Handler("GET", "/fetch/items", HandleRequest(a.FetchItems))
}
