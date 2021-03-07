package domain

import (
	"database/sql"
	"time"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerID sql.NullString `db:"customer_id"`
	Accounts   []string       `db:"account_numbers"`
	Role       string         `db:"role"`
}
