package response

import (
	"net/http"
)

// DefaultResponseBuilder is a response builder with a default configuration
// thats is configured to use a JSONFormatter.
var DefaultResponseBuilder = NewResponseBuilder(
	Config{
		Formatter: JSONFormatter{},
	},
)

// ResponseData contains the data to be used to format a response
type ResponseData struct {
	Content interface{}
	Header  http.Header
	Status  int
}

// Formatter is an interface that defines the methods used to format a response
type Formatter interface {
	FormatBody(ResponseData) []byte
	FormatHeader(ResponseData) http.Header
	FormatStatus(ResponseData) int
}

// Config contains the configuration for a response builder
type Config struct {
	Formatter Formatter
}

type response struct {
	config       Config
	responseData ResponseData
}

// Body returns the content to be written to the response
func (r response) Body() []byte {
	return r.config.Formatter.FormatBody(r.responseData)
}

// Header returns the header map
func (r response) Header() http.Header {
	return r.config.Formatter.FormatHeader(r.responseData)
}

// Status returns the HTTP status code. If no status code has been set,
// 200 OK will be returned.
func (r response) Status() int {
	return r.config.Formatter.FormatStatus(r.responseData)
}

// WithContent returns a copy of the response with the supplied content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r response) WithContent(content interface{}) response {
	r.responseData.Content = content
	return r
}

// WithHeader returns a copy of the response with the supplied header. If the
// header already exists, the values will replace the existing values.
func (r response) WithHeader(name string, value ...string) response {
	r.responseData.Header[name] = value
	return r
}

// WithStatus returns a copy of the response with the supplied status
func (r response) WithStatus(status int) response {
	r.responseData.Status = status
	return r
}

// WithCookie returns a copy of the response with the supplied cookie set
// to the response header
func (r response) WithCookie(cookie *http.Cookie) response {
	if v := cookie.String(); v != "" {
		r.WithHeader("Set-Cookie", v)
	}
	return r
}

type responseBuilder struct {
	Config Config
}

// NewResponseBuilder returns a new response builder with the supplied
// configuration
func NewResponseBuilder(c Config) *responseBuilder {
	return &responseBuilder{
		Config: c,
	}
}

// Content returns a new response with the supplied content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r responseBuilder) Content(content interface{}) response {
	return response{
		config: r.Config,
		responseData: ResponseData{
			Header: make(http.Header),
		},
	}.WithContent(content)
}

// Status returns a new response with the supplied status
func (r responseBuilder) Status(status int) response {
	return response{
		config: r.Config,
		responseData: ResponseData{
			Header: make(http.Header),
		},
	}.WithStatus(status)
}

// BadGateway returns a new response with the status code 502 Bad Gateway
func (r responseBuilder) BadGateway() response {
	return r.Status(http.StatusBadGateway)
}

// BadRequest returns a new response with the status code 400 Bad Request
func (r responseBuilder) BadRequest() response {
	return r.Status(http.StatusBadRequest)
}

// Conflict returns a new response with the status code 409 Conflict
func (r responseBuilder) Conflict() response {
	return r.Status(http.StatusConflict)
}

// Created returns a new response with the status code 201 Created
func (r responseBuilder) Created() response {
	return r.Status(http.StatusCreated)
}

// Forbidden returns a new response with the status code 403 Forbidden
func (r responseBuilder) Forbidden() response {
	return r.Status(http.StatusForbidden)
}

// Found returns a new response with the status code 302 Found
func (r responseBuilder) Found() response {
	return r.Status(http.StatusFound)
}

// GatewayTimeout returns a new response with the status code 504 Gateway Timeout
func (r responseBuilder) GatewayTimeout() response {
	return r.Status(http.StatusGatewayTimeout)
}

// Gone returns a new response with the status code 410 Gone
func (r responseBuilder) InternalServerError() response {
	return r.Status(http.StatusInternalServerError)
}

// OK returns a new response with the status code 200 OK
func (r responseBuilder) OK() response {
	return r.Status(http.StatusOK)
}

// MovedPermanently returns a new response with the status code 301 Moved Permanently
func (r responseBuilder) MovedPermanently() response {
	return r.Status(http.StatusMovedPermanently)
}

// NotFound returns a new response with the status code 404 Not Found
func (r responseBuilder) NotFound() response {
	return r.Status(http.StatusNotFound)
}

// NotImplemented returns a new response with the status code 501 Not Implemented
func (r responseBuilder) NotImplemented() response {
	return r.Status(http.StatusNotImplemented)
}

// Unauthorized returns a new response with the status code 401 Unauthorized
func (r responseBuilder) Unauthorized() response {
	return r.Status(http.StatusUnauthorized)
}
