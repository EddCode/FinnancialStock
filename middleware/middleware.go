package middleware

import (
	"log"
	"net/http"
	"time"
)

func MorganGo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(r.Method, r.RequestURI, time.Since(start))
		}()
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
