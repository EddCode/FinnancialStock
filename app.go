package main

import (
	"fmt"
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
	fmt.Println("Starting Server on localhost:8080")
	log.Println(http.ListenAndServe(a.addr, route))
}
