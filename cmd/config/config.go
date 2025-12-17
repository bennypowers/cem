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
package config

import (
	"bennypowers.dev/cem/serve/middleware/types"
)

type DemoDiscoveryConfig struct {
	FileGlob string `mapstructure:"fileGlob" yaml:"fileGlob"`
	// URLPattern uses standard URLPattern syntax (e.g., "/components/:element/demo/:demo.html")
	URLPattern string `mapstructure:"urlPattern" yaml:"urlPattern"`
	// URLTemplate defines how to generate URLs from captured parameters
	// Uses {{.param}} syntax to interpolate URLPattern parameters
	URLTemplate string `mapstructure:"urlTemplate" yaml:"urlTemplate"`
}

type DesignTokensConfig struct {
	// Path or `npm:@scope/package/path/to/file.json` spec to DTCG format design tokens json module
	Spec string `mapstructure:"spec" yaml:"spec"`
	// Prefix those design tokens use in CSS. If the design tokens are generated
	// by style dictionary and have a `name` field, that will be used instead.
	Prefix string `mapstructure:"prefix" yaml:"prefix"`
}

// CLI or config arguments passed to the generate command
type GenerateConfig struct {
	// List of files or file globs to include in the manifest
	Files []string `mapstructure:"files" yaml:"files"`
	// List of files or file globs to exclude from the manifest
	Exclude []string `mapstructure:"exclude" yaml:"exclude"`
	// Do not exclude files that are excluded by default e.g. *.d.ts files.
	NoDefaultExcludes bool `mapstructure:"noDefaultExcludes" yaml:"noDefaultExcludes"`
	// File path to write output to. If omitted, output will be written to stdout.
	Output string `mapstructure:"output" yaml:"output"`
	// Configuration for design tokens discovery
	DesignTokens DesignTokensConfig `mapstructure:"designTokens" yaml:"designTokens"`
	// Configuration for demo file discovery
	DemoDiscovery DemoDiscoveryConfig `mapstructure:"demoDiscovery" yaml:"demoDiscovery"`
}

type MCPConfig struct {
	// Maximum length for description fields before truncation (default: 2000)
	MaxDescriptionLength int `mapstructure:"maxDescriptionLength" yaml:"maxDescriptionLength"`
}

type ServeConfig struct {
	// Port to run the development server on (default: 8000)
	Port int `mapstructure:"port" yaml:"port"`
	// Whether to automatically open browser on server start
	OpenBrowser bool `mapstructure:"openBrowser" yaml:"openBrowser"`
	// Import map configuration
	ImportMap types.ImportMapConfig `mapstructure:"importMap" yaml:"importMap"`
	// Transform configuration
	Transforms TransformsConfig `mapstructure:"transforms" yaml:"transforms"`
	// Path mappings for src/dist separation (e.g., {"/dist/": "./src/"})
	PathMappings map[string]string `mapstructure:"pathMappings" yaml:"pathMappings"`
}

type TransformsConfig struct {
	TypeScript TypeScriptTransformConfig `mapstructure:"typescript" yaml:"typescript"`
	CSS        CSSTransformConfig        `mapstructure:"css" yaml:"css"`
}

type TypeScriptTransformConfig struct {
	// Enable TypeScript transformation (default: true)
	Enabled bool `mapstructure:"enabled" yaml:"enabled"`
	// Transform target (e.g., "es2022", "es2020") - overridden by --target flag
	Target string `mapstructure:"target" yaml:"target"`
}

type CSSTransformConfig struct {
	// Enable CSS transformation (default: true)
	Enabled bool `mapstructure:"enabled" yaml:"enabled"`
	// Glob patterns for CSS files to include (default: all .css files)
	Include []string `mapstructure:"include" yaml:"include"`
	// Glob patterns for CSS files to exclude
	Exclude []string `mapstructure:"exclude" yaml:"exclude"`
}

type CemConfig struct {
	ProjectDir string `mapstructure:"projectDir" yaml:"projectDir"`
	ConfigFile string `mapstructure:"configFile" yaml:"configFile"`
	// Package name, as would appear in a package.json "name" field
	PackageName string `mapstructure:"packageName" yaml:"packageName"`
	// Generate command options
	Generate GenerateConfig `mapstructure:"generate" yaml:"generate"`
	// MCP server options
	MCP MCPConfig `mapstructure:"mcp" yaml:"mcp"`
	// Serve command options
	Serve ServeConfig `mapstructure:"serve" yaml:"serve"`
	// Canonical public source control URL corresponding to project root on primary branch.
	// e.g. https://github.com/bennypowers/cem/tree/main/
	SourceControlRootUrl string `mapstructure:"sourceControlRootUrl" yaml:"sourceControlRootUrl"`
	// Verbose logging output
	Verbose bool `mapstructure:"verbose" yaml:"verbose"`
}

func (c *CemConfig) Clone() *CemConfig {
	if c == nil {
		return nil
	}
	clone := *c
	// Deep copy slices
	if c.Generate.Files != nil {
		clone.Generate.Files = make([]string, len(c.Generate.Files))
		copy(clone.Generate.Files, c.Generate.Files)
	}
	if c.Generate.Exclude != nil {
		clone.Generate.Exclude = make([]string, len(c.Generate.Exclude))
		copy(clone.Generate.Exclude, c.Generate.Exclude)
	}
	// Deep copy import map config override
	if c.Serve.ImportMap.Override.Imports != nil {
		clone.Serve.ImportMap.Override.Imports = make(map[string]string, len(c.Serve.ImportMap.Override.Imports))
		for k, v := range c.Serve.ImportMap.Override.Imports {
			clone.Serve.ImportMap.Override.Imports[k] = v
		}
	}
	if c.Serve.ImportMap.Override.Scopes != nil {
		clone.Serve.ImportMap.Override.Scopes = make(map[string]map[string]string, len(c.Serve.ImportMap.Override.Scopes))
		for scopeKey, scopeMap := range c.Serve.ImportMap.Override.Scopes {
			clone.Serve.ImportMap.Override.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
			for k, v := range scopeMap {
				clone.Serve.ImportMap.Override.Scopes[scopeKey][k] = v
			}
		}
	}
	// Deep copy path mappings
	if c.Serve.PathMappings != nil {
		clone.Serve.PathMappings = make(map[string]string, len(c.Serve.PathMappings))
		for k, v := range c.Serve.PathMappings {
			clone.Serve.PathMappings[k] = v
		}
	}
	return &clone
}
