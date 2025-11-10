package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type TextFormatter struct{}

// FormatBody converts different types of Content into an io.Reader.
// Supported types: nil, error, io.Reader, string, []byte, map[string]string.
func (f TextFormatter) FormatBody(responseData ResponseData) io.Reader {
	if responseData.Content == nil {
		responseData.Content = map[string]string{
			"message": http.StatusText(responseData.Status),
		}
	}

	switch v := responseData.Content.(type) {
	case io.Reader:
		return v
	case error:
		return bytes.NewReader([]byte(fmt.Sprintf(`{"message":"%s"}`, v.Error())))
	case []byte:
		return bytes.NewReader(v)
	case string:
		return bytes.NewReader([]byte(v))
	case map[string]string:
		var buf bytes.Buffer
		for k, val := range v {
			buf.WriteString(fmt.Sprintf("%s:%s\n", k, val))
		}
		return &buf
	default:
		panic(fmt.Sprintf("unable to format response body of type %T", responseData.Content))
	}
}

// FormatHeader sets default headers. If none exist, it creates a new Header map.
func (f TextFormatter) FormatHeader(responseData ResponseData) http.Header {
	if responseData.Header == nil {
		responseData.Header = http.Header{}
	}
	responseData.Header.Set("Content-Type", "text/plain; charset=utf-8")
	return responseData.Header
}

// FormatStatus returns the status code, defaulting to http.StatusOK if zero.
func (f TextFormatter) FormatStatus(responseData ResponseData) int {
	if responseData.Status == 0 {
		return http.StatusOK
	}
	return responseData.Status
}
