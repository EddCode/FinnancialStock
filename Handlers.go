package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello there new request!")
	io.WriteString(w, "Hello world from root")
}
