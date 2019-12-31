package middleware

import (
	"encoding/json"
	"github.com/Remydeme/iurgence/api/errors"
	"log"
	"net/http"
)

/**
Middle that handle panic and return error 500 message
*/
func RecoverHandler(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic : %+v", err)
				jsonError := &errors.Error{errors.InternalServerErrorId, http.StatusInternalServerError, http.StatusText(500), "something went wrong"}
				w.Header().Set("Content-Type", "application/vnd.api+json")
				w.WriteHeader(jsonError.Status)
				json.NewEncoder(w).Encode(errors.Errors{[]*errors.Error{jsonError}})
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
