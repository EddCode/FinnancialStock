package handler

import (
	"io"
	"net/http"
)

func Root(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "Hello world from root")
}
