package response

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTMLTemplateFormatterFormatBody(t *testing.T) {

	template, err := template.New("test").Parse(
		`<doctype html><html><head><title>Test</title></head><body>{{.Key}}</body></html>`,
	)

	assert.NoError(t, err, "failed to parse template")

	type htmlData struct {
		Key string
	}

	responseData := ResponseData{
		Body: htmlData{
			Key: "value",
		},
	}

	f := HTMLTemplateFormatter{
		Template: template,
	}

	bodyReader := f.FormatBody(responseData)

	body, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, "<doctype html><html><head><title>Test</title></head><body>value</body></html>", string(body))
}

func TestHTMLTemplateFormatterFormatBodyPanicsOnBadTemplate(t *testing.T) {

	template, err := template.New("test").Parse(
		`<doctype html><html><head><title>Test</title></head><body>{{.Wrong}}</body></html>`,
	)

	assert.NoError(t, err, "failed to parse template")

	type htmlData struct {
		Key string
	}

	responseData := ResponseData{
		Body: htmlData{
			Key: "value",
		},
	}

	f := HTMLTemplateFormatter{
		Template: template,
	}

	bodyReader := f.FormatBody(responseData)
	_, err = io.ReadAll(bodyReader)

	assert.Error(t, err)

}

func TestHTMLTemplateFormatterFormatStream(t *testing.T) {

	body := []byte("<doctype html><html><head><title>Test</title></head><body>value</body></html>")

	template, err := template.New("test").Parse(
		`<doctype html><html><head><title>Test</title></head><body>value</body></html>`,
	)

	assert.NoError(t, err, "failed to parse template")

	responseData := ResponseData{
		Body: bytes.NewReader(body),
	}

	f := HTMLTemplateFormatter{
		Template: template,
	}

	bodyReader := f.FormatBody(responseData)

	responseBody, err := io.ReadAll(bodyReader)

	assert.NoError(t, err, "failed to read body reader")
	assert.Equal(t, responseBody, body)
}

func TestHTMLTemplateFormatterFormatHeader(t *testing.T) {

	responseData := ResponseData{
		Body:   struct{}{},
		Header: http.Header{},
	}

	responseData.Header.Set("key", "value")

	f := HTMLTemplateFormatter{}

	header := f.FormatHeader(responseData)

	assert.Equal(t, "text/html", header.Get("Content-Type"))
	assert.Equal(t, "value", header.Get("key"))
	assert.Equal(t, 2, len(header))
}

func TestHTMLTemplateFormatterFormatStatus(t *testing.T) {

	responseData := ResponseData{
		Status: http.StatusBadRequest,
	}

	f := HTMLTemplateFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusBadRequest, status)
}

func TestHTMLTemplateFormatterFormatStatusZeroValue(t *testing.T) {

	responseData := ResponseData{
		Body: map[string]any{},
	}

	f := HTMLTemplateFormatter{}

	status := f.FormatStatus(responseData)

	assert.Equal(t, http.StatusOK, status)
}
