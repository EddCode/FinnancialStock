package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eddcode/fintechApp/handler"
	"github.com/eddcode/fintechApp/model"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MorganGo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(r.Method, r.RequestURI, time.Since(start))
		}()

		next.ServeHTTP(w, r)
	})
}

func ValidateLogin() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var loginUser model.SingIn

			errorResponse := model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "We are sorry, Some happend in the server",
			}

			data := json.NewDecoder(r.Body)

			if dataErr := data.Decode(&loginUser); dataErr != nil {
				handler.ErrorHandler(w, r, errorResponse)
			}

			if loginUser.Email == "" || loginUser.Password == "" {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "Email and Password are required"
				handler.ErrorHandler(w, r, errorResponse)
			}

			ctx := context.WithValue(r.Context(), "user", loginUser)
			next(w, r.WithContext(ctx))
		}
	}
}

func ValidateRegister() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var register model.SingUp
			errorResponse := model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "We are sorry, Some happend in the server",
			}

			data := json.NewDecoder(r.Body)

			if err := data.Decode(&register); err != nil {
				handler.ErrorHandler(w, r, errorResponse)
			}

			if register.Email == "" || register.Name == "" || register.Password == "" {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "All fields are required"
				handler.ErrorHandler(w, r, errorResponse)
			}

			ctx := context.WithValue(r.Context(), "register", register)
			next(w, r.WithContext(ctx))
		}
	}
}
