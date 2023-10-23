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
// http.StatusText(responseData.Status).
func (f JSONFormatter) FormatBody(responseData ResponseData) io.Reader {

	if responseData.Body == nil {
		responseData.Body = map[string]string{
			"message": http.StatusText(responseData.Status),
		}
	} else if msg, ok := responseData.Body.(string); ok {
		responseData.Body = map[string]string{
			"message": msg,
		}
	} else if reader, ok := responseData.Body.(io.Reader); ok {
		return reader
	}

	jsonBytes, err := json.Marshal(responseData.Body)
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
