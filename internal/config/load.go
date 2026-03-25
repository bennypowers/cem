package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// CemConfig holds the complete CEM configuration.
// This is the canonical type definition; cmd/config re-exports it for backward compatibility.
type CemConfig struct {
	ProjectDir   string `mapstructure:"projectDir" yaml:"projectDir" json:"projectDir"`
	ConfigFile   string `mapstructure:"configFile" yaml:"configFile" json:"configFile"`
	PackageName  string `mapstructure:"packageName" yaml:"packageName" json:"packageName"`
	Generate     GenerateConfig                `mapstructure:"generate" yaml:"generate" json:"generate"`
	MCP          MCPConfig                     `mapstructure:"mcp" yaml:"mcp" json:"mcp"`
	Health       HealthConfig                  `mapstructure:"health" yaml:"health" json:"health"`
	Serve        ServeConfig                   `mapstructure:"serve" yaml:"serve" json:"serve"`
	Export       map[string]FrameworkExportConfig `mapstructure:"export" yaml:"export" json:"export"`
	SourceControlRootUrl string `mapstructure:"sourceControlRootUrl" yaml:"sourceControlRootUrl" json:"sourceControlRootUrl"`
	Verbose            bool     `mapstructure:"verbose" yaml:"verbose" json:"verbose"`
	AdditionalPackages []string `mapstructure:"additionalPackages" yaml:"additionalPackages" json:"additionalPackages"`
}

type GenerateConfig struct {
	Files             []string           `mapstructure:"files" yaml:"files" json:"files"`
	Exclude           []string           `mapstructure:"exclude" yaml:"exclude" json:"exclude"`
	NoDefaultExcludes bool               `mapstructure:"noDefaultExcludes" yaml:"noDefaultExcludes" json:"noDefaultExcludes"`
	Output            string             `mapstructure:"output" yaml:"output" json:"output"`
	DesignTokens      DesignTokensConfig `mapstructure:"designTokens" yaml:"designTokens" json:"designTokens"`
	DemoDiscovery     DemoDiscoveryConfig `mapstructure:"demoDiscovery" yaml:"demoDiscovery" json:"demoDiscovery"`
}

type DemoDiscoveryConfig struct {
	FileGlob    string `mapstructure:"fileGlob" yaml:"fileGlob" json:"fileGlob"`
	URLPattern  string `mapstructure:"urlPattern" yaml:"urlPattern" json:"urlPattern"`
	URLTemplate string `mapstructure:"urlTemplate" yaml:"urlTemplate" json:"urlTemplate"`
}

type DesignTokensConfig struct {
	Spec   string `mapstructure:"spec" yaml:"spec" json:"spec"`
	Prefix string `mapstructure:"prefix" yaml:"prefix" json:"prefix"`
}

type MCPConfig struct {
	MaxDescriptionLength int `mapstructure:"maxDescriptionLength" yaml:"maxDescriptionLength" json:"maxDescriptionLength"`
}

type ServeConfig struct {
	Port        int              `mapstructure:"port" yaml:"port" json:"port"`
	OpenBrowser bool             `mapstructure:"openBrowser" yaml:"openBrowser" json:"openBrowser"`
	Transforms  TransformsConfig `mapstructure:"transforms" yaml:"transforms" json:"transforms"`
	URLRewrites []URLRewrite     `mapstructure:"urlRewrites" yaml:"urlRewrites" json:"urlRewrites"`
	Demos       DemosConfig      `mapstructure:"demos" yaml:"demos" json:"demos"`
}

type TransformsConfig struct {
	TypeScript TypeScriptTransformConfig `mapstructure:"typescript" yaml:"typescript" json:"typescript"`
	CSS        CSSTransformConfig        `mapstructure:"css" yaml:"css" json:"css"`
}

type TypeScriptTransformConfig struct {
	Enabled bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Target  string `mapstructure:"target" yaml:"target" json:"target"`
}

type CSSTransformConfig struct {
	Enabled bool     `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Include []string `mapstructure:"include" yaml:"include" json:"include"`
	Exclude []string `mapstructure:"exclude" yaml:"exclude" json:"exclude"`
}

type DemosConfig struct {
	Rendering string `mapstructure:"rendering" yaml:"rendering" json:"rendering"`
}

type URLRewrite struct {
	URLPattern  string `mapstructure:"urlPattern" yaml:"urlPattern" json:"urlPattern"`
	URLTemplate string `mapstructure:"urlTemplate" yaml:"urlTemplate" json:"urlTemplate"`
}

type FrameworkExportConfig struct {
	Output      string `mapstructure:"output" yaml:"output" json:"output"`
	StripPrefix string `mapstructure:"stripPrefix" yaml:"stripPrefix" json:"stripPrefix"`
	PackageName string `mapstructure:"packageName" yaml:"packageName" json:"packageName"`
	ModuleName  string `mapstructure:"moduleName" yaml:"moduleName" json:"moduleName"`
}

type HealthConfig struct {
	FailBelow int      `mapstructure:"failBelow" yaml:"failBelow" json:"failBelow"`
	Disable   []string `mapstructure:"disable" yaml:"disable" json:"disable"`
}

// LoadConfig reads and parses a CEM config file. The format is detected from
// the file extension. JSONC files have comments stripped before parsing.
func LoadConfig(path string) (*CemConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	format := FormatFromPath(path)
	cfg := &CemConfig{}

	switch format {
	case "yaml":
		if len(data) == 0 {
			return cfg, nil
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse yaml config %s: %w", path, err)
		}
	case "json":
		if len(data) == 0 {
			return cfg, nil
		}
		if err := json.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse json config %s: %w", path, err)
		}
	case "jsonc":
		stripped := bytes.TrimSpace(StripComments(data))
		if len(stripped) == 0 {
			return cfg, nil
		}
		if err := json.Unmarshal(stripped, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse jsonc config %s: %w", path, err)
		}
	default:
		return nil, fmt.Errorf("unsupported config format for %s", path)
	}

	return cfg, nil
}
