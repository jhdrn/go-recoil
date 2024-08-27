package response

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterFormatBody(t *testing.T) {

	responseData := ResponseData{
		Content: map[string]any{
			"key": "value",
		},
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, `{"key":"value"}`, string(body))
}

func TestJSONFormatterFormatStream(t *testing.T) {

	body := []byte("body")

	responseData := ResponseData{
		Content: bytes.NewReader(body),
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	responseBody, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, responseBody, body)
}

func TestJSONFormatterFormatBodyNilContent(t *testing.T) {

	responseData := ResponseData{
		Content: nil,
		Status:  http.StatusBadRequest,
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, fmt.Sprintf(`{"message":"%v"}`, http.StatusText(responseData.Status)), string(body))
}

func TestJSONFormatterFormatBodyErrorContent(t *testing.T) {

	err := errors.New("some error")

	responseData := ResponseData{
		Content: err,
		Status:  http.StatusBadRequest,
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, `{"message":"some error"}`, string(body))
}

func TestJSONFormatterFormatBodyBadContent(t *testing.T) {

	responseData := ResponseData{
		Content: func() {},
	}

	f := JSONFormatter{}

	testPanic := func() {
		f.FormatBody(responseData)
	}

	assert.Panics(t, testPanic, "did not panic on bad JSON marshalling input type")
}

func TestJSONFormatterFormatHeader(t *testing.T) {

	responseData := ResponseData{
		Content: map[string]any{
			"key": "value",
		},
		Header: http.Header{},
	}

	responseData.Header.Set("key", "value")

	f := JSONFormatter{}

	header := f.FormatHeader(responseData)

	assert.Equal(t, "application/json", header.Get("Content-Type"))
	assert.Equal(t, "value", header.Get("key"))
	assert.Equal(t, 2, len(header))
}

func TestJSONFormatterFormatStatus(t *testing.T) {

	responseData := ResponseData{
		Status: http.StatusBadRequest,
	}

	f := JSONFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusBadRequest, status)
}

func TestJSONFormatterFormatStatusZeroValue(t *testing.T) {

	responseData := ResponseData{}

	f := JSONFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusOK, status)
}
