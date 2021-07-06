package models

import "net/http"

// Controller Pattern of an rest/stream controller
type Controller interface {
	GetRoutes() Routes
}

// DefaultController Default implementation
type DefaultController struct {
}

// Routes Holder of all routes
type Routes []Route

// Route Describe route
type Route struct {
	Path        string
	Method      string
	HandlerFunc RadixHandlerFunc
}

// RadixHandlerFunc Pattern for handler functions
type RadixHandlerFunc func(Accounts, http.ResponseWriter, *http.Request)
