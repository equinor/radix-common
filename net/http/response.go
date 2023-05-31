package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

// Error Representation of errors in the API. These are divided into a small
// number of categories, essentially distinguished by whose fault the
// error is; i.e., is this error:
//   - a transient problem with the service, so worth trying again?
//   - not going to work until the user takes some other action, e.g., updating config?
type Error struct {
	Type Type
	// a message that can be printed out for the user
	Message string `json:"message"`
	// the underlying error that can be e.g., logged for developers to look at
	Err error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

// Type Type of error
type Type string

const (
	// Server The operation looked fine on paper, but something went wrong
	Server Type = "server"
	// Missing The thing you mentioned, whatever it is, just doesn't exist
	Missing = "missing"
	// User The operation was well-formed, but you asked for something that
	// can't happen at present (e.g., because you've not supplied some
	// config yet)
	User = "user"
)

// MarshalJSON Writes error as json
func (e *Error) MarshalJSON() ([]byte, error) {
	var errMsg string
	if e.Err != nil {
		errMsg = e.Err.Error()
	}
	jsonable := &struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Err     string `json:"error,omitempty"`
	}{
		Type:    string(e.Type),
		Message: e.Message,
		Err:     errMsg,
	}
	return json.Marshal(jsonable)
}

// UnmarshalJSON Parses json
func (e *Error) UnmarshalJSON(data []byte) error {
	jsonable := &struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Err     string `json:"error,omitempty"`
	}{}
	if err := json.Unmarshal(data, &jsonable); err != nil {
		return err
	}
	e.Type = Type(jsonable.Type)
	e.Message = jsonable.Message
	if jsonable.Err != "" {
		e.Err = errors.New(jsonable.Err)
	}
	return nil
}

// UnexpectedError any unexpected error
func UnexpectedError(message string, underlyingError error) error {
	return &Error{
		Type:    Server,
		Err:     underlyingError,
		Message: message,
	}
}

// TypeMissingError indication of underlying type missing
func TypeMissingError(message string, underlyingError error) error {
	return &Error{
		Type:    Missing,
		Err:     underlyingError,
		Message: message,
	}
}

// ValidationError Used for indication of validation errors
func ValidationError(kind, message string) error {
	return &Error{
		Type:    User,
		Err:     fmt.Errorf("%s failed validation", kind),
		Message: message,
	}
}

// NotFoundError No found error
func NotFoundError(message string) error {
	return &Error{
		Type:    Missing,
		Message: message,
	}
}

// ApplicationNotFoundError indication that application was not found. Can also mean a user does not have access to the application.
func ApplicationNotFoundError(message string, underlyingError error) error {
	return &Error{
		Type:    Missing,
		Err:     underlyingError,
		Message: message,
	}
}

// CoverAllError Cover all other errors for requester type Type
func CoverAllError(err error, requesterType Type) *Error {
	return &Error{
		Type:    requesterType,
		Err:     err,
		Message: `Error: ` + err.Error(),
	}
}

func writeErrorWithCode(w http.ResponseWriter, r *http.Request, code int, err *Error) {
	// An Accept header with "application/json" is sent by clients
	// understanding how to decode JSON errors. Older clients don't
	// send an Accept header, so we just give them the error text.
	if len(r.Header.Get("Accept")) > 0 {
		switch NegotiateContentType(r, []string{"application/json", "text/plain"}) {
		case "application/json":
			body, encodeErr := json.Marshal(err)
			if encodeErr != nil {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error encoding error response: %s\n\nOriginal error: %s", encodeErr.Error(), err.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(code)
			w.Write(body)
			return
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(code)
			fmt.Fprint(w, err.Message)
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprint(w, err.Error())
}

// StringResponse Used for textual response data. I.e. log data
func StringResponse(w http.ResponseWriter, r *http.Request, result string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// ByteArrayResponse Used for response data. I.e. image
func ByteArrayResponse(w http.ResponseWriter, r *http.Request, contentType string, result []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// JSONResponse Marshals response with header
func JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	body, err := json.Marshal(result)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// ReaderFileResponse writes the content from the reader to the response,
// and sets Content-Disposition=attachment; filename=<filename arg>
func ReaderFileResponse(w http.ResponseWriter, reader io.Reader, fileName, contentType string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, reader)
}

// ReaderResponse writes the content from the reader to the response,
func ReaderResponse(w http.ResponseWriter, reader io.Reader, contentType string) {
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, reader)
}

// ErrorResponse Marshals error for user requester
func ErrorResponse(w http.ResponseWriter, r *http.Request, apiError error) {
	errorResponseFor(User, w, r, apiError)
}

// ErrorResponseForServer Marshals error for server requester
func ErrorResponseForServer(w http.ResponseWriter, r *http.Request, apiError error) {
	errorResponseFor(Server, w, r, apiError)
}

func errorResponseFor(requesterType Type, w http.ResponseWriter, r *http.Request, apiError error) {
	var outErr *Error
	var code int
	var ok bool

	// Skip error response if the context is cancelled.
	// This will typically happen when a HTTP request is cancelled by the caller.
	if errors.Is(apiError, context.Canceled) {
		log.Info(apiError) // Should we log it as info or debug?
		return
	}

	log.Error(apiError)

	err := errors.Cause(apiError)
	if outErr, ok = err.(*Error); !ok {
		outErr = CoverAllError(apiError, requesterType)
	}

	switch e := apiError.(type) {
	case *url.Error:
		// Reflect any underlying network error
		writeErrorWithCode(w, r, http.StatusInternalServerError, outErr)

	case *k8serrors.StatusError:
		writeErrorWithCode(w, r, int(e.ErrStatus.Code), outErr)

	default:
		switch outErr.Type {
		case Missing:
			code = http.StatusNotFound
		case User:
			code = http.StatusBadRequest
		case Server:
			code = http.StatusInternalServerError
		default:
			code = http.StatusInternalServerError
		}
		writeErrorWithCode(w, r, code, outErr)

	}
}

// NegotiateContentType picks a content type based on the Accept
// header from a request, and a supplied list of available content
// types in order of preference. If the Accept header mentions more
// than one available content type, the one with the highest quality
// (`q`) parameter is chosen; if there are a number of those, the one
// that appears first in the available types is chosen.
func NegotiateContentType(r *http.Request, orderedPref []string) string {
	specs, err := GetAccepts(r.Header)
	if err != nil {
		log.Errorf("error getting header Accept: %v", err)
	}
	if len(specs) == 0 {
		return orderedPref[0]
	}

	var preferred []AcceptSpec
	for _, spec := range specs {
		if indexOf(orderedPref, spec.Value) < len(orderedPref) {
			preferred = append(preferred, spec)
		}
	}
	if len(preferred) > 0 {
		sort.Sort(sortAccept{preferred, orderedPref})
		return preferred[0].Value
	}
	return ""
}
