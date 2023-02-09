package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseBuilderWithCookie(t *testing.T) {
	r := NewResponseBuilder(Config{
		Formatter: JSONFormatter{},
	})

	c := http.Cookie{
		Name:  "foo",
		Value: "bar",
	}

	assert.Equal(t, c.String(), r.OK().WithCookie(&c).Header().Get("Set-Cookie"))
}

func TestResponseBuilderWithHeader(t *testing.T) {
	r := NewResponseBuilder(Config{
		Formatter: NoOpFormatter{},
	})

	h := http.Header{
		"foo": []string{"bar"},
		"bar": []string{"foo"},
	}

	formattedHeader := r.OK().WithHeader(h).Header()
	assert.Equal(t, 2, len(formattedHeader))

	fooValue := formattedHeader["foo"]
	assert.Equal(t, "bar", fooValue[0])
	assert.Equal(t, 1, len(fooValue))

	barValue := formattedHeader["bar"]
	assert.Equal(t, "foo", barValue[0])
	assert.Equal(t, 1, len(barValue))

}

func TestResponseBuilderWithHeaderEntrySingleValue(t *testing.T) {
	r := NewResponseBuilder(Config{
		Formatter: NoOpFormatter{},
	})

	h := http.Header{
		"foo": []string{"bar"},
	}

	assert.Equal(t, h, r.OK().WithHeaderEntry("foo", "bar").Header())
}
