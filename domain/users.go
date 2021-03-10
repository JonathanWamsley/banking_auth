package domain

import (
	"github.com/jonathanwamsley/banking_auth/dto"
	"github.com/jonathanwamsley/banking_auth/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type User struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	Role       string `db:"role"`
	CustomerID string `db:"customer_id"`
	CreatedOn  string `db:"created_on"`
}

type UsersRepository interface {
	CreateUser(User) (*User, *errs.AppError)
	CreateAdmin(User) (*User, *errs.AppError)
	GetUsers() ([]User, *errs.AppError)
}

func NewUser(u dto.UserRequest) User {
	return User{
		Username:   u.Username,
		Password:   u.Password,
		Role:       "user",
		CustomerID: u.CustomerID,
		CreatedOn:  dbTSLayout,
	}
}

func NewAdmin(a dto.AdminRequest) User {
	return User{
		Username:  a.Username,
		Password:  a.Password,
		Role:      "admin",
		CreatedOn: dbTSLayout,
	}
}

func UserResponse(u User) dto.UserResponse {
	return dto.UserResponse{
		Username:   u.Username,
		CustomerID: u.CustomerID,
		CreatedOn:  u.CreatedOn,
	}
}

func AdminResponse(a User) dto.AdminResponse {
	return dto.AdminResponse{
		Username:  a.Username,
		Role:      a.Role,
		CreatedOn: a.CreatedOn,
	}
}
