package response

import (
	"io"
	"net/http"
)

// DefaultConfig is the default configuration for a response builder. It uses
// the JSONFormatter to format responses, but can be modified to use a different
// formatter.
var DefaultConfig = Config{
	Formatter: JSONFormatter{},
}

// ResponseData contains the data to be used to format a response.
type ResponseData struct {
	Body   any
	Header http.Header
	Status int
}

// Formatter is an interface that defines the methods used to format a response.
type Formatter interface {
	FormatBody(ResponseData) io.Reader
	FormatHeader(ResponseData) http.Header
	FormatStatus(ResponseData) int
}

// Config contains the configuration for a response builder.
type Config struct {
	Formatter Formatter
}

// Builder is a builder for creating responses. It implements the
// recoil.Response interface.
type Builder struct {
	config       Config
	responseData ResponseData
}

// Option is a functional option for configuring a response builder.
type Option func(*Builder)

// WithConfig configures the response builder with the given configuration.
func WithConfig(c Config) Option {
	return func(b *Builder) {
		b.config = c
	}
}

// NewBuilder returns a new response builder. By default it uses DefaultConfig
// but can be configured to use another Config using the WithConfig function.
// Will panic if the configuration is missing a formatter.
func NewBuilder(options ...Option) Builder {
	builder := Builder{
		config: DefaultConfig,
		responseData: ResponseData{
			Header: make(http.Header),
		},
	}

	// Apply all the functional options to configure the ALBApp.
	for _, opt := range options {
		opt(&builder)
	}

	if builder.config.Formatter == nil {
		panic("config formatter is nil")
	}

	return builder
}

// Body returns the content to be written to the response.
func (r Builder) Body() io.Reader {
	return r.config.Formatter.FormatBody(r.responseData)
}

// Header returns the header map.
func (r Builder) Header() http.Header {
	return r.config.Formatter.FormatHeader(r.responseData)
}

// Status returns the HTTP status code. If no status code has been set,
// 200 OK will be returned.
func (r Builder) Status() int {
	return r.config.Formatter.FormatStatus(r.responseData)
}

// WithContent returns a copy of the response with the given content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r Builder) WithContent(content any) Builder {
	r.responseData.Body = content
	return r
}

// WithStream returns a copy of the response with the given stream.
func (r Builder) WithStream(stream io.Reader) Builder {
	r.responseData.Body = stream
	return r
}

// WithHeader returns a copy of the response with the given header. It will
// replace the existing header.
func (r Builder) WithHeader(header http.Header) Builder {
	r.responseData.Header = header
	return r
}

// WithHeaderEntry returns a copy of the response with the given header entry.
// If a header entry with the same key already exists, the existing values will
// be replaced.
func (r Builder) WithHeaderEntry(key string, value ...string) Builder {
	r.responseData.Header[key] = value
	return r
}

// WithStatus returns a copy of the response with the given status.
func (r Builder) WithStatus(status int) Builder {
	r.responseData.Status = status
	return r
}

// WithCookie returns a copy of the response with the given cookie set
// to the response header.
func (r Builder) WithCookie(cookie *http.Cookie) Builder {
	if v := cookie.String(); v != "" {
		r.WithHeaderEntry("Set-Cookie", v)
	}
	return r
}

// BadGateway returns a new response with the status code 502 Bad Gateway.
func (r Builder) BadGateway() Builder {
	return r.WithStatus(http.StatusBadGateway)
}

// BadRequest returns a new response with the status code 400 Bad Request.
func (r Builder) BadRequest() Builder {
	return r.WithStatus(http.StatusBadRequest)
}

// Conflict returns a new response with the status code 409 Conflict.
func (r Builder) Conflict() Builder {
	return r.WithStatus(http.StatusConflict)
}

// Created returns a new response with the status code 201 Created.
func (r Builder) Created() Builder {
	return r.WithStatus(http.StatusCreated)
}

// Forbidden returns a new response with the status code 403 Forbidden.
func (r Builder) Forbidden() Builder {
	return r.WithStatus(http.StatusForbidden)
}

// Found returns a new response with the status code 302 Found
// Permanently and the given location set as the Location header.
func (r Builder) Found(location string) Builder {
	return r.WithStatus(http.StatusFound).WithHeaderEntry("Location", location)
}

// GatewayTimeout returns a new response with the status code 504 Gateway Timeout.
func (r Builder) GatewayTimeout() Builder {
	return r.WithStatus(http.StatusGatewayTimeout)
}

// Gone returns a new response with the status code 410 Gone.
func (r Builder) InternalServerError() Builder {
	return r.WithStatus(http.StatusInternalServerError)
}

// OK returns a new response with the status code 200 OK.
func (r Builder) OK() Builder {
	return r.WithStatus(http.StatusOK)
}

// MovedPermanently returns a new response with the status code 301 Moved
// Permanently and the given location set as the Location header.
func (r Builder) MovedPermanently(location string) Builder {
	return r.WithStatus(http.StatusMovedPermanently).WithHeaderEntry("Location", location)
}

// NotFound returns a new response with the status code 404 Not Found.
func (r Builder) NotFound() Builder {
	return r.WithStatus(http.StatusNotFound)
}

// NotImplemented returns a new response with the status code 501 Not Implemented.
func (r Builder) NotImplemented() Builder {
	return r.WithStatus(http.StatusNotImplemented)
}

// PermanentRedirect returns a new response with the status code 308 Permanent Redirect
// and the given location set as the Location header.
func (r Builder) PermanentRedirect(location string) Builder {
	return r.WithStatus(http.StatusPermanentRedirect).WithHeaderEntry("Location", location)
}

// TemporaryRedirect returns a new response with the status code 308 Temporary Redirect
// and the given location set as the Location header.
func (r Builder) TemporaryRedirect(location string) Builder {
	return r.WithStatus(http.StatusTemporaryRedirect).WithHeaderEntry("Location", location)
}

// Unauthorized returns a new response with the status code 401 Unauthorized.
func (r Builder) Unauthorized() Builder {
	return r.WithStatus(http.StatusUnauthorized)
}
