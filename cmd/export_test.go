//go:build e2e

/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestExportReactIntegration(t *testing.T) {
	// Resolve the project root (cmd/ is one level down from root)
	projectRoot, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("failed to resolve project root: %v", err)
	}

	fixtureDir := filepath.Join(projectRoot, "export", "testdata", "integration", "react")
	generatedDir := filepath.Join(fixtureDir, "generated")
	kitchenSinkDir := filepath.Join(projectRoot, "examples", "kitchen-sink")

	// Clean up generated directory before and after
	os.RemoveAll(generatedDir)
	t.Cleanup(func() { os.RemoveAll(generatedDir) })

	// Step 1: Run cem export to generate React wrappers
	t.Log("Running cem export...")
	stdout, stderr := runCemCommand(t, projectRoot,
		"export",
		"--format", "react",
		"-o", generatedDir,
		"-p", kitchenSinkDir,
	)
	t.Logf("cem export stdout: %s", stdout)
	t.Logf("cem export stderr: %s", stderr)

	// Verify generated directory has files
	entries, err := os.ReadDir(generatedDir)
	if err != nil {
		t.Fatalf("generated directory not created: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("no files generated in output directory")
	}

	// Step 2: npm ci
	t.Log("Running npm ci...")
	npmCI := exec.Command("npm", "ci")
	npmCI.Dir = fixtureDir
	npmCI.Stdout = os.Stdout
	npmCI.Stderr = os.Stderr
	if err := npmCI.Run(); err != nil {
		t.Fatalf("npm ci failed: %v", err)
	}

	// Step 3: vitest runtime tests (also validates TypeScript via Vite's compiler)
	t.Log("Running vitest...")
	vitest := exec.Command("npx", "vitest", "run")
	vitest.Dir = fixtureDir
	out, err := vitest.CombinedOutput()
	if err != nil {
		t.Fatalf("vitest run failed:\n%s\n%v", string(out), err)
	}
	t.Logf("vitest output: %s", string(out))
}
