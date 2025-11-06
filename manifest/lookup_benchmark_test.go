package manifest_test

import (
	"encoding/json"
	"fmt"
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
			field.Attribute = "test-attr-" + fmt.Sprintf("%d", i)
			field.Name = "testField" + fmt.Sprintf("%d", i)
			ced.Members[i] = field
		} else {
			field := &manifest.ClassField{}
			field.Name = "regularField" + fmt.Sprintf("%d", i)
			ced.Members[i] = field
		}
	}

	// Create the attribute name to search for (last one for worst case)
	attrName := "test-attr-" + fmt.Sprintf("%d", numMembers-2)

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
				Name:   "TestElement" + fmt.Sprintf("%d", i),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		} else {
			exp := &manifest.JavaScriptExport{}
			exp.Declaration = &manifest.Reference{
				Name:   "TestExport" + fmt.Sprintf("%d", i),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		}
	}

	// Create a declaration to search for (last one for worst case)
	declName := "TestElement" + fmt.Sprintf("%d", numExports-2)

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

// BenchmarkAttributeLookup_Map uses the new O(1) map-based lookup
func BenchmarkAttributeLookup_Map(b *testing.B) {
	b.Run("Small-5members", func(b *testing.B) {
		benchmarkAttributeLookupMap(b, 5)
	})
	b.Run("Medium-20members", func(b *testing.B) {
		benchmarkAttributeLookupMap(b, 20)
	})
	b.Run("Large-100members", func(b *testing.B) {
		benchmarkAttributeLookupMap(b, 100)
	})
}

func benchmarkAttributeLookupMap(b *testing.B, numMembers int) {
	// Same setup as Current version
	ced := &manifest.CustomElementDeclaration{}
	ced.Members = make([]manifest.ClassMember, numMembers)

	for i := 0; i < numMembers; i++ {
		if i%2 == 0 {
			field := &manifest.CustomElementField{}
			field.Attribute = "test-attr-" + fmt.Sprintf("%d", i)
			field.Name = "testField" + fmt.Sprintf("%d", i)
			ced.Members[i] = field
		} else {
			field := &manifest.ClassField{}
			field.Name = "regularField" + fmt.Sprintf("%d", i)
			ced.Members[i] = field
		}
	}

	attrName := "test-attr-" + fmt.Sprintf("%d", numMembers-2)

	b.ResetTimer()

	// Benchmark the new map-based lookup
	for i := 0; i < b.N; i++ {
		field := ced.LookupAttributeField(attrName)
		_ = field // Prevent optimization
	}
}

// BenchmarkExportLookup_Map uses the new O(1) map-based lookup
func BenchmarkExportLookup_Map(b *testing.B) {
	b.Run("Small-5exports", func(b *testing.B) {
		benchmarkExportLookupMap(b, 5)
	})
	b.Run("Medium-20exports", func(b *testing.B) {
		benchmarkExportLookupMap(b, 20)
	})
	b.Run("Large-100exports", func(b *testing.B) {
		benchmarkExportLookupMap(b, 100)
	})
}

func benchmarkExportLookupMap(b *testing.B, numExports int) {
	// Same setup as Current version
	mod := &manifest.Module{}
	mod.Path = "/test/module.js"
	mod.Exports = make([]manifest.Export, numExports)

	for i := 0; i < numExports; i++ {
		if i%2 == 0 {
			exp := &manifest.CustomElementExport{}
			exp.Declaration = &manifest.Reference{
				Name:   "TestElement" + fmt.Sprintf("%d", i),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		} else {
			exp := &manifest.JavaScriptExport{}
			exp.Declaration = &manifest.Reference{
				Name:   "TestExport" + fmt.Sprintf("%d", i),
				Module: mod.Path,
			}
			mod.Exports[i] = exp
		}
	}

	declName := "TestElement" + fmt.Sprintf("%d", numExports-2)

	b.ResetTimer()

	// Benchmark the new map-based lookup
	for i := 0; i < b.N; i++ {
		cee := mod.LookupCustomElementExport(declName)
		je := mod.LookupJavaScriptExport(declName)
		_, _ = cee, je // Prevent optimization
	}
}

// BenchmarkRenderableCreation measures end-to-end renderable package creation.
// This benchmark exercises the full rendering pipeline including attribute field
// lookups and export lookups using the optimized map-based implementation.
func BenchmarkRenderableCreation(b *testing.B) {
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
