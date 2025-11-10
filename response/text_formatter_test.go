package response

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

func TestTextFormatterFormatBody(t *testing.T) {
	tests := []struct {
		name        string
		content     any
		wantContent string
		shouldPanic bool
	}{
		{
			name:        "io.Reader content is returned directly",
			content:     bytes.NewReader([]byte("hello")),
			wantContent: "hello",
		},
		{
			name:        "error content is converted to string",
			content:     errors.New("something went wrong"),
			wantContent: "something went wrong",
		},
		{
			name:        "[]byte content is returned correctly",
			content:     []byte("byte data"),
			wantContent: "byte data",
		},
		{
			name:        "string content is returned correctly",
			content:     "plain text",
			wantContent: "plain text",
		},
		{
			name:        "unsupported content type panics",
			content:     12345,
			shouldPanic: true,
		},
	}

	f := TextFormatter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ResponseData{Content: tt.content}

			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not occur")
					}
				}()
				f.FormatBody(r)
				return
			}

			reader := f.FormatBody(r)
			got, err := io.ReadAll(reader)
			if err != nil {
				t.Fatalf("failed to read result: %v", err)
			}
			if string(got) != tt.wantContent {
				t.Errorf("got %q, want %q", string(got), tt.wantContent)
			}
		})
	}
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
