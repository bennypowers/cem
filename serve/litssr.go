/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"testing"

	litssr "bennypowers.dev/lit-ssr-wasm/go"
	"bennypowers.dev/cem/serve/middleware/routes"
)

// litSSRRenderer is the interface for rendering HTML with Lit SSR.
type litSSRRenderer interface {
	RenderHTML(ctx context.Context, html string) (string, error)
	Close(ctx context.Context) error
}

// initLitSSR initializes the Lit SSR renderer from pre-compiled
// QuickJS bytecode or source bundle embedded in the binary.
func (s *Server) initLitSSR(ctx context.Context) error {
	// Skip SSR in tests (no test exercises SSR, saves ~1s per server)
	if testing.Testing() {
		return nil
	}

	// Try bytecode first (fast path: ~400ms vs ~1s for source eval)
	if bytecode, err := routes.TemplatesFS.ReadFile("templates/ssr-bundle.qbc"); err == nil && len(bytecode) > 0 {
		renderer, err := litssr.NewFromBytecode(ctx, bytecode, 1)
		if err == nil {
			s.litSSR = renderer
			s.logger.Info("Lit SSR renderer initialized from bytecode")
			return nil
		}
		s.logger.Warning("Bytecode init failed, falling back to source: %v", err)
	}

	// Fallback to source eval
	source, err := routes.TemplatesFS.ReadFile("templates/ssr-bundle.js")
	if errors.Is(err, fs.ErrNotExist) || len(source) == 0 {
		return nil // No bundle available (e.g., test environment)
	}
	if err != nil {
		return fmt.Errorf("read SSR bundle: %w", err)
	}

	renderer, err := litssr.New(ctx, string(source), 1)
	if err != nil {
		return fmt.Errorf("init lit-ssr renderer: %w", err)
	}

	s.litSSR = renderer
	s.logger.Info("Lit SSR renderer initialized from source")
	return nil
}
