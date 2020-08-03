package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eddcode/fintechApp/model"
	util "github.com/eddcode/fintechApp/utils"
)

func Login(w http.ResponseWriter, request *http.Request) {

	user := request.Context().Value("user").(model.SingIn)

	token, err := util.SingToken(user.Email)

	if err != nil || token == "" {
		panic(err)
	}

	loginToken := model.SuccessLogin{
		Email: user.Email,
		Token: token,
	}

	responseOk := model.SuccessResponse{
		Code:     http.StatusOK,
		Message:  "Your logged in",
		Response: loginToken,
	}

	jsonResponse, jsonErr := json.Marshal(responseOk)
	if jsonErr != nil {
		panic(jsonErr)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(jsonResponse))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var singupUser model.SingUp
	errorResponse := model.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "We are sorry, Some happend in the server",
	}

	data := json.NewDecoder(r.Body)
	dataErr := data.Decode(&singupUser)

	if dataErr != nil {
		panic(errorResponse)
	}

	token, err := util.SingToken(singupUser.Email)

	if err != nil || token == "" {
		panic(err)
	}

	singToken := model.SuccessLogin{
		Email: singupUser.Email,
		Token: token,
	}

	responseOk := model.SuccessResponse{
		Code:     http.StatusOK,
		Message:  "You`re registred correctly",
		Response: singToken,
	}

	jsonResponse, jsonErr := json.Marshal(responseOk)

	if jsonErr != nil {
		panic(jsonErr)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, string(jsonResponse))
}
