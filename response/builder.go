package response

import (
	"errors"
	"net/http"
)

var DefaultResponseBuilder = NewResponseBuilder(
	Config{
		Formatter: JSONFormatter{},
	},
)

type ResponseData struct {
	Content interface{}
	Header  http.Header
	Status  int
}

type Formatter interface {
	FormatBody(ResponseData) []byte
	FormatHeader(ResponseData) http.Header
}

type ResponseError interface {
	Data() interface{}
	Error() string
	Message() string
	Status() int
}

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
	if r.responseData.Status == 0 {
		return http.StatusOK
	}
	return r.responseData.Status
}

// WithContent returns a copy of the response with the supplied content.
// If the content argument implements the ResponseError interface,
// the status will be set to the status of the ResponseError.
func (r response) WithContent(content interface{}) response {

	if err, ok := content.(error); ok {
		var responseError ResponseError
		if errors.Is(err, responseError) {
			r.responseData.Status = responseError.Status()
		}
	}

	r.responseData.Content = content
	return r
}

// WithStatus returns a copy of the response with the supplied status
func (r response) WithStatus(status int) response {
	r.responseData.Status = status
	return r
}

type responseBuilder struct {
	Config Config
}

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
		responseData: ResponseData{},
	}.WithContent(content)
}

// Status returns a new response with the supplied status
func (r responseBuilder) Status(status int) response {
	return response{
		responseData: ResponseData{},
	}.WithStatus(status)
}

/*
"Status shortcut methods"
*/
func (r responseBuilder) BadGateway() response {
	return r.Status(http.StatusBadGateway)
}

func (r responseBuilder) BadRequest() response {
	return r.Status(http.StatusBadRequest)
}

func (r responseBuilder) Conflict() response {
	return r.Status(http.StatusConflict)
}

func (r responseBuilder) Created() response {
	return r.Status(http.StatusCreated)
}

func (r responseBuilder) Forbidden() response {
	return r.Status(http.StatusForbidden)
}

func (r responseBuilder) Found() response {
	return r.Status(http.StatusFound)
}

func (r responseBuilder) GatewayTimeout() response {
	return r.Status(http.StatusGatewayTimeout)
}

func (r responseBuilder) InternalServerError() response {
	return r.Status(http.StatusInternalServerError)
}

func (r responseBuilder) OK() response {
	return r.Status(http.StatusOK)
}

func (r responseBuilder) MovedPermanently() response {
	return r.Status(http.StatusMovedPermanently)
}

func (r responseBuilder) NotFound() response {
	return r.Status(http.StatusNotFound)
}

func (r responseBuilder) NotImplemented() response {
	return r.Status(http.StatusNotImplemented)
}

func (r responseBuilder) Unauthorized() response {
	return r.Status(http.StatusUnauthorized)
}
