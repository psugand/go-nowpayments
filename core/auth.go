package core

import (
	"fmt"
	"strings"
)

type token struct {
	Token string `json:"token"`
}

// Authenticate is used for obtaining a JWT token.
// JWT is required only for payout request API call
func Authenticate(email, password string) (string, error) {
	r := strings.NewReader(fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
		}`, email, password))

	t := &token{}

	par := &SendParams{
		RouteName: "auth",
		Body:      r,
		Into:      &t,
	}

	err := HTTPSend(par)
	return t.Token, err
}
