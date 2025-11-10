package response

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestTextFormatterFormatBody(t *testing.T) {
	formatter := TextFormatter{}

	tests := []struct {
		name     string
		content  interface{}
		status   int
		expected string
	}{
		{"nil content", nil, http.StatusNotFound, "Not Found"},
		{"string content", "hello world", 0, "hello world"},
		{"byte slice content", []byte("byte content"), 0, "byte content"},
		{"io.Reader content", bytes.NewReader([]byte("reader content")), 0, "reader content"},
		{"error content", io.ErrUnexpectedEOF, 0, `{"message":"unexpected EOF"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respData := ResponseData{
				Content: tt.content,
				Status:  tt.status,
				Header:  nil,
			}

			reader := formatter.FormatBody(respData)
			b, err := io.ReadAll(reader)
			if err != nil {
				t.Fatalf("Failed to read body: %v", err)
			}
			got := string(b)
			if !strings.Contains(got, tt.expected) {
				t.Errorf("expected %q to contain %q", got, tt.expected)
			}
		})
	}
}

func TestTextFormatterFormatBodyPanic(t *testing.T) {
	formatter := TextFormatter{}
	respData := ResponseData{
		Content: 12345,
		Status:  0,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for unsupported type, but did not panic")
		}
	}()

	formatter.FormatBody(respData)
}

func TestTextFormatterFormatHeader(t *testing.T) {
	f := TextFormatter{}
	r := ResponseData{Header: http.Header{}}

	h := f.FormatHeader(r)
	if got := h.Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Errorf("unexpected Content-Type: got %q", got)
	}
}

func TestTextFormatterFormatStatus(t *testing.T) {
	f := TextFormatter{}

	tests := []struct {
		name   string
		status int
		want   int
	}{
		{"default 0 maps to 200 OK", 0, http.StatusOK},
		{"custom status is preserved", http.StatusNotFound, http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ResponseData{Status: tt.status}
			if got := f.FormatStatus(r); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
