package middlewares

import (
	"errors"
	"net/http"

	"sample/api/auth"
	"sample/api/exitcode"
	"sample/api/responses"
)

// SetMiddlewareJSON is...
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			// Handle preflight request
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			res.Header().Set("Access-Control-Max-Age", "86400")
			res.WriteHeader(http.StatusNoContent)
			return
		}
		res.Header().Set("Content-Type", "application/json")

		next(res, req)
	}
}

// SetMiddlewareAuth is...
func SetMiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := auth.TokenValid(req)

		if err != nil {
			responses.ERROR(res, http.StatusUnauthorized, exitcode.UNAUTHORIZED, errors.New("Unauthorized"))

			return
		}

		next(res, req)
	}
}
