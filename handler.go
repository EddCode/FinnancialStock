package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello there new request from root!")
	io.WriteString(w, "Hello world from root")
}

func HandleHello(w http.ResponseWriter, request *http.Request) {
	user := &User{
		Name:  "rafa",
		Email: "rafa@bebe.com",
	}
	fmt.Println("Hello there new request!")
	json.NewEncoder(w).Encode(user)
	io.WriteString(w, "")
}

func HandleShowPokemon(w http.ResponseWriter, request *http.Request) {
	financialapi.RequestApi("pokemon/ditto")
	io.WriteString(w, "Hello pokemon")
}
