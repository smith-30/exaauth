package example

import "github.com/go-chi/jwtauth"

type AuthHandler struct {
	JWTAuth *jwtauth.JWTAuth
}

func New(a *jwtauth.JWTAuth) *AuthHandler {
	return &AuthHandler{
		JWTAuth: a,
	}
}
