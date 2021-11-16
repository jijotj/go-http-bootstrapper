package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"{{.ServiceName}}/pkg/middleware"
)

func TestWrap(t *testing.T) {
	wasHandled := false
	names := []string(nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { wasHandled = true })
	m1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			names = append(names, "m1")
			next.ServeHTTP(w, r)
		})
	}
	m2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			names = append(names, "m2")
			next.ServeHTTP(w, r)
		})
	}
	m3 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			names = append(names, "m3")
			next.ServeHTTP(w, r)
		})
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	middleware.Wrap(handler, m1, m2, m3).ServeHTTP(w, r)

	assert.True(t, wasHandled, "Handler was called")
	assert.Equal(t, []string{"m1", "m2", "m3"}, names, "Incorrect middleware calls")
}

func TestWrapWhenEmpty(t *testing.T) {
	wasHandled := false

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { wasHandled = true })

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	middleware.Wrap(handler).ServeHTTP(w, r)

	assert.True(t, wasHandled, "Handler was called")
}
