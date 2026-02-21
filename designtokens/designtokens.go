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
	"fmt"
	"strings"

	"bennypowers.dev/asimonim/load"
	"bennypowers.dev/asimonim/specifier"
	"bennypowers.dev/asimonim/token"
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
	if err := validatePrefix(prefix); err != nil {
		return nil, err
	}

	spec := cfg.Generate.DesignTokens.Spec
	opts := buildLoadOptions(spec, ctx.Root(), prefix)

	tokens, err := load.Load(context.Background(), spec, opts)
	if err != nil {
		return nil, err
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

// buildLoadOptions constructs load.Options for the given specifier,
// wiring up CDN fallback for package specifiers.
func buildLoadOptions(spec, root, prefix string) load.Options {
	opts := load.Options{
		Root:   root,
		Prefix: prefix,
	}
	parsed := specifier.Parse(spec)
	if parsed.IsNPM() || parsed.IsJSR() {
		opts.Fetcher = load.NewHTTPFetcher(load.DefaultMaxSize)
		opts.CDN = specifier.CDNEsmSh
	}
	return opts
}

// validatePrefix returns an error if prefix starts with dashes.
func validatePrefix(prefix string) error {
	if strings.HasPrefix(prefix, "-") {
		return fmt.Errorf(
			"design token prefix %q should not start with dashes (use %q instead)",
			prefix,
			strings.TrimLeft(prefix, "-"))
	}
	return nil
}
