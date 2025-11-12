package manifest_test

import (
	"testing"

	"bennypowers.dev/cem/manifest"
)

func TestBuildAttributeFieldMap(t *testing.T) {
	t.Run("EmptyMembers", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}
		ced.Members = []manifest.ClassMember{}

		result := manifest.BuildAttributeFieldMap(ced.Members)

		if result == nil {
			t.Error("Expected non-nil map, got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected empty map, got %d entries", len(result))
		}
	})

	t.Run("NoCustomElementFields", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}
		field1 := &manifest.ClassField{}
		field1.Name = "regularField1"
		method1 := &manifest.ClassMethod{}
		method1.Name = "regularMethod1"

		ced.Members = []manifest.ClassMember{field1, method1}

		result := manifest.BuildAttributeFieldMap(ced.Members)

		if len(result) != 0 {
			t.Errorf("Expected empty map for non-CustomElementFields, got %d entries", len(result))
		}
	})

	t.Run("MultipleAttributes", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}

		field1 := &manifest.CustomElementField{}
		field1.Name = "openField"
		field1.Attribute = "open"

		field2 := &manifest.CustomElementField{}
		field2.Name = "disabledField"
		field2.Attribute = "disabled"

		regularField := &manifest.ClassField{}
		regularField.Name = "notAnAttribute"

		ced.Members = []manifest.ClassMember{field1, regularField, field2}

		result := manifest.BuildAttributeFieldMap(ced.Members)

		if len(result) != 2 {
			t.Errorf("Expected 2 entries, got %d", len(result))
		}

		if result["open"] != field1 {
			t.Error("Expected 'open' attribute to map to field1")
		}

		if result["disabled"] != field2 {
			t.Error("Expected 'disabled' attribute to map to field2")
		}

		if result["notAnAttribute"] != nil {
			t.Error("Regular field should not be in attribute map")
		}
	})

	t.Run("DuplicateAttributes", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}

		field1 := &manifest.CustomElementField{}
		field1.Name = "openField1"
		field1.Attribute = "open"

		field2 := &manifest.CustomElementField{}
		field2.Name = "openField2"
		field2.Attribute = "open"

		ced.Members = []manifest.ClassMember{field1, field2}

		result := manifest.BuildAttributeFieldMap(ced.Members)

		if len(result) != 1 {
			t.Errorf("Expected 1 entry for duplicate attribute, got %d", len(result))
		}

		// Last one should win (consistent with current linear search behavior)
		if result["open"] != field2 {
			t.Error("Expected last duplicate attribute to win")
		}
	})

	t.Run("EmptyAttributeName", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}

		field1 := &manifest.CustomElementField{}
		field1.Name = "fieldWithNoAttribute"
		field1.Attribute = "" // Empty attribute name

		field2 := &manifest.CustomElementField{}
		field2.Name = "validField"
		field2.Attribute = "valid"

		ced.Members = []manifest.ClassMember{field1, field2}

		result := manifest.BuildAttributeFieldMap(ced.Members)

		if len(result) != 1 {
			t.Errorf("Expected 1 entry (empty attributes excluded), got %d", len(result))
		}

		if result[""] != nil {
			t.Error("Empty attribute name should not be in map")
		}

		if result["valid"] != field2 {
			t.Error("Valid attribute should be in map")
		}
	})
}

