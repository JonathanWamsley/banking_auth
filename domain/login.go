package domain

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanwamsley/banking_auth/errs"
	"github.com/jonathanwamsley/banking_auth/logger"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerID sql.NullString `db:"customer_id"`
	Accounts   []string       `db:"account_numbers"`
}

func (l Login) GenerateToken() (*string, *errs.AppError) {
	var claims jwt.MapClaims
	if l.CustomerID.Valid {
		claims = l.claimsForUser()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("Failed while signing token: " + err.Error())
		return nil, errs.NewUnexpectedError("cannot generate token")
	}
	return &signedTokenAsString, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	//accounts := strings.Split(l.Accounts.String, ",")

	return jwt.MapClaims{
		"customer_id": l.CustomerID.String,
		"username":    l.Username,
		"accounts":    l.Accounts,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
