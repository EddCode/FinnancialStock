package main

import (
	"net/http"

	"github.com/eddcode/fintechApp/handler"
	"github.com/eddcode/fintechApp/middleware"
	"github.com/gorilla/mux"
)

// Call the next handler, which can be another middleware in the chain, or the final handler.
func AddMiddlewares(handler http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler

}

func Router() *mux.Router {
	router := mux.NewRouter()

	// Add Middleware
	router.Use(middleware.MorganGo)

	router.HandleFunc("/", handler.Root)
	router.HandleFunc("/login", AddMiddlewares(handler.Login, middleware.ValidateLogin())).
		Headers("Content-type", "application/json").
		Methods("POST")
	router.HandleFunc("/singup", handler.Signup).
		Headers("Content-type", "application/json").
		Methods("POST")

	return router
}
