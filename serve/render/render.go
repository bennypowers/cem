/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

// Package render provides in-process page rendering for static site generation.
// It is shared by both the dev server (serve) and the build command (serve --build).
package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
)

// PageResult holds the outcome of rendering a single page.
type PageResult struct {
	Route string
	Body  []byte
	Err   error
}

// Page renders a single URL path through the handler in-process.
// No HTTP server or TCP connection is involved. Panics in the handler
// are recovered and returned as errors.
func Page(handler http.Handler, urlPath string) (body []byte, err error) {
	req, reqErr := http.NewRequest("GET", urlPath, nil)
	if reqErr != nil {
		return nil, fmt.Errorf("creating request for %s: %w", urlPath, reqErr)
	}
	w := httptest.NewRecorder()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic rendering %s: %v", urlPath, r)
		}
	}()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d for %s", w.Code, urlPath)
	}
	return w.Body.Bytes(), nil
}

// Pages renders multiple URL paths concurrently through the handler.
// Concurrency is bounded by the concurrency parameter. Each page has
// independent panic recovery.
func Pages(handler http.Handler, routes []string, concurrency int) []PageResult {
	if concurrency < 1 {
		concurrency = 1
	}

	results := make([]PageResult, len(routes))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, route := range routes {
		wg.Add(1)
		go func(i int, route string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			body, err := Page(handler, route)
			results[i] = PageResult{Route: route, Body: body, Err: err}
		}(i, route)
	}

	wg.Wait()
	return results
}
