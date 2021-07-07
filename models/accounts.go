package models

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

// NewAccounts creates a new Accounts struct
func NewAccounts(
	token string,
	impersonation Impersonation) Accounts {

	return Accounts{
		token:         token,
		impersonation: impersonation,
	}
}

// Accounts contains accounts for accessing k8s API.
type Accounts struct {
	token         string
	impersonation Impersonation
}

// GetUserAccountUserPrincipleName get the user principle name represented in UserAccount
func (accounts Accounts) GetUserAccountUserPrincipleName() (string, error) {
	if accounts.impersonation.PerformImpersonation() {
		return accounts.impersonation.User, nil
	}

	return GetUserPrincipleNameFromToken(accounts.token)
}

// GetToken get the user token
func (accounts Accounts) GetToken() string {
	return accounts.token
}

func GetUserPrincipleNameFromToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	_, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return "", fmt.Errorf("could not parse token (%v)", err)
	}

	userPrincipleName := fmt.Sprintf("%v", claims["upn"])
	return userPrincipleName, nil
}
