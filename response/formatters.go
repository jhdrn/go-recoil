package response

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// NoOpFormatter is a ResponseFormatter that does not format the response in any
// special way.
type NoOpFormatter struct{}

// FormatBody returns the body of the response data as a string.
func (f NoOpFormatter) FormatBody(responseData ResponseData) io.Reader {
	if reader, ok := responseData.Body.(io.Reader); ok {
		return reader
	}
	return bytes.NewReader([]byte(fmt.Sprintf("%v", responseData.Body)))
}

// FormatHeader returns the header of the response data.
func (f NoOpFormatter) FormatHeader(responseData ResponseData) http.Header {
	return responseData.Header
}

// FormatStatus returns the status of the response data.
func (f NoOpFormatter) FormatStatus(responseData ResponseData) int {
	return responseData.Status
}

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

// XMLFormatter is a ResponseFormatter that formats responses as XML.
type XMLFormatter struct{}

// FormatBody formats the response body as XML. If the response body is nil,
// it will be set to a struct with a single field "message" and the value of
// http.StatusText(responseData.Status).
func (f XMLFormatter) FormatBody(responseData ResponseData) io.Reader {

	if responseData.Body == nil {
		responseData.Body = struct {
			XMLName xml.Name `xml:"message"`
			Message string   `xml:",chardata"`
		}{
			Message: http.StatusText(responseData.Status),
		}
	} else if reader, ok := responseData.Body.(io.Reader); ok {
		return reader
	}

	xmlBytes, err := xml.Marshal(responseData.Body)
	if err != nil {
		panic(fmt.Errorf("failed to marshal XML data: %w", err))
	}

	return bytes.NewReader(append([]byte(xml.Header), xmlBytes...))
}

// FormatHeader formats the response header by setting the Content-Type to
// "application/xml".
func (f XMLFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "application/xml")
	return responseData.Header
}

// FormatStatus formats the response status. If the status is 0, it will be
// set to http.StatusOK.
func (f XMLFormatter) FormatStatus(responseData ResponseData) int {
	if responseData.Status == 0 {
		return http.StatusOK
	}
	return responseData.Status
}

// TODO: TemplateFormatter ?
