package middleware

import "net/http"

type MiddlewareFunc func(http.Handler) http.Handler

// NewMiddlewares return registered middlewares as a singgle http.Handler
func NewMiddlewares() MiddlewareFunc {
	return registerMiddlewares(
		CORSMiddleware,
		ThrottleMiddleware,
		TimeoutMiddleware,
	)
}

func registerMiddlewares(middlewares ...func(h http.Handler) http.Handler) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for _, m := range middlewares {
			h = m(h)
		}
		return h
	}
}
