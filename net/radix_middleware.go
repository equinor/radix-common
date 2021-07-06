package net

import (
	"github.com/equinor/radix-common/models"
	httpUtils "github.com/equinor/radix-common/net/http"
	"net/http"
	"time"
)

// RadixMiddleware The middleware between router and radix handler functions
type RadixMiddleware struct {
	path   string
	method string
	next   models.RadixHandlerFunc
}

// NewRadixMiddleware Constructor for radix middleware
func NewRadixMiddleware(path, method string, next models.RadixHandlerFunc) *RadixMiddleware {
	handler := &RadixMiddleware{
		path,
		method,
		next,
	}

	return handler
}

// Handle Wraps radix handler methods
func (handler *RadixMiddleware) Handle(w http.ResponseWriter, r *http.Request, done func(*RadixMiddleware, http.ResponseWriter, *http.Request, time.Time)) {
	startTime := time.Now()

	w.Header().Add("Access-Control-Allow-Origin", "*")

	defer func() {
		done(handler, w, r, startTime)
	}()

	token, err := httpUtils.GetBearerTokenFromHeader(r)
	if err != nil {
		httpUtils.ErrorResponse(w, r, err)
		return
	}

	impersonation, err := httpUtils.GetImpersonationFromHeader(r)
	if err != nil {
		httpUtils.ErrorResponse(w, r, httpUtils.UnexpectedError("Problems impersonating", err))
		return
	}

	accounts := models.NewAccounts(
		token,
		impersonation)

	handler.next(accounts, w, r)
}
