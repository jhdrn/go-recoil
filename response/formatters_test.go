package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterFormatBody(t *testing.T) {

	responseData := ResponseData{
		Body: map[string]interface{}{
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
		Body: bytes.NewReader(body),
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	responseBody, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, responseBody, body)
}

func TestJSONFormatterFormatBodyNilContent(t *testing.T) {

	responseData := ResponseData{
		Body:   nil,
		Status: http.StatusBadRequest,
	}

	f := JSONFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, fmt.Sprintf(`{"message":"%v"}`, http.StatusText(responseData.Status)), string(body))
}

func TestJSONFormatterFormatBodyBadContent(t *testing.T) {

	responseData := ResponseData{
		Body: func() {},
	}

	f := JSONFormatter{}

	testPanic := func() {
		f.FormatBody(responseData)
	}

	assert.Panics(t, testPanic, "did not panic on bad JSON marshalling input type")
}

func TestJSONFormatterFormatHeader(t *testing.T) {

	responseData := ResponseData{
		Body: map[string]interface{}{
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

func TestXMLFormatterFormatBody(t *testing.T) {

	type xmlTest struct {
		Key string `xml:"key"`
	}

	responseData := ResponseData{
		Body: xmlTest{
			Key: "value",
		},
	}

	f := XMLFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<xmlTest><key>value</key></xmlTest>", string(body))
}

func TestXMLFormatterFormatStream(t *testing.T) {

	body := []byte("body")

	responseData := ResponseData{
		Body: bytes.NewReader(body),
	}

	f := XMLFormatter{}

	bodyReader := f.FormatBody(responseData)

	responseBody, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, responseBody, body)
}

func TestXMLFormatterFormatBodyNilContent(t *testing.T) {

	responseData := ResponseData{
		Body:   nil,
		Status: http.StatusBadRequest,
	}

	f := XMLFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<message>%v</message>", http.StatusText(responseData.Status)), string(body))
}

func TestXMLFormatterFormatBodyBadContent(t *testing.T) {

	responseData := ResponseData{
		Body: func() {},
	}

	f := XMLFormatter{}

	testPanic := func() {
		f.FormatBody(responseData)
	}

	assert.Panics(t, testPanic, "did not panic on bad JSON marshalling input type")
}

func TestXMLFormatterFormatHeader(t *testing.T) {

	responseData := ResponseData{
		Body:   nil,
		Header: http.Header{},
	}

	responseData.Header.Set("key", "value")

	f := XMLFormatter{}

	header := f.FormatHeader(responseData)

	assert.Equal(t, "application/xml", header.Get("Content-Type"))
	assert.Equal(t, "value", header.Get("key"))
	assert.Equal(t, 2, len(header))
}

func TestXMLFormatterFormatStatus(t *testing.T) {

	responseData := ResponseData{
		Status: http.StatusBadRequest,
	}

	f := XMLFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusBadRequest, status)
}

func TestXMLFormatterFormatStatusZeroValue(t *testing.T) {

	responseData := ResponseData{}

	f := XMLFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusOK, status)
}
