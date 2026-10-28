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
	"io/fs"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/adrg/xdg"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/viper"
	"golang.org/x/mod/semver"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var englishPrinter = message.NewPrinter(language.English)

// Embedded schemas for all custom-elements-manifest versions
// Includes a speculative 2.1.1 schema to work around issues in 2.1.0
// See: https://github.com/webcomponents/custom-elements-manifest/issues/138
// See: https://github.com/vega/ts-json-schema-generator/pull/2323
//
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

	// NoNetwork prevents validation from fetching a schema that is neither
	// embedded nor cached. It is passed to Schema as WithoutNetwork.
	NoNetwork bool
}

// Validate validates a custom-elements.json manifest.
//
// Set ValidationOptions.NoNetwork to prevent schema fetches; embedded and
// cached schemas still validate. The filesystem argument must implement
// ValidateFS.
func Validate(fsys ValidateFS, manifestPath string, options ValidationOptions) (*ValidationResult, error) {
	manifestData, err := fsys.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error reading manifest file: %w", err)
	}

	pipeline, err := NewValidationPipeline(fsys, manifestData)
	if err != nil {
		return nil, err
	}

	return pipeline.Validate(options)
}

// ValidationPipeline orchestrates the validation process
type ValidationPipeline struct {
	fsys             ValidateFS
	manifestData     []byte
	schemaVersion    string
	navigator        *ManifestNavigator
	warningProcessor *WarningProcessor
	errorProcessor   *ErrorProcessor
}

