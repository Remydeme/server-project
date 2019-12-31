package logger

import (
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		fmt.Printf("[%s] - %s %s \n", r.Method, r.URL.String(), end.Sub(start))
	}
	return http.HandlerFunc(f)
}
