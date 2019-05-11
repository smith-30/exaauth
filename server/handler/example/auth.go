package example

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/smith-30/exaauth/server/request"
	"github.com/smith-30/exaauth/server/response"
)

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	a := &request.Auth{}
	err := json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		response.Json(w, http.StatusBadRequest, nil)
		return
	}
	// c := jwt.MapClaims{"user_id": 123}
	_, tokenString, _ := h.JWTAuth.Encode(jwt.MapClaims{"user_id": 123})

	response.Json(w, http.StatusOK, response.Auth{Token: tokenString})
}
