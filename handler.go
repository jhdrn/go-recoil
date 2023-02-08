// Package: recoil
package recoil

import (
	"net/http"
)

// Response is an interface that defines the methods used to format a response
type Response interface {
	Body() []byte
	Header() http.Header
	Status() int
}

// The Handler type implements the http.Handler interface which to allow the use
// of ordinary functions as HTTP handlers.
type Handler func(r *http.Request) Response

// ServeHTTP calls Handler(r) and writes the Response to w.
func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := f(r)

	for k, v := range response.Header() {
		w.Header()[k] = v
	}
	w.WriteHeader(response.Status())
	w.Write(response.Body())
}

// HandlerFunc creates a standard library compatible handler function
func HandlerFunc(h Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
