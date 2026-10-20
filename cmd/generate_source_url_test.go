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

// setConfigURL appends a top-level sourceControlRootUrl to the project's
// config file so the file-only path can be exercised on the same fixture.
func setConfigURL(t *testing.T, dir, url string) {
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

// setPackageConfigURL writes a per-package config setting only
// sourceControlRootUrl, so package-config precedence can be exercised on the
// temp fixture copy; packages without a config (card) keep none.
func setPackageConfigURL(t *testing.T, projectDir, pkg, url string) {
	t.Helper()
	cfgDir := filepath.Join(projectDir, "packages", pkg, ".config")
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		t.Fatalf("creating packages/%s config dir: %v", pkg, err)
	}
	cfgPath := filepath.Join(cfgDir, "cem.yaml")
	content := "sourceControlRootUrl: " + url + "\n"
	if err := os.WriteFile(cfgPath, []byte(content), 0o644); err != nil {
		t.Fatalf("writing packages/%s config: %v", pkg, err)
	}
}

// readPackageManifest reads a workspace package's generated manifest.
func readPackageManifest(t *testing.T, projectDir, pkg string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(projectDir, "packages", pkg, "custom-elements.json"))
	if err != nil {
		t.Fatalf("manifest should exist for packages/%s: %v", pkg, err)
	}
	return string(data)
}

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

	t.Run("FlagOnly", func(t *testing.T) {
		dir := setupTest(t, "source-href-flag")
		content := generate(t, dir, "--source-control-root-url", sourceURLFlag)

		testutil.CheckGolden(t, "SourceControlRootURLFlagOnly", content, testutil.GoldenOptions{
			Dir:         "testdata/goldens",
			Extension:   ".json",
			UseJSONDiff: true,
		})
	})

	t.Run("FlagMatchesConfigFile", func(t *testing.T) {
		flagDir := setupTest(t, "source-href-flag")
		flagOut := generate(t, flagDir, "--source-control-root-url", sourceURLFlag)

		cfgDir := setupTest(t, "source-href-flag")
		setConfigURL(t, cfgDir, sourceURLFlag)
		cfgOut := generate(t, cfgDir)

		// Inline, not golden: byte-identity between two runs is the property under test, un-golden-able.
		if !bytes.Equal(flagOut, cfgOut) {
			t.Errorf("flag path and config-file path produced different manifests\nflag:    %s\nconfig: %s", flagOut, cfgOut)
		}
	})

	t.Run("FlagOverridesConfigFile", func(t *testing.T) {
		dir := setupTest(t, "source-href-flag")
		setConfigURL(t, dir, sourceURLFile)

		content := string(generate(t, dir, "--source-control-root-url", sourceURLFlag))
		// Inline, not golden: only which URL wins is under test; expected output otherwise equals the FlagOnly golden.
		if strings.Contains(content, sourceURLFile) {
			t.Errorf("config-file value should be overridden by the flag, but the manifest contains it")
		}
		if !strings.Contains(content, sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in manifest source references")
		}
	})
}

