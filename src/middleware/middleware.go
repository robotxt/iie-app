package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func SecureHandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("HTTP_AUTHORIZATION")

		json.NewEncoder(w).Encode(r)
		token := strings.TrimSpace(header)

		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing HTTP_AUTHORIZATION Header")
			return
		}

		// TODO: Check token header is valid

		handler(w, r)
	}
}
