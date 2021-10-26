package api

import (
	"context"
	"encoding/json"
	"github/robotxt/iie-app/src/logging"
	repo "github/robotxt/iie-app/src/repo/firebase"
	"github/robotxt/iie-app/src/service"
	"net/http"
	"os"
	"strings"
)

// UserCtxKey for context user key
type UserCtxKey string

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

type Middleware struct {
	Firebase *repo.FirebaseApp
	Ctx      context.Context
}

func HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func (m *Middleware) SecureApiRequest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var app_api_key = os.Getenv("BASIC_API_KEY")
			var header = r.Header.Get("HTTP_AUTHORIZATION")

			json.NewEncoder(w).Encode(r)
			token := strings.TrimSpace(header)

			logging.Info("secure handler middleware activated: ", token)

			if token == "" {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode("Missing HTTP_AUTHORIZATION Header")
				return
			}

			// TODO check for public API
			if header == app_api_key {
				next.ServeHTTP(w, r)
				return
			}

			decodedToken, err := m.Firebase.VerifyCustomToken(token)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode("Forbidden")
				return
			}

			authUser := service.UserType{}
			authUser.UID = decodedToken.UID
			fbUser, _ := authUser.GetUserByUID(m.Ctx)
			authUser.Email = fbUser.Email

			logging.Info("loggedin user: ", authUser)

			newCtx := context.WithValue(m.Ctx, UserCtxKey("authUser"), authUser)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
