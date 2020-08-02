package model

import jwt "github.com/dgrijalva/jwt-go"

type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

type Claim struct {
	Email string
	User  string
	jwt.StandardClaims
}

// User is struct use for user data
type User struct {
	Name     string
	Email    string
	Password string
}

// SiSingUp is struct to read the request body
type SingUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SinSingInis struct to read the request body
type SingIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessLogin struct {
	Email string `json:"emil"`
	Token string `json:"token"`
}
