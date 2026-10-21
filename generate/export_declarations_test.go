/*
Copyright © 2026 Benny Powers

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
package generate

import (
	"testing"

	"bennypowers.dev/cem/internal/platform"
	M "bennypowers.dev/cem/manifest"
)

// Inline: pure boolean classifiers over single statements, table-driven
// cases are clearer than goldens (same rationale as TestIsTypeOnlyNode)

func TestIsDefaultExportStatement(t *testing.T) {
	tests := []struct {
		name string
		stmt string
		want bool
	}{
		{"default class", "export default class Foo {}", true},
		{"default function", "export default function foo() {}", true},
		{"default identifier", "export default foo;", true},
		{"default anonymous arrow", "export default () => 1;", true},
		{"named class", "export class Foo {}", false},
		{"named function", "export function foo() {}", false},
		{"local named export", "export { foo };", false},
		{"named re-export", "export { foo } from './foo.js';", false},
		{"star re-export", "export * from './foo.js';", false},
		{"star alias re-export", "export * as ns from './foo.js';", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, tree := parseStatement(tt.stmt)
			if tree == nil {
				t.Fatal("failed to parse statement")
			}
			defer tree.Close()
			cursor := root.Walk()
			defer cursor.Close()
			for _, child := range root.NamedChildren(cursor) {
				if got := isDefaultExportStatement(&child); got != tt.want {
					t.Errorf("isDefaultExportStatement(%q) = %v, want %v", tt.stmt, got, tt.want)
				}
			}
		})
	}

	t.Run("nil node", func(t *testing.T) {
		if isDefaultExportStatement(nil) {
			t.Error("isDefaultExportStatement(nil) = true, want false")
		}
	})
}

func TestHasFromClause(t *testing.T) {
	tests := []struct {
		name string
		stmt string
		want bool
	}{
		{"named re-export", "export { foo } from './foo.js';", true},
		{"star re-export", "export * from './foo.js';", true},
		{"star alias re-export", "export * as ns from './foo.js';", true},
		{"named import", "import { foo } from './foo.js';", true},
		{"default import", "import foo from './foo.js';", true},
		{"local named export", "export { foo as bar };", false},
		{"default export", "export default foo;", false},
		{"named class", "export class Foo {}", false},
		{"side-effect import", "import './foo.js';", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, tree := parseStatement(tt.stmt)
			if tree == nil {
				t.Fatal("failed to parse statement")
			}
			defer tree.Close()
			cursor := root.Walk()
			defer cursor.Close()
			for _, child := range root.NamedChildren(cursor) {
				if got := hasFromClause(&child); got != tt.want {
					t.Errorf("hasFromClause(%q) = %v, want %v", tt.stmt, got, tt.want)
				}
			}
		})
	}

	t.Run("nil node", func(t *testing.T) {
		if hasFromClause(nil) {
			t.Error("hasFromClause(nil) = true, want false")
		}
	})
}

// Inline: single-return lookup against a synthetic MapFS, table-driven cases
// are clearer than goldens

func TestResolveExtensionlessSpec(t *testing.T) {
	mfs := platform.NewMapFileSystem(nil)
	for _, f := range []string{
		"src/plain.ts",
		"src/typed.tsx",
		"src/module.mts",
		"src/existing.js",
		"src/withindex/index.ts",
		"src/sub/thing.ts",
		"src/foo.ts",
		"src/nested/decoyed.ts",
	} {
		if err := mfs.WriteFile(f, []byte("export {};\n"), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", f, err)
		}
	}
	// a directory whose name looks like a .ts file must not resolve as one
	if err := mfs.MkdirAll("src/directory.ts", 0755); err != nil {
		t.Fatalf("failed to mkdir: %v", err)
	}

	tests := []struct {
		name         string
		absPath      string
		importSpec   string
		resolvedPath string
		want         string
	}{
		{"ts file", "src/importer.ts", "./plain", "src/plain", "src/plain.js"},
		{"tsx file", "src/importer.ts", "./typed", "src/typed", "src/typed.js"},
		{"mts file", "src/importer.ts", "./module", "src/module", "src/module.mjs"},
		{"js file", "src/importer.ts", "./existing", "src/existing", "src/existing.js"},
		{"index file", "src/importer.ts", "./withindex", "src/withindex", "src/withindex/index.js"},
		{"missing", "src/importer.ts", "./nope", "src/nope", ""},
		{"directory with extension-like name", "src/importer.ts", "./directory", "src/directory", ""},
		{"subdirectory spec", "src/importer.ts", "./sub/thing", "src/sub/thing", "src/sub/thing.js"},
		{"parent directory spec", "src/nested/importer.ts", "../foo", "src/foo", "src/foo.js"},
		{"decoy next to nested importer", "src/nested/importer.ts", "../decoyed", "src/decoyed", ""},
		{"degenerate current-directory spec", "src/importer.ts", "./", "src", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := &ModuleProcessor{fs: mfs, absPath: tt.absPath}
			if got := mp.resolveExtensionlessSpec(tt.importSpec, tt.resolvedPath); got != tt.want {
				t.Errorf("resolveExtensionlessSpec(%q, %q) = %q, want %q", tt.importSpec, tt.resolvedPath, got, tt.want)
			}
		})
	}

	t.Run("nil filesystem", func(t *testing.T) {
		mp := &ModuleProcessor{fs: nil, absPath: "src/importer.ts"}
		if got := mp.resolveExtensionlessSpec("./plain", "src/plain"); got != "" {
			t.Errorf("resolveExtensionlessSpec with nil fs = %q, want %q", got, "")
		}
	})
}

// Inline: single-return path resolution over a synthetic MapFS, table-driven
// cases are clearer than goldens
func TestResolveImportSpecRelativeGate(t *testing.T) {
	mfs := platform.NewMapFileSystem(nil)
	if err := mfs.WriteFile("src/plain.ts", []byte("export {};\n"), 0644); err != nil {
		t.Fatalf("failed to write src/plain.ts: %v", err)
	}
	mp := &ModuleProcessor{
		fs:      mfs,
		absPath: "src/importer.ts",
		module:  M.NewModule("src/importer.js"),
	}

	// bare specifiers name packages, so they must stay on the plain
	// package-resolution path even when a same-named file sits next to the
	// importer on disk
	if got := mp.resolveImportSpec("plain"); got != "src/plain" {
		t.Errorf("resolveImportSpec(\"plain\") = %q, want %q", got, "src/plain")
	}
	if got := mp.resolveImportSpec("./plain"); got != "src/plain.js" {
		t.Errorf("resolveImportSpec(\"./plain\") = %q, want %q", got, "src/plain.js")
	}

	// Root-escaping relative specifiers are rejected to the raw specifier,
	// mirroring the package branch above: a resolved path that is ".." or
	// ".."-prefixed would name a file outside the package, so it stays
	// unresolved rather than emitting a traversal-shaped manifest path.
	// A single ".." from src/ collapses into the root and stays in-root.
	if got := mp.resolveImportSpec("../evil.js"); got != "evil.js" {
		t.Errorf("resolveImportSpec(\"../evil.js\") = %q, want %q", got, "evil.js")
	}
	if got := mp.resolveImportSpec("../../evil.js"); got != "../../evil.js" {
		t.Errorf("resolveImportSpec(\"../../evil.js\") = %q, want %q", got, "../../evil.js")
	}
	if got := mp.resolveImportSpec("../../../evil.js"); got != "../../../evil.js" {
		t.Errorf("resolveImportSpec(\"../../../evil.js\") = %q, want %q", got, "../../../evil.js")
	}

	// in-root ".." resolution keeps working: Join absorbs the segments
	// into the package root rather than escaping it
	deep := &ModuleProcessor{
		fs:      mfs,
		absPath: "elements/x/y.ts",
		module:  M.NewModule("elements/x/y.js"),
	}
	if got := deep.resolveImportSpec("../../base/x.js"); got != "base/x.js" {
		t.Errorf("resolveImportSpec(\"../../base/x.js\") = %q, want %q", got, "base/x.js")
	}
}

// Inline: single-return path resolution for the package-specifier branch,
// table-driven cases are clearer than goldens
func TestResolveImportSpecPackageBranch(t *testing.T) {
	mp := &ModuleProcessor{
		logger: NewLogCtx("src/importer.ts"),
		module: M.NewModule("src/importer.js"),
	}
	tests := []struct {
		name string
		spec string
		want string
	}{
		{"plain subpath", "@scope/pkg/lib/types.js", "lib/types.js"},
		{"interior dot-dot collapses", "@scope/pkg/lib/../types.js", "types.js"},
		{"current-directory segment collapses", "@scope/pkg/./lib/types.js", "lib/types.js"},
		{"escaping dot-dot rejected to raw specifier", "@scope/pkg/../evil.js", "@scope/pkg/../evil.js"},
		{"deep escape rejected to raw specifier", "@scope/pkg/a/../../evil.js", "@scope/pkg/a/../../evil.js"},
		{"two-part scope specifier stays raw", "@scope/pkg", "@scope/pkg"},
		{"trailing slash keeps empty subpath", "@scope/pkg/", ""},
		{"double-slash absolute subpath rejected to raw specifier", "@scope/pkg//../evil.js", "@scope/pkg//../evil.js"},
		{"absolute subpath rejected to raw specifier", "@scope/pkg//etc/evil.js", "@scope/pkg//etc/evil.js"},
		{"dot-dot to package root normalizes to dot", "@scope/pkg/lib/..", "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mp.resolveImportSpec(tt.spec); got != tt.want {
				t.Errorf("resolveImportSpec(%q) = %q, want %q", tt.spec, got, tt.want)
			}
		})
	}
}
