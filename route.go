package main

import (
	"github.com/eddcode/fintechApp/handler"
	"github.com/eddcode/fintechApp/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Add Middleware
	router.Use(middleware.MorganGo)

	router.HandleFunc("/", handler.Root)
	router.HandleFunc("/login", handler.Login).
		Headers("Content-type", "application/json").
		Methods("POST")
	router.HandleFunc("/singup", handler.Signup).
		Headers("Content-type", "application/json").
		Methods("POST")

	return router
}