// NewValidationPipeline creates a new validation pipeline
func NewValidationPipeline(fsys ValidateFS, manifestData []byte) (*ValidationPipeline, error) {
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
		fsys: fsys,
		// Clone so a caller reusing the read buffer cannot change the bytes
		// validated after the navigator was built from the original snapshot.
		manifestData:     slices.Clone(manifestData),
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
	if err := p.validateSchema(result, options); err != nil {
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
	// Validate schema version format. A value that is not a semantic version
	// gets a warning; validateSchema separately decides whether it is unsafe
	// to use as a path component.
	if !isSemverSchemaVersion(p.schemaVersion) {
		result.Warnings = append(result.Warnings, ValidationWarning{
			ID:       "schema-version-invalid",
			Message:  fmt.Sprintf("schemaVersion is not a valid semantic version: %s", p.schemaVersion),
			Category: "schema",
		})
		return nil
	}
	// Warn based on the schema that will actually validate the manifest, since
	// a shortened version such as "2" selects the newest embedded 2.x schema
	// rather than the exact 2.0.0 spelling.
	resolvedVersion := p.schemaVersion
	if embedded, ok := resolveEmbeddedSchemaVersion(p.schemaVersion); ok {
		resolvedVersion = embedded
	}
	// Only warn for versions < 2.1.0, since we handle 2.1.0 with our speculative schema
	if isSemverLessThan(resolvedVersion, "2.1.0") {
		result.Warnings = append(result.Warnings, ValidationWarning{
			ID:       "schema-version-old",
			Message:  fmt.Sprintf("validation for manifests with schemaVersion < 2.1.0 may not produce accurate results (version: %s)", p.schemaVersion),
			Category: "schema",
		})
	}
	return nil
}

// speculativeSchemaVersion is the embedded schema used for the known-broken
// 2.1.0 and 2.1.1 versions and for values that are unsafe as path components.
const speculativeSchemaVersion = "2.1.1-speculative"

// isPathSafeSchemaVersion reports whether version can be used as a path
// component or URL path segment. It permits dist-tags and semver ranges but
// rejects path separators, traversal segments, whitespace, and URL delimiters.
func isPathSafeSchemaVersion(version string) bool {
	if version == "" || strings.ContainsAny(version, `/\`) || strings.Contains(version, "..") {
		return false
	}
	for _, r := range version {
		if unicode.IsSpace(r) || unicode.IsControl(r) || strings.ContainsRune("?#%", r) {
			return false
		}
	}
	return true
}

// canonicalSchemaVersion returns the canonical semantic version for version,
// accepting an optional leading "v" and shortened spellings. It reports false
// when version is not a semantic version, such as a dist-tag or range.
func canonicalSchemaVersion(version string) (string, bool) {
	normalized := strings.TrimPrefix(version, "v")
	if !semver.IsValid("v" + normalized) {
		return "", false
	}
	return strings.TrimPrefix(semver.Canonical("v"+normalized), "v"), true
}

// isSemverSchemaVersion reports whether version is a semantic version with an
// optional leading "v".
func isSemverSchemaVersion(version string) bool {
	_, ok := canonicalSchemaVersion(version)
	return ok
}

// schemaVersionIsShort reports whether version is a shortened semantic
// version spelling with only a major, or a major and minor, component, such as
// "2" or "2.1". Short spellings are ranges: they select the newest embedded
// schema that matches the supplied components. Prerelease and build suffixes
// require a full major.minor.patch spelling.
func schemaVersionIsShort(version string) bool {
	normalized := strings.TrimPrefix(version, "v")
	if !semver.IsValid("v" + normalized) {
		return false
	}
	core := normalized
	if i := strings.IndexAny(core, "-+"); i >= 0 {
		core = core[:i]
	}
	return strings.Count(core, ".") < 2
}

// resolveEmbeddedSchemaVersion returns the embedded schema that version
// selects. Exact versions must match an embedded schema after
// canonicalization; shortened versions select the newest embedded schema with
// the same major, or major and minor, components. It reports false when no
// embedded schema matches, so the caller can fall back to the original
// spelling for cache and network resolution.
func resolveEmbeddedSchemaVersion(version string) (string, bool) {
	canonical, ok := canonicalSchemaVersion(version)
	if !ok {
		return "", false
	}
	if schemaVersionIsShort(version) {
		return newestEmbeddedSchemaVersion(strings.TrimPrefix(version, "v"))
	}
	// 2.1.0 and 2.1.1 are known-broken upstream; use the embedded stand-in.
	if canonical == "2.1.0" || canonical == "2.1.1" {
		return speculativeSchemaVersion, true
	}
	if _, err := getEmbeddedSchema(canonical); err == nil {
		return canonical, true
	}
	return "", false
}

// newestEmbeddedSchemaVersion returns the greatest embedded schema version
// that belongs to short, a major or major.minor spelling. The speculative
// schema participates like any other embedded version.
func newestEmbeddedSchemaVersion(short string) (string, bool) {
	names, err := fs.Glob(embeddedSchemas, "schemas/*.json")
	if err != nil {
		return "", false
	}
	best := ""
	for _, name := range names {
		candidate := strings.TrimSuffix(strings.TrimPrefix(name, "schemas/"), ".json")
		canonical, ok := canonicalSchemaVersion(candidate)
		if !ok || !schemaVersionHasPrefix(canonical, short) {
			continue
		}
		if best == "" || semver.Compare("v"+canonical, "v"+best) > 0 {
			best = canonical
		}
	}
	return best, best != ""
}

// schemaVersionHasPrefix reports whether canonical belongs to the range
// selected by short, a major or major.minor spelling.
func schemaVersionHasPrefix(canonical, short string) bool {
	canonical = strings.TrimPrefix(canonical, "v")
	short = strings.TrimPrefix(short, "v")
	if !strings.Contains(short, ".") {
		major, _, _ := strings.Cut(canonical, ".")
		return major == short
	}
	parts := strings.SplitN(canonical, ".", 3)
	if len(parts) < 2 {
		return false
	}
	return parts[0]+"."+parts[1] == short
}

// isSemverLessThan compares two semantic version strings
func isSemverLessThan(version1, version2 string) bool {
	canonical1, _ := canonicalSchemaVersion(version1)
	canonical2, _ := canonicalSchemaVersion(version2)
	return semver.Compare("v"+canonical1, "v"+canonical2) < 0
}

func (p *ValidationPipeline) validateSchema(result *ValidationResult, options ValidationOptions) error {
	// Hardening: never pass an untrusted schemaVersion into embedded-schema or
	// cache path construction. Invalid values keep their warning from
	// checkSchemaVersion and validation continues against an embedded schema.
	schemaVersion := p.schemaVersion
	unsafeSchemaVersion := !isPathSafeSchemaVersion(schemaVersion)
	if unsafeSchemaVersion {
		schemaVersion = speculativeSchemaVersion
	}
	schemaData, err := getSchema(p.fsys, schemaVersion, schemaOptions{noNetwork: options.NoNetwork})
	if err != nil {
		return fmt.Errorf("error getting schema: %w", err)
	}

	schemaDoc, err := jsonschema.UnmarshalJSON(bytes.NewReader(schemaData))
	if err != nil {
		return fmt.Errorf("error parsing schema: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schemaDoc); err != nil {
		return fmt.Errorf("error adding schema resource: %w", err)
	}

	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return fmt.Errorf("error compiling schema: %w", err)
	}

	v, err := jsonschema.UnmarshalJSON(bytes.NewReader(p.manifestData))
	if err != nil {
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

	// The manifest was checked against a fallback schema, not the schema its
	// schemaVersion names, so it must never be reported as valid. Record a real
	// error so the failure is meaningful to callers and displays, followed by
	// any errors from the fallback validation pass.
	if unsafeSchemaVersion {
		result.Errors = append([]ValidationError{{
			ID:      "schema-version-invalid",
			Message: fmt.Sprintf("schemaVersion %q is unsafe and was rejected; validation used the embedded %s schema instead", p.schemaVersion, speculativeSchemaVersion),
		}}, result.Errors...)
		result.IsValid = false
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
			location := "/" + strings.Join(cause.InstanceLocation, "/")
			message := cause.ErrorKind.LocalizedString(englishPrinter)
			issue := p.errorProcessor.ProcessValidationError(location, message)
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
		if strings.Contains(issue.Message, "missing propert") {
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

// SchemaOption configures Schema.
type SchemaOption func(*schemaOptions)

// schemaOptions is the resolved configuration for Schema.
type schemaOptions struct {
	noNetwork bool
}

// WithoutNetwork prevents Schema from fetching a schema that is neither
// embedded nor cached. When the version cannot be resolved offline, Schema
// returns an error naming the version instead of requesting it from unpkg.
//
// Embedded schemas and cached schemas still resolve normally.
func WithoutNetwork() SchemaOption {
	return func(options *schemaOptions) {
		options.noNetwork = true
	}
}

// Schema returns the JSON schema for a given version, using embedded schemas
// and falling back to fetching from external sources if needed.
//
// Options:
//
//   - WithoutNetwork: do not fetch from unpkg. Schema returns an error when
//     version is neither embedded nor cached. For example:
//
//     schema, err := Schema(fsys, "2.1.2", WithoutNetwork())
//
// fsys may implement only ReadFS. Caching a fetched schema additionally uses
// WriteFile and MkdirAll when fsys implements them, as ValidateFS does; a
// ReadFS-only implementation still returns the fetched schema.
//
// version must not contain path separators, traversal segments, whitespace, or
// URL delimiters, so caller-controlled input cannot escape the embedded
// schema, cache, or fetch paths. Semantic versions with a full
// major.minor.patch spelling, with an optional leading "v", are canonicalized
// and resolved against the embedded schemas first. Shortened spellings are
// ranges: a bare major such as "2" selects the newest embedded schema with
// that major, and a major.minor spelling such as "2.1" selects the newest
// embedded schema with those components. Other path-safe values, such as
// dist-tags, and exact versions with no embedded schema fall through to the
// cache and fetch using the spelling supplied by the caller.
//
// For versions 2.1.0 and 2.1.1, Schema returns the embedded
// 2.1.1-speculative schema instead of the official upstream schema, because
// the upstream 2.1.0 schema is known to be broken. The speculative schema is
// stricter than the official schema: it rejects additional properties.
// Callers that require strict validation against the official 2.1.0 schema
// should not use Schema for those versions.
func Schema(fsys ReadFS, version string, opts ...SchemaOption) ([]byte, error) {
	var options schemaOptions
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	return getSchema(fsys, version, options)
}

func getSchema(fsys ReadFS, version string, options schemaOptions) ([]byte, error) {
	// Hardening: Schema can be called directly with caller-controlled input.
	// Reject values that could escape a path when used as a cache filename or
	// URL segment.
	if !isPathSafeSchemaVersion(version) {
		return nil, fmt.Errorf("invalid schema version %q: must not contain path separators, traversal segments, or URL delimiters", version)
	}

	// Resolve against the embedded schemas first. Canonicalizing valid
	// semantic versions keeps spellings such as "v2.1.0+build" and shortened
	// spellings such as "2.1" resolving to the same embedded schema.
	if embedded, ok := resolveEmbeddedSchemaVersion(version); ok {
		return getEmbeddedSchema(embedded)
	}

	// Nothing embedded matches. Resolve the caller's original spelling: a
	// shortened version such as "0.0" is a range, and a dist-tag is resolved
	// by the cache and npm, so substituting a canonical exact version would
	// change which schema they return.
	return tryFetchSchema(fsys, version, options)
}

func getEmbeddedSchema(version string) ([]byte, error) {
	schemaPath := fmt.Sprintf("schemas/%s.json", version)
	return embeddedSchemas.ReadFile(schemaPath)
}

// schemaCacheDir returns the directory used to cache fetched schemas, or an
// empty string when no cache directory is available (for example, in wasm).
//
// It derives the path from xdg.CacheHome instead of calling xdg.CacheFile:
// xdg.CacheFile creates the directory on the host filesystem, which would
// escape a caller-supplied in-memory or sandboxed FileSystem. All cache reads
// and writes go through fsys.
//
// TODO(#371): wasm targets may want their own cache backed by IndexedDB or similar
func schemaCacheDir() string {
	if xdg.CacheHome == "" {
		return ""
	}
	return filepath.Join(xdg.CacheHome, "cem", "schemas")
}

// schemaCacheWriter is the optional write half of the schema cache. ReadFS
// does not require these methods, so a read-only implementation simply skips
// caching.
type schemaCacheWriter interface {
	WriteFile(name string, data []byte, perm fs.FileMode) error
	MkdirAll(path string, perm fs.FileMode) error
}

// httpGet is the HTTP seam used to fetch schemas. Tests replace it so they
// never reach the network.
var httpGet = http.Get

func tryFetchSchema(fsys ReadFS, version string, options schemaOptions) ([]byte, error) {
	cacheDir := schemaCacheDir()
	if cacheDir != "" {
		schemaPath := filepath.Join(cacheDir, version+".json")
		if _, err := fsys.Stat(schemaPath); err == nil {
			return fsys.ReadFile(schemaPath)
		}
	}

	if options.noNetwork {
		return nil, fmt.Errorf("schema for version %q is not embedded or cached and network access is disabled", version)
	}

	url := fmt.Sprintf("https://unpkg.com/custom-elements-manifest@%s/schema.json", version)
	resp, err := httpGet(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch schema from %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status fetching schema from %s: %s", url, resp.Status)
	}

	schemaData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read schema from response body: %w", err)
	}

	// Cache the result when a cache directory is available and the supplied
	// filesystem can write. A ReadFS-only implementation skips this step.
	if cacheDir != "" {
		if writer, ok := fsys.(schemaCacheWriter); ok {
			schemaPath := filepath.Join(cacheDir, version+".json")
			_ = writer.MkdirAll(cacheDir, 0755)
			_ = writer.WriteFile(schemaPath, schemaData, 0644)
		}
	}

	return schemaData, nil
}
