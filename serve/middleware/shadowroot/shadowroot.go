/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package shadowroot

import (
	"context"
	"maps"
	"net/http"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/types"
)

// SSRRenderer renders HTML with Declarative Shadow DOM injection.
type SSRRenderer interface {
	RenderHTML(ctx context.Context, html string) (string, error)
}

// New creates a shadow root injection middleware that passes HTML
// responses through a Lit SSR renderer to inject Declarative Shadow DOM.
func New(logger types.Logger, renderer SSRRenderer) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip internal routes
			if strings.HasPrefix(r.URL.Path, "/__cem") {
				next.ServeHTTP(w, r)
				return
			}

			if renderer == nil {
				next.ServeHTTP(w, r)
				return
			}

			// Capture response
			rec := middleware.NewResponseRecorder()
			next.ServeHTTP(rec, r)

			// Only process HTML responses
			contentType := rec.Header().Get("Content-Type")
			if !middleware.IsHTMLResponse(contentType) {
				maps.Copy(w.Header(), rec.Header())
				w.WriteHeader(rec.StatusCode())
				_, _ = w.Write(rec.Body())
				return
			}

			// Pass HTML through Lit SSR renderer
			processed, err := renderer.RenderHTML(r.Context(), string(rec.Body()))
			if err != nil {
				logger.Warning("Lit SSR rendering failed, serving without DSD: %v", err)
				maps.Copy(w.Header(), rec.Header())
				w.WriteHeader(rec.StatusCode())
				_, _ = w.Write(rec.Body())
				return
			}

			maps.Copy(w.Header(), rec.Header())
			w.WriteHeader(rec.StatusCode())
			_, _ = w.Write([]byte(processed))
		})
	}
}
