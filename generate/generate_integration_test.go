package generate

import (
	"flag"
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

	// Assuming you have some function to load and parse the config:

	var cfg config.CemConfig
	viper.SetConfigFile(filepath.Join(*projectDirFlag, ".config/cem.yaml"))
	_ = viper.ReadInConfig()
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
