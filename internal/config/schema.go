package config

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"text/template"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
)

//go:embed cem-config.schema.json
var configSchemaJSON []byte

var englishPrinter = message.NewPrinter(language.English)

var globFormat = &jsonschema.Format{
	Name: "x-glob-pattern",
	Validate: func(v any) error {
		s, ok := v.(string)
		if !ok {
			return nil
		}
		if !doublestar.ValidatePattern(s) {
			return fmt.Errorf("invalid glob pattern")
		}
		return nil
	},
}

var urlPatternFormat = &jsonschema.Format{
	Name: "x-url-pattern",
	Validate: func(v any) error {
		s, ok := v.(string)
		if !ok {
			return nil
		}
		if s == "" {
			return nil
		}
		if strings.ContainsAny(s, "{}[]") {
			return fmt.Errorf("invalid URL pattern: contains disallowed characters")
		}
		return nil
	},
}

var goTemplateFormat = &jsonschema.Format{
	Name: "x-go-template",
	Validate: func(v any) error {
		s, ok := v.(string)
		if !ok {
			return nil
		}
		if s == "" {
			return nil
		}
		_, err := template.New("").Parse(s)
		if err != nil {
			return fmt.Errorf("invalid Go template: %w", err)
		}
		return nil
	},
}

// ValidateSchema validates raw config data against the CemConfig JSON Schema.
// It returns schema-level errors (type mismatches, unknown keys, enum violations)
// but not semantic cross-field checks -- those live in [Validate].
// The format parameter ("yaml", "json", "jsonc") controls preprocessing;
// JSONC comments are stripped before parsing.
func ValidateSchema(configData []byte, format string) []ValidationError {
	if format == "jsonc" {
		configData = StripComments(configData)
	}
	schemaDoc, err := jsonschema.UnmarshalJSON(bytes.NewReader(configSchemaJSON))
	if err != nil {
		return []ValidationError{{
			Field:    "(schema)",
			Message:  fmt.Sprintf("internal error: failed to parse config schema: %v", err),
			Severity: SeverityError,
		}}
	}

	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	compiler.RegisterFormat(globFormat)
	compiler.RegisterFormat(urlPatternFormat)
	compiler.RegisterFormat(goTemplateFormat)

	if err := compiler.AddResource("cem-config.schema.json", schemaDoc); err != nil {
		return []ValidationError{{
			Field:    "(schema)",
			Message:  fmt.Sprintf("internal error: failed to add schema resource: %v", err),
			Severity: SeverityError,
		}}
	}

	schema, err := compiler.Compile("cem-config.schema.json")
	if err != nil {
		return []ValidationError{{
			Field:    "(schema)",
			Message:  fmt.Sprintf("internal error: failed to compile config schema: %v", err),
			Severity: SeverityError,
		}}
	}

	var doc any
	if err := yaml.Unmarshal(configData, &doc); err != nil {
		return []ValidationError{{
			Field:    "(document)",
			Message:  fmt.Sprintf("failed to parse config: %v", err),
			Severity: SeverityError,
		}}
	}
	if doc == nil {
		doc = map[string]any{}
	}
	doc = normalizeYAML(doc)

	if err := schema.Validate(doc); err != nil {
		verr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return []ValidationError{{
				Field:    "(schema)",
				Message:  fmt.Sprintf("unexpected validation error type: %v", err),
				Severity: SeverityError,
			}}
		}
		return collectSchemaErrors(verr)
	}

	return nil
}

func collectSchemaErrors(err *jsonschema.ValidationError) []ValidationError {
	var results []ValidationError
	walkSchemaErrors(err, &results)
	slices.SortFunc(results, func(a, b ValidationError) int {
		return cmp.Compare(a.Field, b.Field)
	})
	return results
}

func walkSchemaErrors(err *jsonschema.ValidationError, results *[]ValidationError) {
	for _, cause := range err.Causes {
		if len(cause.Causes) == 0 {
			field := instanceLocationToField(cause.InstanceLocation)
			msg := cause.ErrorKind.LocalizedString(englishPrinter)
			*results = append(*results, ValidationError{
				Field:    field,
				Message:  msg,
				Severity: SeverityError,
			})
		} else {
			walkSchemaErrors(cause, results)
		}
	}
}

func instanceLocationToField(parts []string) string {
	if len(parts) == 0 {
		return "(root)"
	}
	return strings.Join(parts, ".")
}

// normalizeYAML converts YAML-decoded values to JSON-compatible types.
// yaml.v3 produces map[any]any when maps have non-string keys (e.g., integer keys).
// JSON Schema requires map[string]any, so we convert and recurse.
func normalizeYAML(v any) any {
	switch v := v.(type) {
	case map[string]any:
		for k, val := range v {
			v[k] = normalizeYAML(val)
		}
		return v
	case map[any]any:
		m := make(map[string]any, len(v))
		for k, val := range v {
			m[fmt.Sprint(k)] = normalizeYAML(val)
		}
		return m
	case []any:
		for i, val := range v {
			v[i] = normalizeYAML(val)
		}
		return v
	default:
		return v
	}
}
