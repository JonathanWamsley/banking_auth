package app

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanwamsley/banking_auth/dto"
	"github.com/jonathanwamsley/banking_auth/logger"
	"github.com/jonathanwamsley/banking_auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

// Login verifies a user credentials and returns a dummy success msg in place of jwt
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.service.Login(loginRequest)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
