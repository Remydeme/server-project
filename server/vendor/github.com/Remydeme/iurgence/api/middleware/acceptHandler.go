package middleware

import (
	"encoding/json"
	err "github.com/Remydeme/iurgence/api/errors"
	"net/http"
)

func AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// We send a JSON-API error if the Accept header does not have a valid value.
		if r.Header.Get("Accept") != "application/json" {
			jsonErr := &err.Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/json'"}
			w.Header().Set("Content-Type", "application/vnd.api+json")
			w.WriteHeader(jsonErr.Status)
			json.NewEncoder(w).Encode(err.Errors{[]*err.Error{jsonErr}})
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
