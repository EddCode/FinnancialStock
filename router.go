package main

import "net/http"

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.findHandler(request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}

// Check if exists a rigth handler by request path
func (r *Router) findHandler(path string) (http.HandlerFunc, bool) {
	handler, exists := r.routes[path]
	return handler, exists
}
