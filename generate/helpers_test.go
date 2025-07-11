package generate_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/manifest"
	"gopkg.in/yaml.v3"
)

func loadConfig(t testing.TB, ctx manifest.WorkspaceContext) (*config.CemConfig, error) {
	var config config.CemConfig
	config.ProjectDir = ctx.Root()
	cfgFile := ctx.ConfigFile()
	rc, err := ctx.ReadFile(cfgFile)
	if err != nil {
		t.Log(err)
	} else {
		defer rc.Close()
		if err := yaml.NewDecoder(rc).Decode(&config); err != nil {
			return nil, err
		}
		config.ProjectDir = ctx.Root()
		config.ConfigFile = cfgFile
		// Make output path project-root-relative if needed
		if config.Generate.Output != "" && !filepath.IsAbs(config.Generate.Output) {
			config.Generate.Output = filepath.Join(config.ProjectDir, config.Generate.Output)
		}
	}
	t.Logf("%+v", config)
	if config.Generate.Exclude == nil {
		config.Generate.Exclude = make([]string, 0)
	}
	return &config, nil
}
