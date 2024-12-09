package middleware

import "net/http"

func WithJWT(h http.HandlerFunc) http.HandlerFunc {
	return h
}
