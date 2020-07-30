package main

import (
	"log"
	"net/http"
)

type App struct {
	addr string
}

func InitApp(addr string) *App {
	return &App{
		addr: addr,
	}
}

func (a *App) Run() {
	route := Router()
	log.Println("Starting Server")
	log.Println(http.ListenAndServe(a.addr, route))
}
