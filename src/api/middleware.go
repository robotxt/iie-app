package api

import (
	"context"
	"encoding/json"
	"github/robotxt/iie-app/src/logging"
	repo "github/robotxt/iie-app/src/repo/firebase"
	"github/robotxt/iie-app/src/service"
	"net/http"
	"os"
	"reflect"
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

			var apiKey = os.Getenv("BASIC_API_KEY")
			var authHeader = os.Getenv("AUTH_HEADER")

			var authToken = r.Header.Get(authHeader)
			json.NewEncoder(w).Encode(r)
			token := strings.TrimSpace(authToken)

			if token == "" {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode("Missing Authentication Header.")
				return
			}

			verified := false
			if authToken == apiKey {
				// Public API will used public API KEY's
				publicUrlsArray := reflect.ValueOf(PublicURLS)
				for i := 0; i < publicUrlsArray.Len(); i++ {
					if publicUrlsArray.Index(i).Interface() == r.URL.Path {
						verified = true
					}
				}
			}

			if verified {
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
