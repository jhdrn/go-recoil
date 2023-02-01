package recoil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type response struct {
	body   []byte
	header http.Header
	status int
}

func (r response) Body() []byte {
	return r.body
}
func (r response) Header() http.Header {
	return r.header
}
func (r response) Status() int {
	return r.status
}

func TestHandler(t *testing.T) {

	responseObj := response{
		body: []byte("body"),
		header: http.Header{
			"Content-Type": []string{"text/plain"},
		},
		status: 200,
	}

	h := Handler(func(r *http.Request) Response {
		return responseObj
	})

	rw := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://example.org", nil)
	h.ServeHTTP(rw, r)

	respBody, _ := io.ReadAll(rw.Body)

	assert.Equal(t, responseObj.body, respBody)
	assert.Equal(t, responseObj.header, rw.Result().Header)
	assert.Equal(t, responseObj.status, rw.Code)
}

func TestHandlerFunc(t *testing.T) {

	responseObj := response{
		body: []byte("body"),
		header: http.Header{
			"Content-Type": []string{"text/plain"},
		},
		status: 200,
	}

	h := HandlerFunc(func(r *http.Request) Response {
		return responseObj
	})

	rw := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://example.org", nil)

	h(rw, r)

	respBody, _ := io.ReadAll(rw.Body)

	assert.Equal(t, responseObj.body, respBody)
	assert.Equal(t, responseObj.header, rw.Result().Header)
	assert.Equal(t, responseObj.status, rw.Code)
}
