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
package validate

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/viper"
	"golang.org/x/mod/semver"
)

// Embedded schemas for all custom-elements-manifest versions
// Includes a speculative 2.1.1 schema to work around issues in 2.1.0
// See: https://github.com/webcomponents/custom-elements-manifest/issues/138
// See: https://github.com/vega/ts-json-schema-generator/pull/2323
//go:embed schemas/*.json
var embeddedSchemas embed.FS

type ValidationResult struct {
	IsValid       bool                `json:"valid"`
	Path          string              `json:"path,omitempty"`
	SchemaVersion string              `json:"schemaVersion,omitempty"`
	Errors        []ValidationError   `json:"errors"`
	Warnings      []ValidationWarning `json:"warnings"`
}

type ValidationOptions struct {
	IncludeWarnings bool
	DisabledRules   []string
}

// Validate validates a custom-elements.json manifest
func Validate(manifestPath string, options ValidationOptions) (*ValidationResult, error) {
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error reading manifest file: %w", err)
	}

	pipeline, err := NewValidationPipeline(manifestData)
	if err != nil {
		return nil, err
	}

	return pipeline.Validate(options)
}

// ValidationPipeline orchestrates the validation process
type ValidationPipeline struct {
	manifestData      []byte
	schemaVersion     string
	navigator         *ManifestNavigator
	warningProcessor  *WarningProcessor
	errorProcessor    *ErrorProcessor
}

// NewValidationPipeline creates a new validation pipeline
func NewValidationPipeline(manifestData []byte) (*ValidationPipeline, error) {
	// Parse manifest once for efficiency
	var manifestJSON map[string]any
	if err := json.Unmarshal(manifestData, &manifestJSON); err != nil {
		return nil, fmt.Errorf("error parsing manifest: %w", err)
	}

	// Extract schema version
	var versionStruct struct {
		SchemaVersion string `json:"schemaVersion"`
	}
	if err := json.Unmarshal(manifestData, &versionStruct); err != nil {
		return nil, fmt.Errorf("error parsing manifest to find schemaVersion: %w", err)
	}

	if versionStruct.SchemaVersion == "" {
		return nil, fmt.Errorf("schemaVersion not found in manifest")
	}

	// Create navigator and processors
	navigator := NewManifestNavigator(manifestJSON)
	warningProcessor := NewWarningProcessor()
	errorProcessor := NewErrorProcessor(navigator)

	return &ValidationPipeline{
		manifestData:     manifestData,
		schemaVersion:    versionStruct.SchemaVersion,
		navigator:        navigator,
		warningProcessor: warningProcessor,
		errorProcessor:   errorProcessor,
	}, nil
}

// Validate performs the complete validation process
func (p *ValidationPipeline) Validate(options ValidationOptions) (*ValidationResult, error) {
	result := &ValidationResult{
		SchemaVersion: p.schemaVersion,
		Errors:        []ValidationError{},
		Warnings:      []ValidationWarning{},
	}

	// Add schema version warning if needed
	if err := p.checkSchemaVersion(result); err != nil {
		return nil, err
	}

	// Perform JSON schema validation
	if err := p.validateSchema(result); err != nil {
		return nil, err
	}

	// Process warnings if requested
	if options.IncludeWarnings {
		warnings := p.warningProcessor.ProcessWarnings(p.navigator)
		filtered := filterWarningsByConfig(warnings, options.DisabledRules)
		result.Warnings = append(result.Warnings, filtered...)
	}

	return result, nil
}

func (p *ValidationPipeline) checkSchemaVersion(result *ValidationResult) error {
	// Only warn for versions < 2.1.0, since we handle 2.1.0 with our speculative schema
	if isSemverLessThan(p.schemaVersion, "2.1.0") {
		result.Warnings = append(result.Warnings, ValidationWarning{
			ID:       "schema-version-old",
			Message:  fmt.Sprintf("validation for manifests with schemaVersion < 2.1.0 may not produce accurate results (version: %s)", p.schemaVersion),
			Category: "schema",
		})
	}
	return nil
}

// isSemverLessThan compares two semantic version strings
func isSemverLessThan(version1, version2 string) bool {
	return semver.Compare(
		semver.Canonical("v"+version1),
		semver.Canonical("v"+version2),
	) < 0
}

func (p *ValidationPipeline) validateSchema(result *ValidationResult) error {
	schemaData, err := getSchema(p.schemaVersion)
	if err != nil {
		return fmt.Errorf("error getting schema: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", bytes.NewReader(schemaData)); err != nil {
		return fmt.Errorf("error adding schema resource: %w", err)
	}

	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return fmt.Errorf("error compiling schema: %w", err)
	}

	var v any
	if err := json.Unmarshal(p.manifestData, &v); err != nil {
		return fmt.Errorf("error unmarshaling manifest for validation: %w", err)
	}

	if err := schema.Validate(v); err != nil {
		validationError := err.(*jsonschema.ValidationError)
		issues := p.extractValidationErrors(validationError)
		result.Errors = p.deduplicateErrors(issues)
		result.IsValid = false
	} else {
		result.IsValid = true
	}

	return nil
}

func (p *ValidationPipeline) extractValidationErrors(err *jsonschema.ValidationError) []ValidationError {
	var issues []ValidationError
	p.collectErrors(err, &issues)
	return issues
}

func (p *ValidationPipeline) collectErrors(err *jsonschema.ValidationError, issues *[]ValidationError) {
	for _, cause := range err.Causes {
		if len(cause.Causes) == 0 {
			issue := p.errorProcessor.ProcessValidationError(cause, cause.InstanceLocation)
			if issue.Message != "" {
				*issues = append(*issues, issue)
			}
		} else {
			p.collectErrors(cause, issues)
		}
	}
}

