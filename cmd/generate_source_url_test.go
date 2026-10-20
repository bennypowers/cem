//go:build e2e

/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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

package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

const (
	// sourceURLFlag is passed via the CLI flag.
	sourceURLFlag = "https://github.com/example/repo/tree/main/"
	// sourceURLFile is a different URL set only in the config file.
	sourceURLFile = "https://github.com/example/different/tree/v1.0.0/"
)

// TestGenerateSourceControlRootURLFlag pins issue #449: the
// --source-control-root-url persistent flag must produce the same source
// references as the equivalent config-file setting, and must override the
// config file when both are given.
func TestGenerateSourceControlRootURLFlag(t *testing.T) {
	generate := func(t *testing.T, dir string, extraArgs ...string) []byte {
		t.Helper()
		args := append([]string{"generate", "src/element.ts", "--output", "out.json"}, extraArgs...)
		_, stderr := runCemCommand(t, dir, args...)
		content, err := os.ReadFile(filepath.Join(dir, "out.json"))
		if err != nil {
			t.Fatalf("output file not created: %v\nstderr: %s", err, stderr)
		}
		return content
	}

	// setConfigURL appends a top-level sourceControlRootUrl to the fixture's
	// config file so the file-only path can be exercised on the same fixture.
	setConfigURL := func(t *testing.T, dir, url string) {
		t.Helper()
		cfgPath := filepath.Join(dir, ".config", "cem.yaml")
		cfg, err := os.ReadFile(cfgPath)
		if err != nil {
			t.Fatalf("reading fixture config: %v", err)
		}
		cfg = append(cfg, []byte("sourceControlRootUrl: "+url+"\n")...)
		if err := os.WriteFile(cfgPath, cfg, 0o644); err != nil {
			t.Fatalf("writing fixture config: %v", err)
		}
	}

	t.Run("FlagOnly", func(t *testing.T) {
		dir := setupTest(t, "source-href-flag")
		content := generate(t, dir, "--source-control-root-url", sourceURLFlag)

		testutil.CheckGolden(t, "SourceControlRootURLFlagOnly", content, testutil.GoldenOptions{
			Dir:         "testdata/goldens",
			Extension:   ".json",
			UseJSONDiff: true,
		})

		if !strings.Contains(string(content), sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in manifest source references")
		}
	})

	t.Run("FlagMatchesConfigFile", func(t *testing.T) {
		flagDir := setupTest(t, "source-href-flag")
		flagOut := generate(t, flagDir, "--source-control-root-url", sourceURLFlag)

		cfgDir := setupTest(t, "source-href-flag")
		setConfigURL(t, cfgDir, sourceURLFlag)
		cfgOut := generate(t, cfgDir)

		if !bytes.Equal(flagOut, cfgOut) {
			t.Errorf("flag path and config-file path produced different manifests\nflag:    %s\nconfig: %s", flagOut, cfgOut)
		}
	})

	t.Run("FlagOverridesConfigFile", func(t *testing.T) {
		dir := setupTest(t, "source-href-flag")
		setConfigURL(t, dir, sourceURLFile)

		content := string(generate(t, dir, "--source-control-root-url", sourceURLFlag))
		if strings.Contains(content, sourceURLFile) {
			t.Errorf("config-file value should be overridden by the flag, but the manifest contains it")
		}
		if !strings.Contains(content, sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in manifest source references")
		}
	})
}

// TestGenerateWorkspaceSourceControlRootURLFlag verifies the flag also reaches
// every package in workspace mode, where each package loads its own config.
func TestGenerateWorkspaceSourceControlRootURLFlag(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	runCemCommand(t, projectDir, "generate", "--source-control-root-url", sourceURLFlag)

	for _, pkg := range []string{"button", "card"} {
		manifestPath := filepath.Join(projectDir, "packages", pkg, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		if err != nil {
			t.Fatalf("manifest should exist for packages/%s: %v", pkg, err)
		}
		if !strings.Contains(string(data), sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in packages/%s source references:\n%s", pkg, data)
		}
	}
}
