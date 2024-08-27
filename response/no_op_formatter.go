package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// NoOpFormatter is a ResponseFormatter that does not format the response in any
// special way.
type NoOpFormatter struct{}

// FormatBody returns the body of the response data as a string.
func (f NoOpFormatter) FormatBody(responseData ResponseData) io.Reader {
	if reader, ok := responseData.Content.(io.Reader); ok {
		return reader
	}
	return bytes.NewReader([]byte(fmt.Sprintf("%v", responseData.Content)))
}

// FormatHeader returns the header of the response data.
func (f NoOpFormatter) FormatHeader(responseData ResponseData) http.Header {
	return responseData.Header
}

// FormatStatus returns the status of the response data.
func (f NoOpFormatter) FormatStatus(responseData ResponseData) int {
	return responseData.Status
}
