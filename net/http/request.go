package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/equinor/radix-common/models"
)

// GetBearerTokenFromHeader gets bearer token from request header
func GetBearerTokenFromHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("authorization")
	authArr := strings.Split(authorizationHeader, " ")
	var jwtToken string

	if len(authArr) != 2 {
		return "", errors.New("Authentication header is invalid: " + authorizationHeader)
	}

	jwtToken = authArr[1]
	return jwtToken, nil
}

// GetImpersonationFromHeader Gets Impersonation from request header
func GetImpersonationFromHeader(r *http.Request) (models.Impersonation, error) {
	impersonateUser := r.Header.Get("Impersonate-User")
	impersonateGroups := strings.Split(r.Header.Get("Impersonate-Group"), ",")
	return models.NewImpersonation(impersonateUser, impersonateGroups)
}

// GetTokenFromQuery Gets token from query of the request
func GetTokenFromQuery(request *http.Request) string {
	return request.URL.Query().Get("token")
}
