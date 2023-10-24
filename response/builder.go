package response

import (
	"io"
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
	Body   any
	Header http.Header
	Status int
}

// Formatter is an interface that defines the methods used to format a response
type Formatter interface {
	FormatBody(ResponseData) io.Reader
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
func (r response) Body() io.Reader {
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

// WithContent returns a copy of the response with the given content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r response) WithContent(content any) response {
	r.responseData.Body = content
	return r
}

// WithStream returns a copy of the response with the given stream.
func (r response) WithStream(stream io.Reader) response {
	r.responseData.Body = stream
	return r
}

// WithHeader returns a copy of the response with the given header. It will
// replace the existing header.
func (r response) WithHeader(header http.Header) response {
	r.responseData.Header = header
	return r
}

// WithHeaderEntry returns a copy of the response with the given header entry.
// If a header entry with the same key already exists, the existing values will
// be replaced.
func (r response) WithHeaderEntry(key string, value ...string) response {
	r.responseData.Header[key] = value
	return r
}

// WithStatus returns a copy of the response with the given status
func (r response) WithStatus(status int) response {
	r.responseData.Status = status
	return r
}

// WithCookie returns a copy of the response with the given cookie set
// to the response header
func (r response) WithCookie(cookie *http.Cookie) response {
	if v := cookie.String(); v != "" {
		r.WithHeaderEntry("Set-Cookie", v)
	}
	return r
}

type ResponseBuilder struct {
	Config Config
}

// NewResponseBuilder returns a new response builder with the given
// configuration. Will panic if the configuration is missing a formatter.
func NewResponseBuilder(c Config) *ResponseBuilder {
	if c.Formatter == nil {
		panic("config formatter is nil")
	}

	return &ResponseBuilder{
		Config: c,
	}
}

// Content returns a new response with the given content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r ResponseBuilder) Content(content any) response {
	return response{
		config: r.Config,
		responseData: ResponseData{
			Header: make(http.Header),
		},
	}.WithContent(content)
}

// Status returns a new response with the given status
func (r ResponseBuilder) Status(status int) response {
	return response{
		config: r.Config,
		responseData: ResponseData{
			Header: make(http.Header),
		},
	}.WithStatus(status)
}

// BadGateway returns a new response with the status code 502 Bad Gateway
func (r ResponseBuilder) BadGateway() response {
	return r.Status(http.StatusBadGateway)
}

// BadRequest returns a new response with the status code 400 Bad Request
func (r ResponseBuilder) BadRequest() response {
	return r.Status(http.StatusBadRequest)
}

// Conflict returns a new response with the status code 409 Conflict
func (r ResponseBuilder) Conflict() response {
	return r.Status(http.StatusConflict)
}

// Created returns a new response with the status code 201 Created
func (r ResponseBuilder) Created() response {
	return r.Status(http.StatusCreated)
}

// Forbidden returns a new response with the status code 403 Forbidden
func (r ResponseBuilder) Forbidden() response {
	return r.Status(http.StatusForbidden)
}

// Found returns a new response with the status code 302 Found
// Permanently and the given location set as the Location header.
func (r ResponseBuilder) Found(location string) response {
	return r.Status(http.StatusFound).WithHeaderEntry("Location", location)
}

// GatewayTimeout returns a new response with the status code 504 Gateway Timeout
func (r ResponseBuilder) GatewayTimeout() response {
	return r.Status(http.StatusGatewayTimeout)
}

// Gone returns a new response with the status code 410 Gone
func (r ResponseBuilder) InternalServerError() response {
	return r.Status(http.StatusInternalServerError)
}

// OK returns a new response with the status code 200 OK
func (r ResponseBuilder) OK() response {
	return r.Status(http.StatusOK)
}

// MovedPermanently returns a new response with the status code 301 Moved
// Permanently and the given location set as the Location header.
func (r ResponseBuilder) MovedPermanently(location string) response {
	return r.Status(http.StatusMovedPermanently).WithHeaderEntry("Location", location)
}

// NotFound returns a new response with the status code 404 Not Found
func (r ResponseBuilder) NotFound() response {
	return r.Status(http.StatusNotFound)
}

// NotImplemented returns a new response with the status code 501 Not Implemented
func (r ResponseBuilder) NotImplemented() response {
	return r.Status(http.StatusNotImplemented)
}

// PermanentRedirect returns a new response with the status code 308 Permanent Redirect
// and the given location set as the Location header.
func (r ResponseBuilder) PermanentRedirect(location string) response {
	return r.Status(http.StatusPermanentRedirect).WithHeaderEntry("Location", location)
}

// TemporaryRedirect returns a new response with the status code 308 Temporary Redirect
// and the given location set as the Location header.
func (r ResponseBuilder) TemporaryRedirect(location string) response {
	return r.Status(http.StatusTemporaryRedirect).WithHeaderEntry("Location", location)
}

// Unauthorized returns a new response with the status code 401 Unauthorized
func (r ResponseBuilder) Unauthorized() response {
	return r.Status(http.StatusUnauthorized)
}
