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
package designtokens

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"bennypowers.dev/asimonim/load"
	"bennypowers.dev/asimonim/specifier"
	"bennypowers.dev/asimonim/token"
	"bennypowers.dev/cem/internal/logging"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
)

// DesignTokens provides access to design tokens by name.
type DesignTokens struct {
	tokens *token.Map
}

// Get returns the Token for the given name, prepending the prefix if it's not present.
func (dt *DesignTokens) Get(name string) (*token.Token, bool) {
	return dt.tokens.Get(name)
}

// LoadDesignTokens loads tokens from a path or Deno-style specifier and returns a DesignTokens struct.
// The prefix is prepended to all token names on load.
// Configuration is loaded from:
// 1. CEM config (generate.designTokens.prefix, generate.designTokens.spec) - highest priority
// 2. Asimonim config (.config/design-tokens.{yaml,json}) - fallback for prefix/groupMarkers
func LoadDesignTokens(ctx types.WorkspaceContext) (*DesignTokens, error) {
	cfg, err := ctx.Config()
	if err != nil {
		return nil, err
	}

	prefix := cfg.Generate.DesignTokens.Prefix
	if strings.HasPrefix(prefix, "-") {
		return nil, fmt.Errorf(
			"design token prefix %q should not start with dashes (use %q instead)",
			prefix,
			strings.TrimLeft(prefix, "-"))
	}

	spec := cfg.Generate.DesignTokens.Spec
	opts := load.Options{
		Root:   ctx.Root(),
		Prefix: prefix,
	}

	// Try loading via asimonim (handles local files and node_modules)
	tokens, err := load.Load(spec, opts)
	if err != nil {
		// Only fall back to network for resolution/not-found errors
		// Other errors (parsing, alias resolution, etc.) should be reported directly
		if isResolutionError(err) {
			logging.Debug("local resolution failed for %q, attempting network fallback: %v", spec, err)
			tokens, err = loadFromNetwork(spec, opts)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &DesignTokens{tokens: tokens}, nil
}

// Loader implements the types.DesignTokensLoader interface
type Loader struct{}

// NewLoader creates a new design tokens loader
func NewLoader() types.DesignTokensLoader {
	return &Loader{}
}

// Load implements types.DesignTokensLoader.Load
func (l *Loader) Load(ctx types.WorkspaceContext) (types.DesignTokens, error) {
	// Cast the context to our minimal types.WorkspaceContext interface
	// This is safe because the workspace package will pass the correct type
	return LoadDesignTokens(ctx)
}

func MergeDesignTokensToModule(module *M.Module, designTokens types.DesignTokens) {
	for i, d := range module.Declarations {
		if d, ok := d.(*M.CustomElementDeclaration); ok {
			for i, p := range d.CustomElement.CssProperties {
				if tok, ok := designTokens.Get(p.Name); ok {
					// Merge user's description with design token description
					// If user has a description, concatenate with two newlines
					// If user has no description, use only the design token description
					if tok.Description != "" {
						if p.Description != "" {
							p.Description = p.Description + "\n\n" + tok.Description
						} else {
							p.Description = tok.Description
						}
					}
					p.Syntax = tok.CSSSyntax()
					// Set default value from design token if not already set
					if p.Default == "" {
						p.Default = tok.DisplayValue()
					}
					// Merge deprecation status from design token
					if tok.Deprecated && p.Deprecated == nil {
						if tok.DeprecationMessage != "" {
							p.Deprecated = M.DeprecatedReason(tok.DeprecationMessage)
						} else {
							p.Deprecated = M.DeprecatedFlag(true)
						}
					}
					d.CustomElement.CssProperties[i] = p
				}
			}
		}
		module.Declarations[i] = d
	}
}

const (
	// networkTimeout is the maximum time to wait for a network fetch
	networkTimeout = 30 * time.Second
	// maxResponseSize is the maximum allowed response size (10 MB)
	maxResponseSize = 10 * 1024 * 1024
)

// isResolutionError returns true if the error indicates a resolution or not-found issue
// that can be recovered by falling back to network fetch.
func isResolutionError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	// Check for resolution errors from asimonim (wraps with "failed to resolve specifier")
	// or file not found errors (from filesystem operations)
	return strings.Contains(errStr, "failed to resolve specifier") ||
		strings.Contains(errStr, "no such file or directory") ||
		strings.Contains(errStr, "file does not exist") ||
		strings.Contains(errStr, "cannot find module")
}

// loadFromNetwork fetches tokens from CDN when local resolution fails.
// TODO: Remove once asimonim supports network fetching.
func loadFromNetwork(spec string, opts load.Options) (*token.Map, error) {
	parsed := specifier.Parse(spec)
	if !parsed.IsNPM() && !parsed.IsJSR() {
		return nil, errors.New("network fallback only supported for npm:/jsr: specifiers")
	}

	// Fetch from unpkg.com with timeout
	url := "https://unpkg.com/" + parsed.Package + "/" + parsed.File

	ctx, cancel := context.WithTimeout(context.Background(), networkTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("timeout fetching %s after %s", url, networkTimeout)
		}
		return nil, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch from CDN: %s returned %s", url, resp.Status)
	}

	// Limit response size to prevent excessive memory use
	limitedReader := io.LimitReader(resp.Body, maxResponseSize+1)
	content, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from %s: %w", url, err)
	}
	if len(content) > maxResponseSize {
		return nil, fmt.Errorf("response from %s exceeds maximum size of %d bytes", url, maxResponseSize)
	}

	// Parse fetched content using lower-level API
	return ParseTokensWithAsimonim(content, ParseOptions{
		Prefix: opts.Prefix,
	})
}