func TestBuildExportMaps(t *testing.T) {
	t.Run("EmptyExports", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"
		mod.Exports = []manifest.Export{}

		ceeMap, jeMap := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if ceeMap == nil || jeMap == nil {
			t.Error("Expected non-nil maps, got nil")
		}
		if len(ceeMap) != 0 || len(jeMap) != 0 {
			t.Error("Expected empty maps for no exports")
		}
	})

	t.Run("SingleCustomElementExport", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		cee := &manifest.CustomElementExport{}
		cee.Declaration = &manifest.Reference{
			Name:   "MyElement",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{cee}

		ceeMap, jeMap := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if len(ceeMap) != 1 {
			t.Errorf("Expected 1 CustomElementExport, got %d", len(ceeMap))
		}
		if len(jeMap) != 0 {
			t.Errorf("Expected 0 JavaScriptExports, got %d", len(jeMap))
		}
		if ceeMap["MyElement"] != cee {
			t.Error("Expected MyElement to map to cee")
		}
	})

	t.Run("SingleJavaScriptExport", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		je := &manifest.JavaScriptExport{}
		je.Declaration = &manifest.Reference{
			Name:   "myFunction",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{je}

		ceeMap, jeMap := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if len(ceeMap) != 0 {
			t.Errorf("Expected 0 CustomElementExports, got %d", len(ceeMap))
		}
		if len(jeMap) != 1 {
			t.Errorf("Expected 1 JavaScriptExport, got %d", len(jeMap))
		}
		if jeMap["myFunction"] != je {
			t.Error("Expected myFunction to map to je")
		}
	})

	t.Run("MixedExports", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		cee := &manifest.CustomElementExport{}
		cee.Declaration = &manifest.Reference{
			Name:   "MyElement",
			Module: mod.Path,
		}

		je := &manifest.JavaScriptExport{}
		je.Declaration = &manifest.Reference{
			Name:   "myFunction",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{cee, je}

		ceeMap, jeMap := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if len(ceeMap) != 1 || len(jeMap) != 1 {
			t.Errorf("Expected 1 of each type, got %d CEE, %d JE", len(ceeMap), len(jeMap))
		}
	})

	t.Run("DuplicateDeclarationNames", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		cee1 := &manifest.CustomElementExport{}
		cee1.Declaration = &manifest.Reference{
			Name:   "MyElement",
			Module: mod.Path,
		}

		cee2 := &manifest.CustomElementExport{}
		cee2.Declaration = &manifest.Reference{
			Name:   "MyElement",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{cee1, cee2}

		ceeMap, _ := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if len(ceeMap) != 1 {
			t.Errorf("Expected 1 entry for duplicate names, got %d", len(ceeMap))
		}

		// Last one should win
		if ceeMap["MyElement"] != cee2 {
			t.Error("Expected last duplicate to win")
		}
	})

	t.Run("ModulePathFiltering", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		// Export with matching module path
		cee1 := &manifest.CustomElementExport{}
		cee1.Declaration = &manifest.Reference{
			Name:   "LocalElement",
			Module: mod.Path,
		}

		// Export with empty module (should be included)
		cee2 := &manifest.CustomElementExport{}
		cee2.Declaration = &manifest.Reference{
			Name:   "EmptyModuleElement",
			Module: "",
		}

		// Export with different module path (should be excluded)
		cee3 := &manifest.CustomElementExport{}
		cee3.Declaration = &manifest.Reference{
			Name:   "OtherModuleElement",
			Module: "/other/module.js",
		}

		mod.Exports = []manifest.Export{cee1, cee2, cee3}

		ceeMap, _ := manifest.BuildExportMaps(mod.Exports, mod.Path)

		if len(ceeMap) != 2 {
			t.Errorf("Expected 2 entries (matching and empty module), got %d", len(ceeMap))
		}

		if ceeMap["LocalElement"] != cee1 {
			t.Error("Expected LocalElement to be included")
		}

		if ceeMap["EmptyModuleElement"] != cee2 {
			t.Error("Expected EmptyModuleElement to be included")
		}

		if ceeMap["OtherModuleElement"] != nil {
			t.Error("Expected OtherModuleElement to be excluded (different module path)")
		}
	})
}

func TestLookupAttributeField(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}

		field := &manifest.CustomElementField{}
		field.Name = "openField"
		field.Attribute = "open"

		ced.Members = []manifest.ClassMember{field}

		result := ced.LookupAttributeField("open")

		if result != field {
			t.Error("Expected to find field for 'open' attribute")
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		ced := &manifest.CustomElementDeclaration{}

		field := &manifest.CustomElementField{}
		field.Name = "openField"
		field.Attribute = "open"

		ced.Members = []manifest.ClassMember{field}

		result := ced.LookupAttributeField("disabled")

		if result != nil {
			t.Error("Expected nil for non-existent attribute")
		}
	})

	t.Run("NilDeclaration", func(t *testing.T) {
		var ced *manifest.CustomElementDeclaration

		result := ced.LookupAttributeField("open")

		if result != nil {
			t.Error("Expected nil for nil declaration")
		}
	})
}

func TestLookupExports(t *testing.T) {
	t.Run("LookupCustomElementExport_Found", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		cee := &manifest.CustomElementExport{}
		cee.Declaration = &manifest.Reference{
			Name:   "MyElement",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{cee}

		result := mod.LookupCustomElementExport("MyElement")

		if result != cee {
			t.Error("Expected to find CustomElementExport")
		}
	})

	t.Run("LookupCustomElementExport_NotFound", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"
		mod.Exports = []manifest.Export{}

		result := mod.LookupCustomElementExport("NonExistent")

		if result != nil {
			t.Error("Expected nil for non-existent export")
		}
	})

	t.Run("LookupJavaScriptExport_Found", func(t *testing.T) {
		mod := &manifest.Module{}
		mod.Path = "/test/module.js"

		je := &manifest.JavaScriptExport{}
		je.Declaration = &manifest.Reference{
			Name:   "myFunction",
			Module: mod.Path,
		}

		mod.Exports = []manifest.Export{je}

		result := mod.LookupJavaScriptExport("myFunction")

		if result != je {
			t.Error("Expected to find JavaScriptExport")
		}
	})

	t.Run("NilModule", func(t *testing.T) {
		var mod *manifest.Module

		ceeResult := mod.LookupCustomElementExport("test")
		jeResult := mod.LookupJavaScriptExport("test")

		if ceeResult != nil || jeResult != nil {
			t.Error("Expected nil results for nil module")
		}
	})
}
