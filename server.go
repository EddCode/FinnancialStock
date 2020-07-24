package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Listen() {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleRoot)
	router.HandleFunc("/hello", HandleHello)

	log.Fatal(http.ListenAndServe(s.port, router))
}
