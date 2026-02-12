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
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func assertGeneratedFiles(t *testing.T, generatedDir string) {
	t.Helper()
	entries, err := os.ReadDir(generatedDir)
	if err != nil {
		t.Fatalf("generated directory not created: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("no files generated in output directory")
	}
}

func runNpm(t *testing.T, dir string, args ...string) {
	t.Helper()
	t.Logf("Running npm %s...", args[0])
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)
	cmd := exec.CommandContext(ctx, "npm", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("npm %s failed:\n%s\n%v", args[0], string(out), err)
	}
	t.Logf("npm %s output: %s", args[0], string(out))
}

func runVitest(t *testing.T, dir string) {
	t.Helper()
	t.Log("Running vitest...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)
	cmd := exec.CommandContext(ctx, "npx", "vitest", "run")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("vitest run failed:\n%s\n%v", string(out), err)
	}
	t.Logf("vitest output: %s", string(out))
}

func runExportIntegrationTest(t *testing.T, framework string) {
	t.Helper()
	projectRoot, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("failed to resolve project root: %v", err)
	}

	kitchenSinkDir := filepath.Join(projectRoot, "examples", "kitchen-sink")

	t.Run("library-author", func(t *testing.T) {
		fixtureDir := filepath.Join(projectRoot, "export", "testdata", "integration", framework)
		generatedDir := filepath.Join(fixtureDir, "generated")

		os.RemoveAll(generatedDir)
		t.Cleanup(func() { os.RemoveAll(generatedDir) })

		t.Log("Running cem export...")
		stdout, stderr := runCemCommand(t, projectRoot,
			"export",
			"--format", framework,
			"-o", generatedDir,
			"-p", kitchenSinkDir,
		)
		t.Logf("cem export stdout: %s", stdout)
		t.Logf("cem export stderr: %s", stderr)

		assertGeneratedFiles(t, generatedDir)
		runNpm(t, fixtureDir, "ci")
		runVitest(t, fixtureDir)
	})

	t.Run("consumer", func(t *testing.T) {
		fixtureDir := filepath.Join(projectRoot, "export", "testdata", "integration", framework+"-consumer")
		generatedDir := filepath.Join(fixtureDir, "generated")

		os.RemoveAll(generatedDir)
		t.Cleanup(func() { os.RemoveAll(generatedDir) })

		runNpm(t, fixtureDir, "install")

		t.Log("Running cem export with npm: specifier...")
		stdout, stderr := runCemCommand(t, fixtureDir,
			"export",
			"--format", framework,
			"-o", generatedDir,
			"-p", "npm:@cem-examples/kitchen-sink",
		)
		t.Logf("cem export stdout: %s", stdout)
		t.Logf("cem export stderr: %s", stderr)

		assertGeneratedFiles(t, generatedDir)
		runVitest(t, fixtureDir)
	})
}

func TestExportReactIntegration(t *testing.T) {
	runExportIntegrationTest(t, "react")
}

func TestExportVueIntegration(t *testing.T) {
	runExportIntegrationTest(t, "vue")
}

func TestExportAngularIntegration(t *testing.T) {
	runExportIntegrationTest(t, "angular")
}
