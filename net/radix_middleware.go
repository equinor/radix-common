package net

import (
	"net/http"
	"time"

	"github.com/equinor/radix-common/models"
	httpUtils "github.com/equinor/radix-common/net/http"
	"github.com/rs/zerolog"
)

// RadixMiddleware The middleware between router and radix handler functions
type RadixMiddleware struct {
	Path    string
	Method  string
	next    models.RadixHandlerFunc
	handled func(*RadixMiddleware, http.ResponseWriter, *http.Request, time.Time)
}

// NewRadixMiddleware Constructor for radix middleware
func NewRadixMiddleware(path, method string, next models.RadixHandlerFunc, handled func(*RadixMiddleware, http.ResponseWriter, *http.Request, time.Time)) *RadixMiddleware {
	handler := &RadixMiddleware{
		Path:    path,
		Method:  method,
		next:    next,
		handled: handled,
	}
	return handler
}

// Handle Wraps radix handler methods
func (handler *RadixMiddleware) Handle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	logger := zerolog.Ctx(r.Context())
	w.Header().Add("Access-Control-Allow-Origin", "*")

	defer func() {
		if handler.handled != nil {
			handler.handled(handler, w, r, startTime)
		}
	}()

	token, err := httpUtils.GetBearerTokenFromHeader(r)
	if err != nil {
		if err := httpUtils.ErrorResponse(w, r, err); err != nil {
			logger.Error().Err(err).Msg("unable to write auth error response")
		}
	}

	impersonation, err := httpUtils.GetImpersonationFromHeader(r)
	if err != nil {
		if err := httpUtils.ErrorResponse(w, r, httpUtils.UnexpectedError("Problems impersonating", err)); err != nil {
			logger.Error().Err(err).Msg("unable to write impersonating error response")
		}
	}

	accounts := models.NewAccounts(
		token,
		impersonation)

	handler.next(accounts, w, r)
}
