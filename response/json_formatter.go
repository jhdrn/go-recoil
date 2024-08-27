package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSONFormatter is a ResponseFormatter that formats responses as JSON.
type JSONFormatter struct{}

// FormatBody formats the response body as JSON. If the response body is nil,
// it will be set to a map with a single key "message" and the value of
// http.StatusText(responseData.Status). If the response body is an error,
// the response body will be set to a map with a single key "message" and the
// value of the error message. If the response body is an io.Reader, it will be
// returned as is. Otherwise, the response body will be marshaled to JSON.
func (f JSONFormatter) FormatBody(responseData ResponseData) io.Reader {

	if responseData.Content == nil {
		responseData.Content = map[string]string{
			"message": http.StatusText(responseData.Status),
		}
	} else if reader, ok := responseData.Content.(io.Reader); ok {
		return reader
	} else if err, ok := responseData.Content.(error); ok {
		responseData.Content = map[string]string{
			"message": err.Error(),
		}
	}

	jsonBytes, err := json.Marshal(responseData.Content)
	if err != nil {
		panic(fmt.Errorf("failed to marshal JSON data: %w", err))
	}

	return bytes.NewReader(jsonBytes)
}

// FormatHeader formats the response header by setting the Content-Type to
// "application/json".
func (f JSONFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "application/json")
	return responseData.Header
}

// FormatStatus formats the response status. If the status is 0, it will be
// set to http.StatusOK.
func (f JSONFormatter) FormatStatus(responseData ResponseData) int {
	if responseData.Status == 0 {
		return http.StatusOK
	}
	return responseData.Status
}
