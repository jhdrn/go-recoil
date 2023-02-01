package response

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type JSONFormatter struct{}

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

func (f JSONFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "application/json")
	return responseData.Header
}

type XMLFormatter struct{}

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

func (f XMLFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "application/xml")
	return responseData.Header
}
