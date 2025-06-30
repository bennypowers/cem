package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

type DemoDiscoveryConfig struct {
	FileGlob  string               `mapstructure:"fileGlob"`
	URLPattern   string            `mapstructure:"urlPattern"`
	URLTemplate  string            `mapstructure:"urlTemplate"`
}

type DesignTokensConfig struct {
	// Path or `npm:@scope/package/path/to/file.json` spec to DTCG format design
	// tokens json module
	Spec string
	// Prefix those design tokens use in CSS. If the design tokens are generated
	// by style dictionary and have a `name` field, that will be used instead.
	Prefix string
}

// CLI or config arguments passed to the generate command
type GenerateConfig struct {
	// List of files or file globs to include in the manifest
	Files              []string
	// List of files or file globs to exclude from the manifest
	Exclude            []string
	// Do not exclude files that are excluded by default e.g. *.d.ts files.
	NoDefaultExcludes  bool
	// File path to write output to. If omitted, output will be written to stdout.
	Output             string
	// Configuration for design tokens discovery
	DesignTokens       DesignTokensConfig
	// Configuration for demo file discovery
	DemoDiscovery      DemoDiscoveryConfig
}

type CemConfig struct {
	projectDir           string           `mapstructure:"-"`
	configFile           string           `mapstructure:"-"`
	// Package name, as would appear in a package.json "name" field
	PackageName          string           `mapstructure:"packageName"`
	// Generate command options
	Generate             GenerateConfig   `mapstructure:"generate"`
	// Canonical public source control URL corresponding to project root on primary branch.
	// e.g. https://github.com/bennypowers/cem/tree/main/
	SourceControlRootUrl string           `mapstructure:"sourceControlRootUrl"`
	// Verbose logging output
	Verbose              bool             `mapstructure:"verbose"`
}

func (cfg *CemConfig) GetProjectDir() string { return cfg.projectDir }
func (cfg *CemConfig) GetConfigFile() string { return cfg.configFile }

func resolveProjectDir(configPath, projectDirFlag string) string {
	if projectDirFlag != "" {
		abs, err := expandPath(projectDirFlag)
		if err != nil {
			pterm.Fatal.Printf("Invalid --project-dir: %v", err)
		}
		return abs
	}
	configAbs, err := filepath.Abs(configPath)
	if err != nil {
		pterm.Fatal.Printf("Invalid --config: %v", err)
	}
	configDir := filepath.Dir(configAbs)
	base := filepath.Base(configDir)
	if base == ".config" || base == "config" {
		return filepath.Dir(configDir)
	}
	// fallback: use current working directory
	cwd, err := os.Getwd()
	if err != nil {
		pterm.Fatal.Printf("Unable to get current working directory: %v", err)
	}
	if !strings.HasPrefix(configAbs, cwd) {
		pterm.Warning.Printf("Warning: --config is outside of current dir, guessing project root as %s\n", cwd)
	}
	return cwd
}

// expandPath expands ~, handles relative and absolute paths
func expandPath(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		// Support ~/ and ~
		if path == "~" {
			path = home
		} else if strings.HasPrefix(path, "~/") {
			path = filepath.Join(home, path[2:])
		}
		// Note: ~user/ is not supported (Go stdlib doesn't provide this)
	}
	return filepath.Abs(path)
}

func LoadConfig(cfgFile string, projectDir string) (config *CemConfig, err error) {
	v := viper.New()
	if cfgFile != "" {
		// Use config file from the flag.
		cfgFile, err = expandPath(cfgFile)
		if err != nil {
			return nil, err
		}
		v.SetConfigFile(cfgFile)
	} else {
		// Search config in local project .config directory with name "cem.yaml"
		cfgFile, err = expandPath(filepath.Join(projectDir, "./.config", "cem.yaml"))
		if err != nil {
			return nil, err
		}
		v.AddConfigPath(filepath.Dir(cfgFile))
		v.SetConfigType("yaml")
		v.SetConfigName("cem.yaml")
	}
	v.AutomaticEnv() // read in environment variables that match
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	if projectDir != "" {
		projectDir, err = expandPath(projectDir)
		if err != nil {
			return nil, err
		}
	}
	config.projectDir = resolveProjectDir(cfgFile, projectDir)
	config.configFile = cfgFile
	if !filepath.IsAbs(config.Generate.Output) {
		config.Generate.Output = filepath.Join(projectDir, config.Generate.Output)
	}
	if config.Verbose {
		pterm.EnableDebugMessages()
	} else {
		pterm.DisableDebugMessages()
	}
	return config, nil
}
