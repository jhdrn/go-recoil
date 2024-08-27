package response

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// HTMLTemplateFormatter is a ResponseFormatter that formats responses using Go
// HTML templates.
type HTMLTemplateFormatter struct {
	Template *template.Template
}

func (f HTMLTemplateFormatter) FormatBody(responseData ResponseData) io.Reader {

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeWriter.Close()

		err := f.Template.Execute(pipeWriter, responseData.Content)
		if err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("failed to execute template: %w", err))
		}
	}()

	return pipeReader
}

func (f HTMLTemplateFormatter) FormatHeader(responseData ResponseData) http.Header {
	responseData.Header.Set("Content-Type", "text/html")
	return responseData.Header
}

func (f HTMLTemplateFormatter) FormatStatus(responseData ResponseData) int {
	if responseData.Status == 0 {
		return http.StatusOK
	}
	return responseData.Status
}
