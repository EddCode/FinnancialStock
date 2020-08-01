package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eddcode/fintechApp/model"
	util "github.com/eddcode/fintechApp/utils"
)

func Login(w http.ResponseWriter, request *http.Request) {
	response := model.Claim{
		Email: "12312",
		User:  "hoy en 10 minutos",
	}

	json.NewEncoder(w).Encode(response)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var singupUser model.SingUp
	data := json.NewDecoder(r.Body)
	dataErr := data.Decode(&singupUser)

	if dataErr != nil {
		panic(dataErr)
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
