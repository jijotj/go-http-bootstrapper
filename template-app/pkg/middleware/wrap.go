package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func Wrap(handler http.Handler, middleware ...Middleware) http.Handler {
	last := len(middleware) - 1
	for m := range middleware {
		handler = middleware[last-m](handler)
	}
	return handler
}
