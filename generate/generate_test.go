/*
Copyright © 2025 Benny Powers

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

	DT "bennypowers.dev/cem/designtokens"
	"bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/types"
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

			// First create a basic context to check if design tokens are configured
			baseCtx := W.NewFileSystemWorkspaceContextWithDefaults(projectDir)
			if err := baseCtx.Init(); err != nil {
				t.Fatalf("TestGenerate: %v", err)
			}

			cfg, err := baseCtx.Config()
			if err != nil {
				t.Fatalf("TestGenerate: %v", err)
			}

			// If design tokens are configured, create context with design tokens loader
			var ctx types.WorkspaceContext
			if cfg.Generate.DesignTokens.Spec != "" {
				designTokensLoader := DT.NewLoader()
				ctx = W.NewFileSystemWorkspaceContext(projectDir, designTokensLoader)
				if err := ctx.Init(); err != nil {
					t.Fatalf("TestGenerate: %v", err)
				}
			} else {
				ctx = baseCtx
			}

			var cases []testcase
			if err := filepath.WalkDir(filepath.Join(projectDir, "src"), func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() && strings.HasSuffix(d.Name(), ".ts") {
					name := strings.TrimSuffix(d.Name(), ".ts")
					if re != nil && !re.MatchString(name) {
						return nil
					}
					rel, err := filepath.Rel(filepath.Join(projectDir, "src"), path)
					if err != nil {
						return err
					}
					cases = append(cases, testcase{
						name: name,
						path: filepath.Join("src", rel),
					})
				}
				return nil
			}); err != nil {
				t.Fatalf("Error walking directory: %v", err)
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
						diff, str := jsondiff.Compare(expected, []byte(*actual), &options)
						if diff == jsondiff.FullMatch {
							t.Logf("Semantic match, string mismatch: %s", str)
						} else {
							t.Error(diff, str)
						}
					}
				})
			}
		})
	}
}
