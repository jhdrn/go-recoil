package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type TextFormatter struct{}

func (f TextFormatter) FormatBody(responseData ResponseData) io.Reader {

	if reader, ok := responseData.Content.(io.Reader); ok {
		return reader
	} else if err, ok := responseData.Content.(error); ok {
		responseData.Content = err.Error()
	}
	if b, ok := responseData.Content.([]byte); ok {
		return bytes.NewReader(b)
	}

	if str, ok := responseData.Content.(string); ok {
		return bytes.NewReader([]byte(str))
	}

	panic(fmt.Sprintf("unable to format response body of type %T", responseData.Content))
}

// FormatHeader formats the response header by setting the Content-Type to
// "application/json".
func (f TextFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "text/plain; charset=utf-8")
	return responseData.Header
}

// FormatStatus formats the response status. If the status is 0, it will be
// set to http.StatusOK.
func (f TextFormatter) FormatStatus(responseData ResponseData) int {
	if responseData.Status == 0 {
		return http.StatusOK
	}
	return responseData.Status
}
