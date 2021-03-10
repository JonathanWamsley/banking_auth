package dto

import "github.com/jonathanwamsley/banking_auth/errs"

type UserRequest struct {
	Username   string `json"username`
	Password   string `json:"password"`
	CustomerID string `json:"customer_id"`
}

type UserResponse struct {
	Username   string `json"username`
	CustomerID string `json:"customer_id"`
	CreatedOn  string `json:"created_on"`
}

type AdminRequest struct {
	Username string `json"username`
	Password string `json:"password"`
}

type AdminResponse struct {
	Username  string `json"username`
	Role      string `json:"role"`
	CreatedOn string `json:"created_on"`
}

func (u *UserRequest) Validate() *errs.AppError {
	if u.Username == "" {
		return errs.NewBadRequest("invalid username")
	}

	if u.Password == "" {
		return errs.NewBadRequest("invalid password")
	}

	if u.CustomerID == "" {
		return errs.NewBadRequest("invalid customer id")
	}

	return nil
}

func (u *AdminRequest) Validate() *errs.AppError {
	if u.Username == "" {
		return errs.NewBadRequest("invalid username")
	}

	if u.Password == "" {
		return errs.NewBadRequest("invalid password")
	}

	return nil
}
