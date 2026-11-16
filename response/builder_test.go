package response

import (
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseBuilderWithCookie(t *testing.T) {
	r := NewBuilder(WithConfig(
		Config{
			Formatter: JSONFormatter{},
		},
	))

	c := http.Cookie{
		Name:  "foo",
		Value: "bar",
	}

	assert.Equal(t, c.String(), r.OK().WithCookie(&c).Header().Get("Set-Cookie"))
}

func TestResponseBuilderWithHeader(t *testing.T) {
	r := NewBuilder(WithConfig(
		Config{
			Formatter: NoOpFormatter{},
		},
	))

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
	r := NewBuilder(WithConfig(
		Config{
			Formatter: NoOpFormatter{},
		},
	))

	h := http.Header{
		"foo": []string{"bar"},
	}

	assert.Equal(t, h, r.OK().WithHeaderEntry("foo", "bar").Header())
}

func TestResponseBuilderWithHeaderEntryDoesNotMutateReceiver(t *testing.T) {
	base := NewBuilder(WithConfig(Config{Formatter: NoOpFormatter{}})).
		WithHeaderEntry("X-App", "1")

	derived := base.WithHeaderEntry("X-User", "alice")

	assert.Equal(t, http.Header{"X-App": {"1"}}, base.Header())
	assert.Equal(t, http.Header{"X-App": {"1"}, "X-User": {"alice"}}, derived.Header())
}

func TestResponseBuilderWithHeaderEntryCopiesValues(t *testing.T) {
	values := []string{"a"}
	r := NewBuilder(WithConfig(Config{Formatter: NoOpFormatter{}})).
		WithHeaderEntry("X-Foo", values...)

	values[0] = "changed"

	assert.Equal(t, "a", r.Header().Get("X-Foo"))
}

func TestResponseBuilderWithHeaderCopiesInput(t *testing.T) {
	h := http.Header{"X-Foo": {"bar"}}
	r := NewBuilder(WithConfig(Config{Formatter: NoOpFormatter{}})).WithHeader(h)

	h.Set("X-Foo", "changed")
	h.Set("X-Other", "leak")

	assert.Equal(t, http.Header{"X-Foo": {"bar"}}, r.Header())
}

func TestResponseBuilderWithHeaderNilThenEntry(t *testing.T) {
	r := NewBuilder(WithConfig(Config{Formatter: NoOpFormatter{}}))

	assert.NotPanics(t, func() {
		r = r.WithHeader(nil).WithHeaderEntry("X-Foo", "bar")
	})
	assert.Equal(t, "bar", r.Header().Get("X-Foo"))
}

func TestResponseBuilderWithMultipleCookies(t *testing.T) {
	base := NewBuilder()
	a := http.Cookie{Name: "a", Value: "1"}
	b := http.Cookie{Name: "b", Value: "2"}

	r := base.WithCookie(&a).WithCookie(&b)

	assert.Equal(t, []string{a.String(), b.String()}, r.Header().Values("Set-Cookie"))
	assert.Empty(t, base.Header().Values("Set-Cookie"))
}

func TestResponseBuilderHeaderDoesNotMutateBuilder(t *testing.T) {
	r := NewBuilder(WithConfig(Config{Formatter: JSONFormatter{}}))

	h := r.Header()
	h.Set("X-Leak", "1")

	assert.Equal(t, "application/json", h.Get("Content-Type"))
	assert.Empty(t, r.responseData.Header)
	assert.Empty(t, r.Header().Get("X-Leak"))
}

// Run with -race: a shared base builder must be safe to derive from and
// format concurrently.
func TestResponseBuilderConcurrentUseOfSharedBase(t *testing.T) {
	base := NewBuilder().WithHeaderEntry("X-App", "1")

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := base.OK().
				WithHeaderEntry("X-User", "alice").
				WithCookie(&http.Cookie{Name: "a", Value: "1"})
			_ = r.Header()
		}()
	}
	wg.Wait()

	assert.Equal(t, http.Header{"X-App": {"1"}}, base.responseData.Header)
}
