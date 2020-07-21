package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) HandleRoute(path string, handler http.HandlerFunc) {
	s.router.routes[path] = handler
}

func (s *Server) Listen() {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println("Error occure to up server ")
		panic(err)
	}

	fmt.Printf("Server is runing on port %s", s.port)
}
