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
