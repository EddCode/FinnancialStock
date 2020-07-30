package handler

import (
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "The Login")
}
