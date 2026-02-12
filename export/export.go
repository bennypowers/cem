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
package export

import (
	"fmt"
	"os"
	"path/filepath"

	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

// FrameworkExportConfig holds per-framework export settings from config or flags.
type FrameworkExportConfig struct {
	Output      string `mapstructure:"output" yaml:"output"`
	StripPrefix string `mapstructure:"stripPrefix" yaml:"stripPrefix"`
	PackageName string `mapstructure:"packageName" yaml:"packageName"`
	ModuleName  string `mapstructure:"moduleName" yaml:"moduleName"`
}

// ExportElement is a template-friendly representation of a custom element.
type ExportElement struct {
	TagName    string
	ClassName  string
	ModulePath string
	ImportPath    string
	Summary       string
	Description   string

	Attributes    []ExportAttribute
	Properties    []ExportProperty
	Events        []ExportEvent
	Slots         []ExportSlot
	CssParts      []ExportCssPart
	CssProperties []ExportCssProperty

	HasDefaultSlot bool
	HasNamedSlots  bool
}

// ExportAttribute represents an attribute for template rendering.
type ExportAttribute struct {
	Name      string
	FieldName string
	Type      string
	Default   string
	Summary   string
	IsBoolean bool
}

// ExportProperty represents a property (public field not already an attribute).
type ExportProperty struct {
	Name     string
	Type     string
	Default  string
	Summary  string
	Readonly bool
}

// ExportEvent represents an event with pre-computed framework-specific names.
type ExportEvent struct {
	Name        string
	Type        string
	ReactName   string
	VueName     string
	AngularName string
	Summary     string
	// IsNative is true for standard DOM events (click, focus, etc.) that React
	// handles natively through HTMLAttributes. React wrappers should skip these
	// to avoid type conflicts.
	IsNative bool
}

// ExportSlot represents a slot.
type ExportSlot struct {
	Name    string
	Summary string
}

// ExportCssPart represents a CSS part.
type ExportCssPart struct {
	Name    string
	Summary string
}

// ExportCssProperty represents a CSS custom property.
type ExportCssProperty struct {
	Name    string
	Default string
	Syntax  string
	Summary string
}

// FrameworkExporter generates framework wrapper files for custom elements.
type FrameworkExporter interface {
	// Name returns the framework name (e.g. "react", "vue", "angular").
	Name() string
	// ExportElement generates files for a single custom element.
	// Returns a map of filename -> content.
	ExportElement(element ExportElement, cfg FrameworkExportConfig) (map[string]string, error)
	// ExportIndex generates barrel/index files for all elements.
	// Returns a map of filename -> content.
	ExportIndex(elements []ExportElement, cfg FrameworkExportConfig) (map[string]string, error)
}

// Options configures the export operation.
type Options struct {
	// Manifest is the parsed custom elements manifest.
	Manifest *M.Package
	// PackageName is the npm package name for import paths.
	PackageName string
	// Frameworks maps framework name to its config.
	Frameworks map[string]FrameworkExportConfig
}

// Export generates framework wrappers for all custom elements in the manifest.
func Export(opts Options) error {
	if opts.Manifest == nil {
		return fmt.Errorf("manifest is required")
	}

	elements := buildExportElements(opts.Manifest, opts.PackageName)
	if len(elements) == 0 {
		pterm.Warning.Println("No custom elements found in manifest")
		return nil
	}

	pterm.Info.Printfln("Found %d custom element(s)", len(elements))

	exporters := map[string]FrameworkExporter{
		"react":   &ReactExporter{},
		"vue":     &VueExporter{},
		"angular": &AngularExporter{},
	}

	for name, cfg := range opts.Frameworks {
		exporter, ok := exporters[name]
		if !ok {
			return fmt.Errorf("unknown framework: %q", name)
		}

		if cfg.Output == "" {
			return fmt.Errorf("output directory is required for framework %q", name)
		}

		pterm.Info.Printfln("Exporting %s wrappers to %s", name, cfg.Output)

		if err := os.MkdirAll(cfg.Output, 0o755); err != nil {
			return fmt.Errorf("creating output directory for %s: %w", name, err)
		}

		for _, elem := range elements {
			files, err := exporter.ExportElement(elem, cfg)
			if err != nil {
				return fmt.Errorf("exporting %s for %s: %w", elem.TagName, name, err)
			}
			for filename, content := range files {
				path := filepath.Join(cfg.Output, filename)
				if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
					return fmt.Errorf("writing %s: %w", path, err)
				}
			}
		}

		indexFiles, err := exporter.ExportIndex(elements, cfg)
		if err != nil {
			return fmt.Errorf("exporting index for %s: %w", name, err)
		}
		for filename, content := range indexFiles {
			path := filepath.Join(cfg.Output, filename)
			if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
				return fmt.Errorf("writing %s: %w", path, err)
			}
		}

		pterm.Success.Printfln("Exported %d %s wrapper(s)", len(elements), name)
	}

	return nil
}

