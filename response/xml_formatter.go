package response

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// XMLFormatter is a ResponseFormatter that formats responses as XML.
type XMLFormatter struct{}

// FormatBody formats the response body as XML. If the response body is nil,
// it will be set to a struct with a single field "message" and the value of
// http.StatusText(responseData.Status). If the response body is an error, it
// will be set to a struct with a single field "message" and the value of the
// error message. If the response body is an io.Reader, it will be returned as
// is. Otherwise, the response body will be marshaled to XML.
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
	} else if err, ok := responseData.Body.(error); ok {
		responseData.Body = struct {
			XMLName xml.Name `xml:"message"`
			Message string   `xml:",chardata"`
		}{
			Message: err.Error(),
		}
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
