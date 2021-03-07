package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking_auth/errs"
	"github.com/jonathanwamsley/banking_auth/logger"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
}

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client}
}

// queries
const (
	validLogin = "select customer_id from banking.users where username = ? and password = ?;"

	getAccountID = "SELECT a.account_id as account_numbers FROM banking.users u LEFT JOIN banking.accounts a ON a.customer_id = u.customer_id WHERE u.username = ? and password = ?;"
)

// FindBy verifies a login credentials and then returns associated accounts for that customer
func (d AuthRepositoryDB) FindBy(username, password string) (*Login, *errs.AppError) {
	var login Login
	// first verify matching username/password
	err := d.client.Get(&login, validLogin, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewUnexpectedError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	// now get the verify users account ids
	accounts := make([]string, 0)
	err = d.client.Select(&accounts, getAccountID, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			// pass, a user is allowed to have no accounts
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	login.Accounts = accounts
	return &login, nil
}
