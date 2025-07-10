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
	"path/filepath"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"

	"bennypowers.dev/cem/manifest"
)

type DemoDiscoveryConfig struct {
	FileGlob    string `mapstructure:"fileGlob" yaml:"fileGlob"`
	URLPattern  string `mapstructure:"urlPattern" yaml:"urlPattern"`
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

type CemConfig struct {
	ProjectDir string `mapstructure:"projectDir" yaml:"projectDir"`
	ConfigFile string `mapstructure:"configFile" yaml:"configFile"`
	// Package name, as would appear in a package.json "name" field
	PackageName string `mapstructure:"packageName" yaml:"packageName"`
	// Generate command options
	Generate GenerateConfig `mapstructure:"generate" yaml:"generate"`
	// Canonical public source control URL corresponding to project root on primary branch.
	// e.g. https://github.com/bennypowers/cem/tree/main/
	SourceControlRootUrl string `mapstructure:"sourceControlRootUrl" yaml:"sourceControlRootUrl"`
	// Verbose logging output
	Verbose bool `mapstructure:"verbose" yaml:"verbose"`
}

// LoadConfig reads CemConfig from the project context's config file.
func LoadConfig(ctx manifest.ProjectContext) (*CemConfig, error) {
	cfgFile, err := ctx.ConfigFile()
	if err != nil {
		return nil, err
	}
	rc, err := ctx.ReadFile(cfgFile)
	if err != nil {
		// If config file not found, return default config
		return &CemConfig{
			ProjectDir: ctx.Root(),
			ConfigFile: cfgFile,
		}, nil
	}
	defer rc.Close()

	var config CemConfig
	if err := yaml.NewDecoder(rc).Decode(&config); err != nil {
		return nil, err
	}

	config.ProjectDir = ctx.Root()
	config.ConfigFile = cfgFile

	// Make output path project-root-relative if needed
	if config.Generate.Output != "" && !filepath.IsAbs(config.Generate.Output) {
		config.Generate.Output = filepath.Join(config.ProjectDir, config.Generate.Output)
	}

	// Set debug verbosity
	if config.Verbose {
		pterm.EnableDebugMessages()
	} else {
		pterm.DisableDebugMessages()
	}

	return &config, nil
}
