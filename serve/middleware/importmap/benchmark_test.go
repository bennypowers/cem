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

package importmap

import (
	"os"
	"path/filepath"
	"testing"
)

// BenchmarkGenerate_SmallProject benchmarks import map generation for a small project
func BenchmarkGenerate_SmallProject(b *testing.B) {
	tmpDir := b.TempDir()

	// Create a small project with a few dependencies
	packageJSON := `{
  "name": "small-project",
  "dependencies": {
    "lit": "^3.0.0",
    "tslib": "^2.0.0"
  }
}`
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		b.Fatal(err)
	}

	// Create node_modules with mock packages
	packages := map[string]string{
		"lit": `{
  "name": "lit",
  "exports": {
    ".": {"import": "./index.js"},
    "./decorators.js": {"import": "./decorators.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0",
    "lit-element": "^4.0.0"
  }
}`,
		"tslib": `{
  "name": "tslib",
  "exports": {
    ".": {"import": "./tslib.es6.mjs"}
  }
}`,
		"@lit/reactive-element": `{
  "name": "@lit/reactive-element",
  "exports": {
    ".": {"import": "./reactive-element.js"}
  }
}`,
		"lit-html": `{
  "name": "lit-html",
  "exports": {
    ".": {"import": "./lit-html.js"},
    "./directives/repeat.js": {"import": "./directives/repeat.js"}
  }
}`,
		"lit-element": `{
  "name": "lit-element",
  "exports": {
    ".": {"import": "./lit-element.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0"
  }
}`,
	}

	for name, content := range packages {
		pkgDir := filepath.Join(tmpDir, "node_modules", name)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Generate(tmpDir, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerate_MediumProject benchmarks import map generation for a medium project
func BenchmarkGenerate_MediumProject(b *testing.B) {
	tmpDir := b.TempDir()

	// Create a medium project with more dependencies
	packageJSON := `{
  "name": "medium-project",
  "dependencies": {
    "lit": "^3.0.0",
    "tslib": "^2.0.0",
    "@patternfly/elements": "^3.0.0",
    "fuse.js": "^7.0.0",
    "prismjs": "^1.29.0"
  }
}`
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		b.Fatal(err)
	}

	// Create node_modules with mock packages
	packages := map[string]string{
		"lit": `{
  "name": "lit",
  "exports": {
    ".": {"import": "./index.js"},
    "./decorators.js": {"import": "./decorators.js"},
    "./directives/*": {"import": "./directives/*.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0",
    "lit-element": "^4.0.0"
  }
}`,
		"tslib": `{
  "name": "tslib",
  "exports": {
    ".": {"import": "./tslib.es6.mjs"}
  }
}`,
		"@lit/reactive-element": `{
  "name": "@lit/reactive-element",
  "exports": {
    ".": {"import": "./reactive-element.js"},
    "./decorators/*": {"import": "./decorators/*.js"}
  }
}`,
		"lit-html": `{
  "name": "lit-html",
  "exports": {
    ".": {"import": "./lit-html.js"},
    "./directives/*": {"import": "./directives/*.js"}
  }
}`,
		"lit-element": `{
  "name": "lit-element",
  "exports": {
    ".": {"import": "./lit-element.js"},
    "./decorators.js": {"import": "./decorators.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0"
  }
}`,
		"@patternfly/elements": `{
  "name": "@patternfly/elements",
  "exports": {
    ".": {"import": "./pfe.js"},
    "./pf-button/pf-button.js": {"import": "./pf-button/pf-button.js"},
    "./pf-card/pf-card.js": {"import": "./pf-card/pf-card.js"},
    "./pf-accordion/pf-accordion.js": {"import": "./pf-accordion/pf-accordion.js"}
  },
  "dependencies": {
    "lit": "^3.0.0",
    "@patternfly/pfe-core": "^3.0.0"
  }
}`,
		"@patternfly/pfe-core": `{
  "name": "@patternfly/pfe-core",
  "exports": {
    ".": {"import": "./core.js"},
    "./decorators.js": {"import": "./decorators.js"}
  },
  "dependencies": {
    "lit": "^3.0.0"
  }
}`,
		"fuse.js": `{
  "name": "fuse.js",
  "exports": {
    ".": {"import": "./dist/fuse.mjs"}
  }
}`,
		"prismjs": `{
  "name": "prismjs",
  "main": "prism.js"
}`,
	}

	for name, content := range packages {
		pkgDir := filepath.Join(tmpDir, "node_modules", name)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Generate(tmpDir, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerate_RealProject_RHDS benchmarks import map generation for RHDS
// This requires the RHDS project to be available at the specified path
func BenchmarkGenerate_RealProject_RHDS(b *testing.B) {
	rhdsPath := "/home/bennyp/Developer/redhat-ux/red-hat-design-system"
	if _, err := os.Stat(filepath.Join(rhdsPath, "package.json")); os.IsNotExist(err) {
		b.Skip("RHDS project not available at " + rhdsPath)
	}
	if _, err := os.Stat(filepath.Join(rhdsPath, "node_modules")); os.IsNotExist(err) {
		b.Skip("RHDS node_modules not installed")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Generate(rhdsPath, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerate_RealProject_PFE benchmarks import map generation for PatternFly Elements
// This requires the PFE project to be available at the specified path
func BenchmarkGenerate_RealProject_PFE(b *testing.B) {
	pfePath := "/home/bennyp/Developer/patternfly/patternfly-elements"
	if _, err := os.Stat(filepath.Join(pfePath, "package.json")); os.IsNotExist(err) {
		b.Skip("PFE project not available at " + pfePath)
	}
	if _, err := os.Stat(filepath.Join(pfePath, "node_modules")); os.IsNotExist(err) {
		b.Skip("PFE node_modules not installed")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Generate(pfePath, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerateWithGraph_SmallProject benchmarks initial resolution with graph
func BenchmarkGenerateWithGraph_SmallProject(b *testing.B) {
	tmpDir := b.TempDir()

	// Create a small project with a few dependencies
	packageJSON := `{
  "name": "small-project",
  "dependencies": {
    "lit": "^3.0.0",
    "tslib": "^2.0.0"
  }
}`
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		b.Fatal(err)
	}

	// Create node_modules with mock packages
	packages := map[string]string{
		"lit": `{
  "name": "lit",
  "exports": {
    ".": {"import": "./index.js"},
    "./decorators.js": {"import": "./decorators.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0",
    "lit-element": "^4.0.0"
  }
}`,
		"tslib": `{
  "name": "tslib",
  "exports": {
    ".": {"import": "./tslib.es6.mjs"}
  }
}`,
		"@lit/reactive-element": `{
  "name": "@lit/reactive-element",
  "exports": {
    ".": {"import": "./reactive-element.js"}
  }
}`,
		"lit-html": `{
  "name": "lit-html",
  "exports": {
    ".": {"import": "./lit-html.js"},
    "./directives/repeat.js": {"import": "./directives/repeat.js"}
  }
}`,
		"lit-element": `{
  "name": "lit-element",
  "exports": {
    ".": {"import": "./lit-element.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0"
  }
}`,
	}

	for name, content := range packages {
		pkgDir := filepath.Join(tmpDir, "node_modules", name)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateWithGraph(tmpDir, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerateIncremental_SmallProject benchmarks incremental resolution
func BenchmarkGenerateIncremental_SmallProject(b *testing.B) {
	tmpDir := b.TempDir()

	// Create a small project with a few dependencies
	packageJSON := `{
  "name": "small-project",
  "dependencies": {
    "lit": "^3.0.0",
    "tslib": "^2.0.0"
  }
}`
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		b.Fatal(err)
	}

	// Create node_modules with mock packages
	packages := map[string]string{
		"lit": `{
  "name": "lit",
  "exports": {
    ".": {"import": "./index.js"},
    "./decorators.js": {"import": "./decorators.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0",
    "lit-element": "^4.0.0"
  }
}`,
		"tslib": `{
  "name": "tslib",
  "exports": {
    ".": {"import": "./tslib.es6.mjs"}
  }
}`,
		"@lit/reactive-element": `{
  "name": "@lit/reactive-element",
  "exports": {
    ".": {"import": "./reactive-element.js"}
  }
}`,
		"lit-html": `{
  "name": "lit-html",
  "exports": {
    ".": {"import": "./lit-html.js"},
    "./directives/repeat.js": {"import": "./directives/repeat.js"}
  }
}`,
		"lit-element": `{
  "name": "lit-element",
  "exports": {
    ".": {"import": "./lit-element.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0"
  }
}`,
	}

	for name, content := range packages {
		pkgDir := filepath.Join(tmpDir, "node_modules", name)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	// Generate initial state with graph
	initial, err := GenerateWithGraph(tmpDir, nil)
	if err != nil {
		b.Fatal(err)
	}

	update := IncrementalUpdate{
		ChangedPackages: []string{"lit"},
		PreviousMap:     initial.ImportMap,
		PreviousGraph:   initial.Graph,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateIncremental(tmpDir, nil, update)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerateIncremental_RealProject_RHDS benchmarks incremental resolution for RHDS
func BenchmarkGenerateIncremental_RealProject_RHDS(b *testing.B) {
	rhdsPath := "/home/bennyp/Developer/redhat-ux/red-hat-design-system"
	if _, err := os.Stat(filepath.Join(rhdsPath, "package.json")); os.IsNotExist(err) {
		b.Skip("RHDS project not available at " + rhdsPath)
	}
	if _, err := os.Stat(filepath.Join(rhdsPath, "node_modules")); os.IsNotExist(err) {
		b.Skip("RHDS node_modules not installed")
	}

	// Generate initial state with graph
	initial, err := GenerateWithGraph(rhdsPath, nil)
	if err != nil {
		b.Fatal(err)
	}

	update := IncrementalUpdate{
		ChangedPackages: []string{"lit"},
		PreviousMap:     initial.ImportMap,
		PreviousGraph:   initial.Graph,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateIncremental(rhdsPath, nil, update)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGenerate_WorkspaceMode benchmarks workspace mode import map generation
func BenchmarkGenerate_WorkspaceMode(b *testing.B) {
	tmpDir := b.TempDir()

	// Create a monorepo with workspace packages
	packageJSON := `{
  "name": "monorepo",
  "workspaces": ["packages/*"]
}`
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		b.Fatal(err)
	}

	// Create workspace packages
	workspacePackages := map[string]string{
		"packages/ui": `{
  "name": "@myorg/ui",
  "exports": {
    ".": {"import": "./index.js"},
    "./button": {"import": "./button.js"},
    "./card": {"import": "./card.js"}
  },
  "dependencies": {
    "lit": "^3.0.0"
  }
}`,
		"packages/icons": `{
  "name": "@myorg/icons",
  "exports": {
    ".": {"import": "./index.js"},
    "./*": {"import": "./*.js"}
  }
}`,
		"packages/utils": `{
  "name": "@myorg/utils",
  "exports": {
    ".": {"import": "./index.js"}
  }
}`,
	}

	for path, content := range workspacePackages {
		pkgDir := filepath.Join(tmpDir, path)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	// Create node_modules
	nodeModulesPackages := map[string]string{
		"lit": `{
  "name": "lit",
  "exports": {
    ".": {"import": "./index.js"}
  },
  "dependencies": {
    "@lit/reactive-element": "^2.0.0",
    "lit-html": "^3.0.0"
  }
}`,
		"@lit/reactive-element": `{
  "name": "@lit/reactive-element",
  "exports": {
    ".": {"import": "./reactive-element.js"}
  }
}`,
		"lit-html": `{
  "name": "lit-html",
  "exports": {
    ".": {"import": "./lit-html.js"}
  }
}`,
	}

	for name, content := range nodeModulesPackages {
		pkgDir := filepath.Join(tmpDir, "node_modules", name)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			b.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(content), 0644); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Generate(tmpDir, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
