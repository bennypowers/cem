package manifest_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/manifest"
)

// BenchmarkAttributeLookup_Current establishes baseline performance for
// the current linear search implementation through Members.
func BenchmarkAttributeLookup_Current(b *testing.B) {
	b.Run("Small-5members", func(b *testing.B) {
		benchmarkAttributeLookupCurrent(b, 5)
	})
	b.Run("Medium-20members", func(b *testing.B) {
		benchmarkAttributeLookupCurrent(b, 20)
	})
	b.Run("Large-100members", func(b *testing.B) {
		benchmarkAttributeLookupCurrent(b, 100)
	})
}

func benchmarkAttributeLookupCurrent(b *testing.B, numMembers int) {
	// Create a CustomElementDeclaration with N members
	ced := &manifest.CustomElementDeclaration{}
	ced.Members = make([]manifest.ClassMember, numMembers)

	// Fill with half CustomElementFields (with attributes) and half regular fields
	for i := 0; i < numMembers; i++ {
		if i%2 == 0 {
			field := &manifest.CustomElementField{}
			field.Attribute = "test-attr-" + string(rune('a'+i))
			field.Name = "testField" + string(rune('A'+i))
			ced.Members[i] = field
		} else {
			field := &manifest.ClassField{}
			field.Name = "regularField" + string(rune('A'+i))
			ced.Members[i] = field
		}
	}

	// Create the attribute name to search for (last one for worst case)
	attrName := "test-attr-" + string(rune('a'+numMembers-2))

	b.ResetTimer()

	// Benchmark the current linear search pattern
	for i := 0; i < b.N; i++ {
		var field *manifest.CustomElementField
		// Current implementation: O(n) linear search
		for _, f := range ced.Members {
			if cef, ok := f.(*manifest.CustomElementField); ok {
				if cef.Attribute == attrName {
					field = cef
					break
				}
			}
		}
		_ = field // Prevent optimization
	}
}

// BenchmarkExportLookup_Current establishes baseline for export lookups
func BenchmarkExportLookup_Current(b *testing.B) {
	b.Run("Small-5exports", func(b *testing.B) {
		benchmarkExportLookupCurrent(b, 5)
	})
	b.Run("Medium-20exports", func(b *testing.B) {
		benchmarkExportLookupCurrent(b, 20)
	})
	b.Run("Large-100exports", func(b *testing.B) {
		benchmarkExportLookupCurrent(b, 100)
	})
}

func benchmarkExportLookupCurrent(b *testing.B, numExports int) {
	// Create a Module with N exports
	mod := &manifest.Module{}
	mod.Path = "/test/module.js"
	mod.Exports = make([]manifest.Export, numExports)

	// Fill with mix of CustomElementExport and JavaScriptExport
	for i := 0; i < numExports; i++ {
		if i%2 == 0 {
			exp := &manifest.CustomElementExport{}
			exp.Declaration = &manifest.Reference{
				Name:   "TestElement" + string(rune('A'+i)),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		} else {
			exp := &manifest.JavaScriptExport{}
			exp.Declaration = &manifest.Reference{
				Name:   "TestExport" + string(rune('A'+i)),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		}
	}

	// Create a declaration to search for (last one for worst case)
	declName := "TestElement" + string(rune('A'+numExports-2))

	b.ResetTimer()

	// Benchmark the current linear search pattern
	for i := 0; i < b.N; i++ {
		var cee *manifest.CustomElementExport
		var je *manifest.JavaScriptExport

		// Current implementation: O(n) linear search
		for j, exp := range mod.Exports {
			if ecee, ok := exp.(*manifest.CustomElementExport); ok {
				if ecee != nil && ecee.Declaration != nil &&
					ecee.Declaration.Name == declName &&
					(ecee.Declaration.Module == "" || ecee.Declaration.Module == mod.Path) {
					cee = mod.Exports[j].(*manifest.CustomElementExport)
				}
			}
			if eje, ok := exp.(*manifest.JavaScriptExport); ok {
				if eje.Declaration.Name == declName &&
					(eje.Declaration.Module == "" || eje.Declaration.Module == mod.Path) {
					je = mod.Exports[j].(*manifest.JavaScriptExport)
				}
			}
		}
		_, _ = cee, je // Prevent optimization
	}
}

// BenchmarkRenderableCreation_Current measures end-to-end renderable creation
func BenchmarkRenderableCreation_Current(b *testing.B) {
	// Load a realistic manifest with multiple elements
	manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "custom-element-member-grouping.json"))
	if err != nil {
		b.Fatalf("Failed to load fixture: %v", err)
	}

	var pkg manifest.Package
	if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
		b.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = manifest.NewRenderablePackage(&pkg)
	}
}