// TestGenerateWorkspaceSourceControlRootURLFlag verifies workspace-mode
// precedence for sourceControlRootUrl: the URL points at the monorepo root,
// so it is a workspace-level setting. Effective value: the flag when passed
// (non-empty), else the root config's value, else each package's own config.
// A package value conflicting with the workspace value loses, with a warning.
func TestGenerateWorkspaceSourceControlRootURLFlag(t *testing.T) {
	t.Run("FlagOverridesPackageConfig", func(t *testing.T) {
		projectDir := setupTest(t, "workspace-generate")
		setPackageConfigURL(t, projectDir, "button", sourceURLFile)

		_, stderr := runCemCommand(t, projectDir, "generate", "--source-control-root-url", sourceURLFlag)

		// Inline, not golden: per-package substring checks across two packages, comparative property not a fixed manifest.
		button := readPackageManifest(t, projectDir, "button")
		if !strings.Contains(button, sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in packages/button source references:\n%s", button)
		}
		if strings.Contains(button, sourceURLFile) {
			t.Errorf("package config value should be overridden by the flag, but packages/button contains it:\n%s", button)
		}
		card := readPackageManifest(t, projectDir, "card")
		if !strings.Contains(card, sourceURLFlag) {
			t.Errorf("--source-control-root-url flag value not reflected in packages/card source references:\n%s", card)
		}
		if !strings.Contains(stderr, "conflicts with the workspace root URL") {
			t.Errorf("button's conflicting package config should produce a warning, but stderr has none:\n%s", stderr)
		}
	})

	t.Run("RootConfigCascadesToPackages", func(t *testing.T) {
		projectDir := setupTest(t, "workspace-generate")
		setConfigURL(t, projectDir, sourceURLFlag)

		runCemCommand(t, projectDir, "generate")

		// Inline, not golden: per-package substring checks across two packages, comparative property not a fixed manifest.
		for _, pkg := range []string{"button", "card"} {
			manifest := readPackageManifest(t, projectDir, pkg)
			if !strings.Contains(manifest, sourceURLFlag) {
				t.Errorf("root config value should cascade to packages/%s source references:\n%s", pkg, manifest)
			}
		}
	})

	t.Run("RootConfigOverridesConflictingPackageConfig", func(t *testing.T) {
		projectDir := setupTest(t, "workspace-generate")
		setPackageConfigURL(t, projectDir, "button", sourceURLFile)
		setConfigURL(t, projectDir, sourceURLFlag)

		_, stderr := runCemCommand(t, projectDir, "generate")

		// Inline, not golden: per-package substring checks across two packages, comparative property not a fixed manifest.
		button := readPackageManifest(t, projectDir, "button")
		if !strings.Contains(button, sourceURLFlag) {
			t.Errorf("root config value should win over packages/button's conflicting config:\n%s", button)
		}
		if strings.Contains(button, sourceURLFile) {
			t.Errorf("conflicting package config value should not reach packages/button's manifest:\n%s", button)
		}
		card := readPackageManifest(t, projectDir, "card")
		if !strings.Contains(card, sourceURLFlag) {
			t.Errorf("root config value should cascade to packages/card source references:\n%s", card)
		}
		if !strings.Contains(stderr, "conflicts with the workspace root URL") {
			t.Errorf("button's conflicting package config should produce a warning, but stderr has none:\n%s", stderr)
		}
	})

	t.Run("PackageConfigAppliesWithoutRootValue", func(t *testing.T) {
		projectDir := setupTest(t, "workspace-generate")
		setPackageConfigURL(t, projectDir, "button", sourceURLFile)

		runCemCommand(t, projectDir, "generate")

		// Inline, not golden: per-package substring checks across two packages, comparative property not a fixed manifest.
		button := readPackageManifest(t, projectDir, "button")
		if !strings.Contains(button, sourceURLFile) {
			t.Errorf("package config value should apply when no workspace-level value exists:\n%s", button)
		}
		card := readPackageManifest(t, projectDir, "card")
		if strings.Contains(card, sourceURLFile) {
			t.Errorf("packages/card has no config and no workspace value; its manifest should have no source references:\n%s", card)
		}
	})

	t.Run("FlagOverridesRootConfig", func(t *testing.T) {
		projectDir := setupTest(t, "workspace-generate")
		setConfigURL(t, projectDir, sourceURLFile)

		runCemCommand(t, projectDir, "generate", "--source-control-root-url", sourceURLFlag)

		// Inline, not golden: per-package substring checks across two packages, comparative property not a fixed manifest.
		for _, pkg := range []string{"button", "card"} {
			manifest := readPackageManifest(t, projectDir, pkg)
			if !strings.Contains(manifest, sourceURLFlag) {
				t.Errorf("flag value should override the root config value for packages/%s:\n%s", pkg, manifest)
			}
			if strings.Contains(manifest, sourceURLFile) {
				t.Errorf("root config value should be overridden by the flag, but packages/%s contains it:\n%s", pkg, manifest)
			}
		}
	})
}
