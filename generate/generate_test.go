package generate_test

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"bennypowers.dev/cem/generate"
	W "bennypowers.dev/cem/workspace"
	"github.com/nsf/jsondiff"
)

var update = flag.Bool("update", false, "update golden files")

type testcase struct {
	name string
	path string
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

	projects, err := os.ReadDir("test/fixtures")
	if err != nil {
		log.Fatal(err)
	}

	for _, projectEntry := range projects {
		if !projectEntry.IsDir() || projectEntry.Name() == ".config" {
			continue
		}
		projectDir := filepath.Join("test/fixtures/", projectEntry.Name())

		t.Run(projectEntry.Name(), func(t *testing.T) {
			projectGoldenDir := "golden"
			if err := os.MkdirAll(filepath.Join(projectDir, projectGoldenDir), 0755); err != nil {
				t.Fatalf("failed to create %s: %v", projectGoldenDir, err)
			}

			ctx := W.NewFileSystemWorkspaceContext(projectDir)
			if err := ctx.Init(); err != nil {
				t.Fatalf("TestGenerate: %v", err)
			}

			cfg, err := ctx.Config()
			if err != nil {
				t.Fatalf("TestGenerate: %v", err)
			}
			fixtures, err := os.ReadDir(filepath.Join(projectDir, "src"))
			if err != nil {
				t.Fatalf("TestGenerate: %v", err)
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
					name: name,
					path: filepath.Join("src", fixture.Name()),
				})
			}

			for _, tc := range cases {
				// capture range variable
				t.Run(tc.name, func(t *testing.T) {
					cfg.Generate.Files = []string{tc.path}
					actual, err := generate.Generate(ctx)
					if err != nil {
						t.Fatal(err)
					}
					golden := filepath.Join(projectDir, projectGoldenDir, tc.name+".json")
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
