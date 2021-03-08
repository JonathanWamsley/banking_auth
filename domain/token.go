package domain

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanwamsley/banking_auth/errs"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

type Claims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"exp"`
}


func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, *errs.AppError) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, errs.NewValidationError("invalid json format for token")
	}
	var c Claims
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return nil,errs.NewValidationError("invalid token claim")
	}
	return &c, nil
}

func (c Claims) IsValidCustomerId(customerId string) bool {
	return c.CustomerId == customerId
}

func (c Claims) IsValidAccountId(accountId string) bool {
	if accountId != "" {
		accountFound := false
		for _, a := range c.Accounts {
			if a == accountId {
				accountFound = true
				break
			}
		}
		return accountFound
	}
	return true
}

func (c Claims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.CustomerId != urlParams["customer_id"] {
		return false
	}

	if !c.IsValidAccountId(urlParams["account_id"]) {
		return false
	}
	return true
}