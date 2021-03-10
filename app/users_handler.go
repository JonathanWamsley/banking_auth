package app

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanwamsley/banking_auth/dto"
	"github.com/jonathanwamsley/banking_auth/service"
)

type UserHandler struct {
	service service.UsersService
}

func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		writeResponse(w, http.StatusBadRequest, "invalid json")
		return
	}
	user, err := uh.service.CreateUser(userRequest)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, user)
}

func (uh UserHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminRequest dto.AdminRequest
	if err := json.NewDecoder(r.Body).Decode(&adminRequest); err != nil {
		writeResponse(w, http.StatusBadRequest, "invalid json")
		return
	}
	admin, err := uh.service.CreateAdmin(adminRequest)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, admin)
}

func (uh UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetUsers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, users)
}
