package generate

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"github.com/nsf/jsondiff"
	"github.com/spf13/viper"
)

var update = flag.Bool("update", false, "update golden files")

type testcase struct {
	name   string
	path   string
	config *config.CemConfig
}

func TestGenerate(t *testing.T) {
	// Accept a flexible pattern for fixture filtering
	fixturePattern := os.Getenv("FIXTURE_PATTERN")
	var re *regexp.Regexp
	var err error
	if fixturePattern != "" {
		re, err = regexp.Compile(fixturePattern)
		if err != nil {
			t.Fatalf("failed to compile FIXTURE_PATTERN %q: %v", fixturePattern, err)
		}
	}

	projects, err := os.ReadDir(filepath.Join("../test/fixtures"))
	if err != nil {
		log.Fatal(err)
	}

	for _, projectEntry := range projects {
		if !projectEntry.IsDir() {
			continue
		}
		projectDir := filepath.Join("../test/fixtures", projectEntry.Name())
		t.Run(projectEntry.Name(), func(t *testing.T) {
			oldWd, _ := os.Getwd()
			if err := os.Chdir(projectDir); err != nil {
				t.Fatalf("failed to chdir to %s: %v", projectDir, err)
			}
			defer os.Chdir(oldWd)

			projectGoldenDir := "golden"
			if err := os.MkdirAll(projectGoldenDir, 0755); err != nil {
				t.Fatalf("failed to create %s: %v", projectGoldenDir, err)
			}

			configPath := filepath.Join(".config", "cem.yaml")
			cfg, err := config.LoadConfig(viper.New(), configPath, ".")
			if err != nil {
				t.Fatalf("failed to load config: %v", err)
			}
			fixtures, err := os.ReadDir("src")
			if err != nil {
				t.Fatalf("cannot read src directory: %v", err)
			}

			var cases []testcase
			for _, fixture := range fixtures {
				if !fixture.Type().IsRegular() || filepath.Ext(fixture.Name()) != ".ts" {
					continue
				}
				name := strings.TrimSuffix(fixture.Name(), ".ts")
				if re != nil && !re.MatchString(name) {
					continue
				}
				cases = append(cases, testcase{
					name:   name,
					path:   filepath.Join("src", fixture.Name()),
					config: cfg,
				})
			}

			for _, tc := range cases {
				tc := tc // capture range variable
				t.Run(tc.name, func(t *testing.T) {
					tc.config.Generate.Files = []string{tc.path}
					actual, err := Generate(tc.config)
					if err != nil {
						t.Fatal(err)
					}
					golden := filepath.Join(projectGoldenDir, tc.name+".json")
					if *update {
						if err := os.WriteFile(golden, []byte(*actual), 0644); err != nil {
							t.Fatalf("failed to write golden file: %v", err)
						}
					}
					expected, err := os.ReadFile(golden)
					if err != nil {
						t.Fatalf("golden file missing: %s (have you run with -update?)\nerror: %v", golden, err)
					}
					// Validate JSON
					var jsExpected, jsActual any
					if err := json.Unmarshal(expected, &jsExpected); err != nil {
						t.Fatalf("expected golden file is invalid JSON: %v", err)
					}
					if err := json.Unmarshal([]byte(*actual), &jsActual); err != nil {
						t.Fatalf("actual output is invalid JSON: %v\noutput:\n%s", err, *actual)
					}
					if string(expected) != *actual {
						options := jsondiff.DefaultConsoleOptions()
						t.Error(jsondiff.Compare(expected, []byte(*actual), &options))
					}
				})
			}
		})
	}
}
