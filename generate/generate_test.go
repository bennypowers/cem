package generate

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
			path := fixture.Name()
			name := strings.Split(fixture.Name(), ".ts")[0]
			cases = append(cases, struct{name string; path string}{name, path})
		}
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Generate([]string{filepath.Join("test-fixtures", tc.path)}, []string{})
			golden := filepath.Join("test-golden", tc.name + ".json")
			if *update {
				os.WriteFile(golden, []byte(actual), 0644)
			}
			expected, _ := os.ReadFile(golden)
			if string(expected) != actual {
				t.Fail()
			}
		})
	}
}