func (p *ValidationPipeline) deduplicateErrors(issues []ValidationError) []ValidationError {
	seen := make(map[string]bool)
	enumGroups := make(map[string][]string)
	missingPropertyContexts := make(map[string]bool)
	var deduplicated []ValidationError

	for _, issue := range issues {
		// Track contexts that have missing property errors
		if strings.Contains(issue.Message, "missing properties") {
			contextKey := fmt.Sprintf("%s::%s::%s::%s",
				issue.Module, issue.Declaration, issue.Member, issue.Property)
			missingPropertyContexts[contextKey] = true
		}

		// Handle enum validation errors specially
		if strings.Contains(issue.Message, "value must be") {
			contextKey := fmt.Sprintf("%s::%s::%s::%s",
				issue.Module, issue.Declaration, issue.Member, issue.Property)

			// Extract valid value from message
			if strings.Contains(issue.Message, "value must be \"") {
				start := strings.Index(issue.Message, "value must be \"") + len("value must be \"")
				end := strings.Index(issue.Message[start:], "\"")
				if end > 0 {
					validValue := issue.Message[start : start+end]
					enumGroups[contextKey] = append(enumGroups[contextKey], validValue)
				}
			}
			continue
		}

		// Clean up additionalProperties messages
		if strings.Contains(issue.Message, "additionalProperties") && strings.Contains(issue.Message, "not allowed") {
			start := strings.Index(issue.Message, "'")
			end := strings.LastIndex(issue.Message, "'")
			if start >= 0 && end > start {
				props := issue.Message[start+1 : end]
				issue.Message = fmt.Sprintf("property '%s' not allowed", props)
			}
		}

		// Deduplicate based on context and message
		key := fmt.Sprintf("%s::%s::%s::%s::%s",
			issue.Module, issue.Declaration, issue.Member, issue.Property, issue.Message)

		if !seen[key] {
			seen[key] = true
			deduplicated = append(deduplicated, issue)
		}
	}

	// Add consolidated enum errors, but only if there's no missing property error for the same context
	for contextKey, validValues := range enumGroups {
		// Skip enum errors if we already have a missing property error for this context
		if missingPropertyContexts[contextKey] {
			continue
		}

		parts := strings.Split(contextKey, "::")
		if len(parts) == 4 {
			// Deduplicate values
			uniqueValues := make(map[string]bool)
			for _, val := range validValues {
				uniqueValues[val] = true
			}

			var deduped []string
			for val := range uniqueValues {
				deduped = append(deduped, val)
			}

			// Determine message and ID
			isKindError := strings.Contains(contextKey, "kind")
			var message, id string

			if isKindError && (strings.Contains(parts[2], "invalid-") || strings.Contains(parts[2], "accessor")) {
				message = "invalid kind"
				id = "schema-invalid-kind"
			} else {
				message = fmt.Sprintf("invalid value, must be one of: %s", strings.Join(deduped, ", "))
				id = "schema-invalid-enum"
			}

			issue := ValidationError{
				ID:          id,
				Module:      parts[0],
				Declaration: parts[1],
				Member:      parts[2],
				Property:    parts[3],
				Message:     message,
			}
			deduplicated = append(deduplicated, issue)
		}
	}

	return deduplicated
}

func filterWarningsByConfig(warnings []ValidationWarning, disabledRules []string) []ValidationWarning {
	var filtered []ValidationWarning

	for _, warning := range warnings {
		if !isWarningDisabled(warning, disabledRules) {
			filtered = append(filtered, warning)
		}
	}

	return filtered
}

func isWarningDisabled(warning ValidationWarning, disabledRules []string) bool {
	// Also check Viper config for backwards compatibility
	viperDisabledRules := viper.GetStringSlice("warnings.disable")
	allDisabledRules := append(disabledRules, viperDisabledRules...)

	// Check if this specific warning ID or its category is disabled
	for _, rule := range allDisabledRules {
		if rule == warning.ID || rule == warning.Category {
			return true
		}
	}

	return false
}

func getSchema(version string) ([]byte, error) {
	// First try to use embedded schema
	if schemaData, err := getEmbeddedSchema(version); err == nil {
		return schemaData, nil
	}

	// If embedded schema loading fails, apply fallback logic
	if version == "2.1.0" || version == "2.1.1" {
		// Use speculative fallback for known problematic versions
		return getEmbeddedSchema("2.1.1-speculative")
	}

	// For other versions, try fetching from CDN
	return tryFetchSchema(version)
}

func getEmbeddedSchema(version string) ([]byte, error) {
	schemaPath := fmt.Sprintf("schemas/%s.json", version)
	return embeddedSchemas.ReadFile(schemaPath)
}

func tryFetchSchema(version string) ([]byte, error) {
	cacheDir, err := xdg.CacheFile(filepath.Join("cem", "schemas"))
	if err != nil {
		return nil, fmt.Errorf("could not get cache directory: %w", err)
	}

	schemaPath := filepath.Join(cacheDir, version+".json")

	if _, err := os.Stat(schemaPath); err == nil {
		return os.ReadFile(schemaPath)
	}

	url := fmt.Sprintf("https://unpkg.com/custom-elements-manifest@%s/schema.json", version)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch schema from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status fetching schema from %s: %s", url, resp.Status)
	}

	schemaData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read schema from response body: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(schemaPath), 0755); err != nil {
		return nil, fmt.Errorf("could not create cache directory: %w", err)
	}

	if err := os.WriteFile(schemaPath, schemaData, 0644); err != nil {
		return nil, fmt.Errorf("could not write schema to cache: %w", err)
	}

	return schemaData, nil
}
