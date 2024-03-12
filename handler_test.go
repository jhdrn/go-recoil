package recoil

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type response struct {
	body   io.Reader
	header http.Header
	status int
}

func (r response) Body() io.Reader {
	return r.body
}
func (r response) Header() http.Header {
	return r.header
}
func (r response) Status() int {
	return r.status
}

func TestHandler(t *testing.T) {

	body := []byte("body")

	responseObj := response{
		body: bytes.NewReader(body),
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

	assert.Equal(t, body, respBody)
	assert.Equal(t, responseObj.header, rw.Result().Header)
	assert.Equal(t, responseObj.status, rw.Code)
}

type closer struct {
	io.Reader
	closed bool
}

func (c *closer) Close() error {
	if c.closed {
		return errors.New("already closed")
	}
	c.closed = true
	return nil
}

func TestCloserBody(t *testing.T) {

	body := []byte("body")

	closer := &closer{bytes.NewReader(body), false}

	responseObj := response{
		body: closer,
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

	assert.Equal(t, closer.closed, true)

	// should panic if closed again
	assert.Panics(t, func() {
		h.ServeHTTP(rw, r)
	})
}

func TestHandlerFunc(t *testing.T) {

	body := []byte("body")
	responseObj := response{
		body: bytes.NewReader(body),
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

	assert.Equal(t, body, respBody)
	assert.Equal(t, responseObj.header, rw.Result().Header)
	assert.Equal(t, responseObj.status, rw.Code)
}
