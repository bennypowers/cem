package generate

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	"github.com/nsf/jsondiff"
)

var update = flag.Bool("update", false, "update golden files")

func TestGenerate(t *testing.T) {
	fixtures, err := os.ReadDir(filepath.Join("test-fixtures"))
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct{name string; path string}{}

	for _, fixture := range fixtures {
		if fixture.Type().IsRegular() {
			pathToFile := fixture.Name()
			ext := filepath.Ext(pathToFile)
			if ext == ".ts" {
				name := strings.Split(fixture.Name(), ".ts")[0]
				cases = append(cases, struct{name string; path string}{name, pathToFile})
			}
		}
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := C.CemConfig{
				Generate: C.GenerateConfig{
					Files: []string{filepath.Join("test-fixtures", tc.path)},
					DesignTokensSpec: "./" + filepath.Join("test-fixtures", "design-tokens.json"),
					DesignTokensPrefix: "token",
				},
			}
			actual, err := Generate(&cfg)
			if err != nil {
				t.Error(err)
			}
			golden := filepath.Join("test-golden", tc.name + ".json")
			if *update {
				os.WriteFile(golden, []byte(*actual), 0644)
			}
			expected, _ := os.ReadFile(golden)
			if string(expected) != *actual {
				options := jsondiff.DefaultConsoleOptions()
				t.Error(jsondiff.Compare(expected, []byte(*actual), &options))
			}
		})
	}
}
