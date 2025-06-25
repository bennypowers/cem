package config

type DemoDiscoveryConfig struct {
	FilePattern  string            `mapstructure:"filePattern"`
	// Canonical public source control URL corresponding to project root on primary branch.
	// e.g. https://github.com/bennypowers/cem/tree/main/
	SourceControlUrl string        `mapstructure:"sourceControlUrl"`
	URLPattern   string            `mapstructure:"urlPattern"`
	URLTemplate  string            `mapstructure:"urlTemplate"`
}

// CLI or config arguments passed to the generate command
type GenerateConfig struct {
	// List of files or file globs to include in the manifest
	Files              []string
	// List of files or file globs to exclude from the manifest
	Exclude            []string
	// Path or `npm:@scope/package/path/to/file.json` spec to DTCG format design
	// tokens json module
	DesignTokensSpec   string
	// Prefix those design tokens use in CSS. If the design tokens are generated
	// by style dictionary and have a `name` field, that will be used instead.
	DesignTokensPrefix string
	// Do not exclude files that are excluded by default e.g. *.d.ts files.
	NoDefaultExcludes  bool
	// File path to write output to. If omitted, output will be written to stdout.
	Output             string
	DemoDiscovery      DemoDiscoveryConfig
}

type CemConfig struct {
	Generate GenerateConfig `mapstructure:"generate"`
	TagPrefix string `mapstructure:"tagPrefix"`
}