// nativeDOMEvents lists standard DOM events that React handles natively through
// HTMLAttributes. Wrappers should not re-declare these to avoid type conflicts.
var nativeDOMEvents = map[string]bool{
	"click": true, "dblclick": true, "mousedown": true, "mouseup": true,
	"mouseover": true, "mouseout": true, "mousemove": true, "mouseenter": true,
	"mouseleave": true, "contextmenu": true,
	"keydown": true, "keyup": true, "keypress": true,
	"focus": true, "blur": true, "focusin": true, "focusout": true,
	"input": true, "change": true, "submit": true, "reset": true, "invalid": true,
	"scroll": true, "wheel": true,
	"drag": true, "dragstart": true, "dragend": true, "dragenter": true,
	"dragleave": true, "dragover": true, "drop": true,
	"touchstart": true, "touchmove": true, "touchend": true, "touchcancel": true,
	"pointerdown": true, "pointerup": true, "pointermove": true,
	"pointerenter": true, "pointerleave": true, "pointerover": true, "pointerout": true,
	"copy": true, "cut": true, "paste": true,
	"compositionstart": true, "compositionupdate": true, "compositionend": true,
	"select":         true,
	"animationstart": true, "animationend": true, "animationiteration": true,
	"transitionend": true,
	"load":          true, "error": true, "abort": true,
}

// buildExportElements transforms manifest declarations into template-friendly ExportElements.
func buildExportElements(pkg *M.Package, packageName string) []ExportElement {
	var elements []ExportElement
	for i := range pkg.Modules {
		mod := &pkg.Modules[i]
		for _, decl := range mod.Declarations {
			switch d := decl.(type) {
			case *M.CustomElementDeclaration:
				if d.TagName == "" {
					continue
				}
				elements = append(elements, buildExportElement(d, mod, packageName))
			case *M.CustomElementMixinDeclaration:
				if d.TagName == "" {
					continue
				}
				elements = append(elements, buildExportElement(&d.CustomElementDeclaration, mod, packageName))
			}
		}
	}
	return elements
}

func buildExportElement(ced *M.CustomElementDeclaration, mod *M.Module, packageName string) ExportElement {
	importPath := mod.Path
	if packageName != "" {
		importPath = packageName + "/" + mod.Path
	}

	elem := ExportElement{
		TagName:     ced.TagName,
		ClassName:   ced.Name(),
		ModulePath:  mod.Path,
		ImportPath:  importPath,
		Summary:     ced.Summary,
		Description: ced.Description,
	}

	// Attributes
	attrFieldNames := make(map[string]bool)
	for _, attr := range ced.Attributes() {
		typeName := ""
		if attr.Type != nil {
			typeName = attr.Type.Text
		}
		ea := ExportAttribute{
			Name:      attr.Name,
			FieldName: attr.FieldName,
			Type:      MapCEMType(typeName),
			Default:   attr.Default,
			Summary:   attr.Summary,
			IsBoolean: IsBooleanType(typeName),
		}
		elem.Attributes = append(elem.Attributes, ea)
		if attr.FieldName != "" {
			attrFieldNames[attr.FieldName] = true
		}
	}

	// Properties: public fields not already covered by attributes
	for _, member := range ced.Fields() {
		var cf *M.ClassField
		switch f := member.(type) {
		case *M.CustomElementField:
			cf = &f.ClassField
		case *M.ClassField:
			cf = f
		}
		if cf == nil || cf.Privacy == M.Private || cf.Privacy == M.Protected {
			continue
		}
		if attrFieldNames[cf.Name] {
			continue
		}
		typeName := ""
		if cf.Type != nil {
			typeName = cf.Type.Text
		}
		elem.Properties = append(elem.Properties, ExportProperty{
			Name:     cf.Name,
			Type:     MapCEMType(typeName),
			Default:  cf.Default,
			Summary:  cf.Summary,
			Readonly: cf.Readonly,
		})
	}

	// Events
	for _, ev := range ced.Events() {
		typeName := ""
		if ev.Type != nil {
			typeName = ev.Type.Text
		}
		elem.Events = append(elem.Events, ExportEvent{
			Name:        ev.Name,
			Type:        MapCEMType(typeName),
			ReactName:   ToReactEventName(ev.Name),
			VueName:     ToVueEventName(ev.Name),
			AngularName: ToAngularEventName(ev.Name),
			Summary:     ev.Summary,
			IsNative:    nativeDOMEvents[ev.Name],
		})
	}

	// Slots
	for _, slot := range ced.Slots() {
		elem.Slots = append(elem.Slots, ExportSlot{
			Name:    slot.Name,
			Summary: slot.Summary,
		})
		if slot.Name == "" {
			elem.HasDefaultSlot = true
		} else {
			elem.HasNamedSlots = true
		}
	}

	// CSS Parts
	for _, part := range ced.CssParts() {
		elem.CssParts = append(elem.CssParts, ExportCssPart{
			Name:    part.Name,
			Summary: part.Summary,
		})
	}

	// CSS Properties
	for _, prop := range ced.CssProperties() {
		elem.CssProperties = append(elem.CssProperties, ExportCssProperty{
			Name:    prop.Name,
			Default: prop.Default,
			Syntax:  prop.Syntax,
			Summary: prop.Summary,
		})
	}

	return elem
}
