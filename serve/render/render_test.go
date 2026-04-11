/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package render

import (
	"net/http"
	"strings"
	"testing"
)

func TestPage_Success(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("hello"))
	})

	body, err := Page(handler, "/test")
	if err != nil {
		t.Fatalf("Page() error: %v", err)
	}
	if string(body) != "hello" {
		t.Errorf("Page() = %q, want %q", string(body), "hello")
	}
}

func TestPage_NonOKStatus(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := Page(handler, "/fail")
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
	if !strings.Contains(err.Error(), "HTTP 500") {
		t.Errorf("expected HTTP 500 error, got: %v", err)
	}
}

func TestPage_RecoversPanic(t *testing.T) {
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("renderer exploded")
	})

	_, err := Page(handler, "/boom")
	if err == nil {
		t.Fatal("expected error from recovered panic")
	}
	if !strings.Contains(err.Error(), "panic rendering /boom") {
		t.Errorf("expected panic error, got: %v", err)
	}
}

func TestPages_Concurrent(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("page:" + r.URL.Path))
	})

	routes := []string{"/a", "/b", "/c", "/d"}
	results := Pages(handler, routes, 2)

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}

	// Results should be in the same order as input routes
	for i, res := range results {
		if res.Route != routes[i] {
			t.Errorf("result[%d].Route = %q, want %q", i, res.Route, routes[i])
		}
		if res.Err != nil {
			t.Errorf("result[%d] unexpected error: %v", i, res.Err)
		}
		want := "page:" + routes[i]
		if string(res.Body) != want {
			t.Errorf("result[%d].Body = %q, want %q", i, string(res.Body), want)
		}
	}
}

func TestPages_PanicIsolation(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bad" {
			panic("nil pointer")
		}
		_, _ = w.Write([]byte("ok"))
	})

	routes := []string{"/good", "/bad", "/also-good"}
	results := Pages(handler, routes, 2)

	// /good and /also-good should succeed
	if results[0].Err != nil {
		t.Errorf("/good should succeed, got: %v", results[0].Err)
	}
	if results[2].Err != nil {
		t.Errorf("/also-good should succeed, got: %v", results[2].Err)
	}

	// /bad should have a panic error
	if results[1].Err == nil {
		t.Fatal("/bad should have a panic error")
	}
	if !strings.Contains(results[1].Err.Error(), "panic") {
		t.Errorf("expected panic error for /bad, got: %v", results[1].Err)
	}
}
