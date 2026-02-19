/*
Copyright 2026 Benny Powers. All rights reserved.
Use of this source code is governed by the GPLv3
license that can be found in the LICENSE file.
*/

package designtokens

// TODO: Delete this entire file once asimonim/load supports network fetching.
// This file only exists to support loadFromNetwork() in designtokens.go.
// When asimonim can fetch from CDN, load.Load() will handle everything and
// this lower-level parsing code becomes unnecessary.

import (
	"fmt"
	"strings"

	asimonimParser "bennypowers.dev/asimonim/parser"
	"bennypowers.dev/asimonim/resolver"
	"bennypowers.dev/asimonim/schema"
	"bennypowers.dev/asimonim/token"
)

// ParseOptions configures design token parsing.
// TODO: Remove when asimonim/load supports network fetching.
type ParseOptions struct {
	Prefix        string
	GroupMarkers  []string
	SchemaVersion schema.Version
}

// ParseTokensWithAsimonim parses design tokens from raw bytes.
// TODO: Remove when asimonim/load supports network fetching.
func ParseTokensWithAsimonim(data []byte, opts ParseOptions) (*token.Map, error) {
	if strings.HasPrefix(opts.Prefix, "-") {
		return nil, fmt.Errorf(
			"design token prefix %q should not start with dashes (use %q instead)",
			opts.Prefix,
			strings.TrimLeft(opts.Prefix, "-"))
	}

	parser := asimonimParser.NewJSONParser()

	// Detect schema version from content if not specified
	schemaVersion := opts.SchemaVersion
	if schemaVersion == schema.Unknown {
		detected, err := schema.DetectVersion(data, nil)
		if err != nil {
			return nil, fmt.Errorf("detecting schema version: %w", err)
		}
		schemaVersion = detected
	}

	parserOpts := asimonimParser.Options{
		Prefix:        opts.Prefix,
		GroupMarkers:  opts.GroupMarkers,
		SchemaVersion: schemaVersion,
		SkipSort:      true, // we don't need sorted output for map lookup
		SkipPositions: true, // cem doesn't need LSP position tracking
	}

	tokens, err := parser.Parse(data, parserOpts)
	if err != nil {
		return nil, fmt.Errorf("parsing tokens: %w", err)
	}

	// Resolve $extends relationships (v2025.10 feature, no-op for Draft schema)
	tokens, err = resolver.ResolveGroupExtensions(tokens, data)
	if err != nil {
		return nil, fmt.Errorf("resolving group extensions: %w", err)
	}

	// Resolve aliases so token references are expanded
	if err := resolver.ResolveAliases(tokens, schemaVersion); err != nil {
		return nil, fmt.Errorf("resolving aliases: %w", err)
	}

	return token.NewMap(tokens, opts.Prefix), nil
}
