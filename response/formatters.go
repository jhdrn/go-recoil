package response

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// JSONFormatter is a ResponseFormatter that formats responses as JSON.
type JSONFormatter struct{}

// FormatBody formats the response body as JSON. If the response body is nil,
// it will be set to a map with a single key "message" and the value of
// http.StatusText(responseData.Status).
func (f JSONFormatter) FormatBody(responseData ResponseData) []byte {

	if responseData.Content == nil {
		responseData.Content = map[string]string{
			"message": http.StatusText(responseData.Status),
		}
	}

	jsonBytes, err := json.Marshal(responseData.Content)
	if err != nil {
		panic(fmt.Errorf("failed to marshal JSON data: %w", err))
	}

	return jsonBytes
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
func (f XMLFormatter) FormatBody(responseData ResponseData) []byte {

	if responseData.Content == nil {
		responseData.Content = struct {
			XMLName xml.Name `xml:"message"`
			Message string   `xml:",chardata"`
		}{
			Message: http.StatusText(responseData.Status),
		}
	}

	xmlBytes, err := xml.Marshal(responseData.Content)
	if err != nil {
		panic(fmt.Errorf("failed to marshal XML data: %w", err))
	}

	return append([]byte(xml.Header), xmlBytes...)
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
