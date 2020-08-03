package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eddcode/fintechApp/model"
)

func Root(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "Hello world from root")
}

func ErrorHandler(response http.ResponseWriter, req *http.Request, errMessage model.ErrorResponse) {
	httpResponse := &model.ErrorResponse{Code: errMessage.Code, Message: errMessage.Message}
	jsonResponse, err := json.Marshal(httpResponse)

	if err != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errMessage.Code)
	response.Write(jsonResponse)
	req.Body.Close()
}
