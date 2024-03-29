package response

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContent(t *testing.T) {
	r := Content(map[string]any{"key": "value"})

	bodyReader := r.Body()

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, `{"key":"value"}`, string(body))
	assert.Equal(t, http.StatusOK, r.Status())
}

func TestStatus(t *testing.T) {
	r := Status(http.StatusConflict)

	assert.Equal(t, http.StatusConflict, r.Status())
}

func TestBadGateway(t *testing.T) {
	r := BadGateway()

	assert.Equal(t, http.StatusBadGateway, r.Status())
}

func TestBadRequest(t *testing.T) {
	r := BadRequest()

	assert.Equal(t, http.StatusBadRequest, r.Status())
}

func TestConflict(t *testing.T) {
	r := Conflict()

	assert.Equal(t, http.StatusConflict, r.Status())
}

func TestCreated(t *testing.T) {
	r := Created()

	assert.Equal(t, http.StatusCreated, r.Status())
}

func TestForbidden(t *testing.T) {
	r := Forbidden()

	assert.Equal(t, http.StatusForbidden, r.Status())
}

func TestFound(t *testing.T) {
	r := Found("https://example.org")

	assert.Equal(t, http.StatusFound, r.Status())
	assert.Equal(t, "https://example.org", r.Header().Get("Location"))
}

func TestGatewayTimeout(t *testing.T) {
	r := GatewayTimeout()

	assert.Equal(t, http.StatusGatewayTimeout, r.Status())
}

func TestInternalServerError(t *testing.T) {
	r := InternalServerError()

	assert.Equal(t, http.StatusInternalServerError, r.Status())
}

func TestOK(t *testing.T) {
	r := OK()

	assert.Equal(t, http.StatusOK, r.Status())
}

func TestMovedPermanently(t *testing.T) {
	r := MovedPermanently("https://example.org")

	assert.Equal(t, http.StatusMovedPermanently, r.Status())
	assert.Equal(t, "https://example.org", r.Header().Get("Location"))
}

func TestNotFound(t *testing.T) {
	r := NotFound()

	assert.Equal(t, http.StatusNotFound, r.Status())
}

func TestNotImplemented(t *testing.T) {
	r := NotImplemented()

	assert.Equal(t, http.StatusNotImplemented, r.Status())
}

func TestPermanentRedirect(t *testing.T) {
	r := PermanentRedirect("https://example.org")

	assert.Equal(t, http.StatusPermanentRedirect, r.Status())
	assert.Equal(t, "https://example.org", r.Header().Get("Location"))
}

func TestTemporaryRedirect(t *testing.T) {
	r := TemporaryRedirect("https://example.org")

	assert.Equal(t, http.StatusTemporaryRedirect, r.Status())
	assert.Equal(t, "https://example.org", r.Header().Get("Location"))
}

func TestUnauthorized(t *testing.T) {
	r := Unauthorized()

	assert.Equal(t, http.StatusUnauthorized, r.Status())
}
