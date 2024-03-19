package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestXMLFormatterFormatBodyErrorContent(t *testing.T) {

	responseData := ResponseData{
		Body:   fmt.Errorf("error message"),
		Status: http.StatusBadRequest,
	}

	f := XMLFormatter{}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<message>error message</message>", string(body))
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
