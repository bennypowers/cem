/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package middleware

import (
	"net/http"
	"strings"
)

// Middleware is the standard middleware function signature.
// It wraps an http.Handler and returns a new http.Handler.
type Middleware func(http.Handler) http.Handler

// Chain applies middlewares in reverse order (last to first).
// This ensures the first middleware in the list wraps everything else.
//
// Example:
//
//	handler := middleware.Chain(
//	    mux,
//	    thirdMiddleware,  // Applied last (outermost)
//	    secondMiddleware,
//	    firstMiddleware,  // Applied first (innermost)
//	)
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	// Apply in reverse so first middleware wraps everything
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// ResponseRecorder captures HTTP response for modification by middleware.
// This is useful for middlewares that need to inspect or modify the response
// before sending it to the client (e.g., HTML injection, compression).
type ResponseRecorder struct {
	header        http.Header
	body          []byte
	statusCode    int
	headerWritten bool
}

// NewResponseRecorder creates a new response recorder
func NewResponseRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		header:     make(http.Header),
		body:       make([]byte, 0, 1024), // Pre-allocate 1KB
		statusCode: http.StatusOK,
	}
}

// Write implements http.ResponseWriter
func (r *ResponseRecorder) Write(b []byte) (int, error) {
	if !r.headerWritten {
		r.WriteHeader(http.StatusOK)
	}
	r.body = append(r.body, b...)
	return len(b), nil
}

// WriteHeader implements http.ResponseWriter
func (r *ResponseRecorder) WriteHeader(statusCode int) {
	if !r.headerWritten {
		r.statusCode = statusCode
		r.headerWritten = true
	}
}

// Header implements http.ResponseWriter
func (r *ResponseRecorder) Header() http.Header {
	return r.header
}

// Body returns the captured response body
func (r *ResponseRecorder) Body() []byte {
	return r.body
}

// StatusCode returns the captured status code
func (r *ResponseRecorder) StatusCode() int {
	return r.statusCode
}

// Ensure ResponseRecorder implements http.ResponseWriter
var _ http.ResponseWriter = (*ResponseRecorder)(nil)

// CopyHeaders copies all headers from src to dst, optionally excluding certain headers
func CopyHeaders(dst, src http.Header, exclude ...string) {
	excludeMap := make(map[string]bool)
	for _, key := range exclude {
		excludeMap[key] = true
	}

	for k, v := range src {
		if !excludeMap[k] {
			for _, vv := range v {
				dst.Add(k, vv)
			}
		}
	}
}

// IsHTMLResponse checks if a response is HTML based on content type
func IsHTMLResponse(contentType string) bool {
	if contentType == "" {
		return false
	}
	return strings.Contains(contentType, "text/html")
}
