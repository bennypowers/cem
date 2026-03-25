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
	"fmt"
	"maps"

	IC "bennypowers.dev/cem/internal/config"
)

// Type aliases re-exported from internal/config for backward compatibility.
type CemConfig = IC.CemConfig
type GenerateConfig = IC.GenerateConfig
type DemoDiscoveryConfig = IC.DemoDiscoveryConfig
type DesignTokensConfig = IC.DesignTokensConfig
type MCPConfig = IC.MCPConfig
type ServeConfig = IC.ServeConfig
type TransformsConfig = IC.TransformsConfig
type TypeScriptTransformConfig = IC.TypeScriptTransformConfig
type CSSTransformConfig = IC.CSSTransformConfig
type DemosConfig = IC.DemosConfig
type URLRewrite = IC.URLRewrite
type FrameworkExportConfig = IC.FrameworkExportConfig
type HealthConfig = IC.HealthConfig

// Validate validates the configuration and returns an error if invalid
func Validate(c *CemConfig) error {
	if c == nil {
		return nil
	}

	// Validate demo rendering mode
	rendering := c.Serve.Demos.Rendering
	if rendering != "" {
		switch rendering {
		case "light", "shadow":
			// Valid modes
		case "iframe":
			return fmt.Errorf("serve.demos.rendering: 'iframe' mode is not yet implemented - use 'light' or 'shadow'")
		default:
			return fmt.Errorf("serve.demos.rendering: invalid value '%s' - must be 'light', 'shadow', or 'iframe'", rendering)
		}
	}

	return nil
}

func Clone(c *CemConfig) *CemConfig {
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
		maps.Copy(clone.Serve.ImportMap.Override.Imports, c.Serve.ImportMap.Override.Imports)
	}
	if c.Serve.ImportMap.Override.Scopes != nil {
		clone.Serve.ImportMap.Override.Scopes = make(map[string]map[string]string, len(c.Serve.ImportMap.Override.Scopes))
		for scopeKey, scopeMap := range c.Serve.ImportMap.Override.Scopes {
			clone.Serve.ImportMap.Override.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
			maps.Copy(clone.Serve.ImportMap.Override.Scopes[scopeKey], scopeMap)
		}
	}
	// Deep copy URL rewrites
	if c.Serve.URLRewrites != nil {
		clone.Serve.URLRewrites = make([]URLRewrite, len(c.Serve.URLRewrites))
		copy(clone.Serve.URLRewrites, c.Serve.URLRewrites)
	}
	// Deep copy additional packages
	if c.AdditionalPackages != nil {
		clone.AdditionalPackages = make([]string, len(c.AdditionalPackages))
		copy(clone.AdditionalPackages, c.AdditionalPackages)
	}
	// Deep copy export config
	if c.Export != nil {
		clone.Export = make(map[string]FrameworkExportConfig, len(c.Export))
		maps.Copy(clone.Export, c.Export)
	}
	return &clone
}
