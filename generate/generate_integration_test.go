package generate

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"github.com/spf13/viper"
)

var projectDirFlag = flag.String("projectdir", "", "Path to the project directory to run integration tests against")

func TestIntegration_GenerateOnRealProject(t *testing.T) {
	flag.Parse()
	if *projectDirFlag == "" {
		t.Skip("Skipping integration test; -projectdir not provided")
	}

	projectDir, err := filepath.Abs(*projectDirFlag)
	t.Log(projectDir)
	if err != nil {
		t.Fatalf("could not resolve absolute path for %q: %v", *projectDirFlag, err)
	}

	oldWd, _ := os.Getwd()
	if err := os.Chdir(projectDir); err != nil {
		t.Fatalf("failed to chdir to %s: %v", projectDir, err)
	}
	defer os.Chdir(oldWd)

	var cfg config.CemConfig
	viper.SetConfigFile(filepath.Join(".config", "cem.yaml"))
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Unable to read config: %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		t.Fatal("Unable to decode config:", err)
	}
	cemConfig, err := LoadConfig([]string{}, cfg)
	if err != nil {
		t.Fatalf("Could not load config: %v", err)
	}

	if _, err := Generate(&cemConfig); err != nil {
		t.Fatalf("Generator failed: %v", err)
	}
}
