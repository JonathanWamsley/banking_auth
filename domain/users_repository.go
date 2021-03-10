package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking_auth/errs"
	"github.com/jonathanwamsley/banking_auth/logger"
)

const (
	queryInsertUser  = "insert into banking.users(username, password, role, customer_id, created_on) value(?, ?, ?, ?, ?);"
	queryInsertAdmin = "insert into banking.users(username, password, role, created_on) value(?, ?, ?, ?);"
	queryGetUsers    = `select username, customer_id from banking.users where role = "user";`
)

type UsersRepositoryDB struct {
	client *sqlx.DB
}

func NewUsersRepository(client *sqlx.DB) UsersRepositoryDB {
	return UsersRepositoryDB{client}
}

func (d UsersRepositoryDB) CreateUser(u User) (*User, *errs.AppError) {
	_, err := d.client.Exec(queryInsertUser, u.Username, u.Password, u.Role, u.CustomerID, u.CreatedOn)
	if err != nil {
		logger.Error("error while creating new user " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	u.Password = ""
	return &u, nil
}

func (d UsersRepositoryDB) CreateAdmin(u User) (*User, *errs.AppError) {
	_, err := d.client.Exec(queryInsertAdmin, u.Username, u.Password, u.Role, u.CreatedOn)
	if err != nil {
		logger.Error("error while creating new admin " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	u.Password = ""
	return &u, nil
}

func (d UsersRepositoryDB) GetUsers() ([]User, *errs.AppError) {
	users := make([]User, 0)
	err := d.client.Select(&users, queryGetUsers)
	if err != nil {
		logger.Error("error while querying new user " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	return users, nil
}
