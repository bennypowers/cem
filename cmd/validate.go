package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"bennypowers.dev/cem/workspace"
	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang.org/x/mod/semver"
)

func init() {
	validateCmd.Flags().BoolP("verbose", "v", false, "Show detailed information including schema version")
	validateCmd.Flags().Bool("no-warnings", false, "Suppress validation warnings")
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a custom-elements.json manifest",
	Long:  `Validate a custom-elements.json manifest against its JSON schema.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting workspace context: %v\n", err)
			os.Exit(1)
		}

		manifestPath := ctx.CustomElementsManifestPath()
		if len(args) > 0 {
			manifestPath = args[0]
		}

		if manifestPath == "" {
			fmt.Fprintln(os.Stderr, "Could not find custom-elements.json")
			os.Exit(1)
		}

		manifestData, err := os.ReadFile(manifestPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading manifest file: %v\n", err)
			os.Exit(1)
		}

		var manifest struct {
			SchemaVersion string `json:"schemaVersion"`
		}
		if err := json.Unmarshal(manifestData, &manifest); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing manifest to find schemaVersion: %v\n", err)
			os.Exit(1)
		}

		if manifest.SchemaVersion == "" {
			fmt.Fprintln(os.Stderr, "Error: schemaVersion not found in manifest.")
			os.Exit(1)
		}

		if semver.Compare(
			semver.Canonical("v"+manifest.SchemaVersion),
			semver.Canonical("v2.1.1"),
		) < 0 {
			pterm.Warning.Printfln("validation for manifests with schemaVersion <= 2.1.0 may not produce accurate results (version: %s)", manifest.SchemaVersion)
		}

		schemaData, err := getSchema(manifest.SchemaVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting schema: %v\n", err)
			os.Exit(1)
		}

		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource("schema.json", bytes.NewReader(schemaData)); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding schema resource: %v\n", err)
			os.Exit(1)
		}
		schema, err := compiler.Compile("schema.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error compiling schema: %v\n", err)
			os.Exit(1)
		}

		var v any
		if err := json.Unmarshal(manifestData, &v); err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling manifest for validation: %v\n", err)
			os.Exit(1)
		}

		verbose, _ := cmd.Flags().GetBool("verbose")
		noWarnings, _ := cmd.Flags().GetBool("no-warnings")
		
		if err := schema.Validate(v); err != nil {
			validationError := err.(*jsonschema.ValidationError)
			printValidationResults(manifestPath, manifest.SchemaVersion, validationError, verbose)
			os.Exit(1)
		}

		// Check for warnings on valid manifests
		if !noWarnings {
			warnings := detectWarnings(manifestData)
			filteredWarnings := filterWarningsByConfig(warnings)
			if len(filteredWarnings) > 0 {
				printValidationWarnings(manifestPath, filteredWarnings, verbose)
			} else {
				printValidationSuccess(manifestPath, manifest.SchemaVersion, verbose)
			}
		} else {
			printValidationSuccess(manifestPath, manifest.SchemaVersion, verbose)
		}
	},
}

func printValidationSuccess(manifestPath, schemaVersion string, verbose bool) {
	pterm.Success.Printf("✓ Manifest is valid (%s)\n", manifestPath)
	if verbose {
		pterm.Info.Printf("  Schema version: %s\n", schemaVersion)
	}
}

func printValidationResults(manifestPath, schemaVersion string, err *jsonschema.ValidationError, verbose bool) {
	// Read manifest for context extraction
	manifestData, _ := os.ReadFile(manifestPath)
	var manifestJSON map[string]interface{}
	json.Unmarshal(manifestData, &manifestJSON)
	
	// Group and process validation errors
	issues := extractValidationIssues(err, manifestJSON)
	deduplicatedIssues := deduplicateIssues(issues)
	groupedIssues := groupIssuesByContext(deduplicatedIssues)
	
	// Always use consistent format
	pterm.Error.Printf("✗ Validation failed with %d issue%s (%s)\n", len(deduplicatedIssues), func() string {
		if len(deduplicatedIssues) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	pterm.Println()
	
	printGroupedIssues(groupedIssues)
	
	if verbose {
		pterm.Info.Printf("Schema version: %s\n", schemaVersion)
	}
}

type ValidationIssue struct {
	Module       string
	Declaration  string
	Member       string
	Property     string
	Message      string
	Location     string
	Index        int // For array items
}

type GroupedIssues struct {
	Module       string
	Declaration  string
	Issues       []ValidationIssue
}

type Warning struct {
	ID           string // Unique identifier for this warning rule
	Module       string
	Declaration  string
	Member       string
	Property     string
	Message      string
	Category     string // "lifecycle", "private", "verbose", etc.
}

type GroupedWarnings struct {
	Module       string
	Declaration  string
	Warnings     []Warning
}

func extractValidationIssues(err *jsonschema.ValidationError, manifest map[string]interface{}) []ValidationIssue {
	var issues []ValidationIssue
	collectIssues(err, manifest, &issues)
	return issues
}

func collectIssues(err *jsonschema.ValidationError, manifest map[string]interface{}, issues *[]ValidationIssue) {
	for _, cause := range err.Causes {
		if len(cause.Causes) == 0 {
			issue := parseValidationIssue(cause, manifest)
			if issue.Message != "" {
				*issues = append(*issues, issue)
			}
		} else {
			collectIssues(cause, manifest, issues)
		}
	}
}

func parseValidationIssue(cause *jsonschema.ValidationError, manifest map[string]interface{}) ValidationIssue {
	location := cause.InstanceLocation
	if location == "" {
		location = "root"
	}
	
	issue := ValidationIssue{
		Message:  cause.Message,
		Location: location,
		Index:    -1, // Default to -1 for non-array items
	}
	
	// Parse JSON path to extract meaningful context
	parts := strings.Split(strings.TrimPrefix(location, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return issue
	}
	
	// Navigate through the JSON structure to build context
	for i, part := range parts {
		// Handle array indices
		if matched, _ := regexp.MatchString(`^\d+$`, part); matched {
			index, _ := strconv.Atoi(part)
			
			if i > 0 {
				parentKey := parts[i-1]
				switch parentKey {
				case "modules":
					// Extract module path
					if modules, ok := manifest["modules"].([]interface{}); ok && index < len(modules) {
						if module, ok := modules[index].(map[string]interface{}); ok {
							if path, ok := module["path"].(string); ok {
								issue.Module = path
							}
						}
					}
				case "declarations":
					// Extract declaration info - need to find the right module first
					moduleIndex := -1
					for j := i - 2; j >= 0; j-- {
						if parts[j] == "modules" && j+1 < len(parts) {
							if matched, _ := regexp.MatchString(`^\d+$`, parts[j+1]); matched {
								moduleIndex, _ = strconv.Atoi(parts[j+1])
								break
							}
						}
					}
					if moduleIndex >= 0 {
						if modules, ok := manifest["modules"].([]interface{}); ok && moduleIndex < len(modules) {
							if module, ok := modules[moduleIndex].(map[string]interface{}); ok {
								if declarations, ok := module["declarations"].([]interface{}); ok && index < len(declarations) {
									if decl, ok := declarations[index].(map[string]interface{}); ok {
										if name, ok := decl["name"].(string); ok {
											if kind, ok := decl["kind"].(string); ok {
												issue.Declaration = fmt.Sprintf("%s %s", kind, name)
											} else {
												issue.Declaration = name
											}
										} else {
											if kind, ok := decl["kind"].(string); ok {
												issue.Declaration = fmt.Sprintf("%s[%d]", kind, index)
											} else {
												issue.Declaration = fmt.Sprintf("declaration[%d]", index)
											}
										}
									}
								}
							}
						}
					}
				case "members":
					// Extract member info - need to find declaration
					moduleIndex, declIndex := -1, -1
					for j := i - 2; j >= 0; j-- {
						if parts[j] == "declarations" && j+1 < len(parts) {
							if matched, _ := regexp.MatchString(`^\d+$`, parts[j+1]); matched {
								declIndex, _ = strconv.Atoi(parts[j+1])
								break
							}
						}
					}
					for j := i - 4; j >= 0; j-- {
						if parts[j] == "modules" && j+1 < len(parts) {
							if matched, _ := regexp.MatchString(`^\d+$`, parts[j+1]); matched {
								moduleIndex, _ = strconv.Atoi(parts[j+1])
								break
							}
						}
					}
					if moduleIndex >= 0 && declIndex >= 0 {
						if modules, ok := manifest["modules"].([]interface{}); ok && moduleIndex < len(modules) {
							if module, ok := modules[moduleIndex].(map[string]interface{}); ok {
								if declarations, ok := module["declarations"].([]interface{}); ok && declIndex < len(declarations) {
									if decl, ok := declarations[declIndex].(map[string]interface{}); ok {
										if members, ok := decl["members"].([]interface{}); ok && index < len(members) {
											issue.Index = index
											if member, ok := members[index].(map[string]interface{}); ok {
												if name, ok := member["name"].(string); ok {
													if kind, ok := member["kind"].(string); ok {
														issue.Member = fmt.Sprintf("%s %s", kind, name)
													} else {
														issue.Member = name
													}
												} else {
													// For members without names, show the index and any available kind
													if kind, ok := member["kind"].(string); ok {
														issue.Member = fmt.Sprintf("%s[%d]", kind, index)
													} else {
														issue.Member = fmt.Sprintf("member[%d]", index)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case "attributes", "events", "slots", "cssProperties", "cssParts", "cssStates":
					// Extract property info
					singularName := strings.TrimSuffix(parentKey, "s")
					if parentKey == "cssProperties" {
						singularName = "CSS property"
					} else if parentKey == "cssParts" {
						singularName = "CSS part"
					} else if parentKey == "cssStates" {
						singularName = "CSS state"
					}
					
					// Find the declaration this property belongs to
					moduleIndex, declIndex := -1, -1
					for j := i - 2; j >= 0; j-- {
						if parts[j] == "declarations" && j+1 < len(parts) {
							if matched, _ := regexp.MatchString(`^\d+$`, parts[j+1]); matched {
								declIndex, _ = strconv.Atoi(parts[j+1])
								break
							}
						}
					}
					for j := i - 4; j >= 0; j-- {
						if parts[j] == "modules" && j+1 < len(parts) {
							if matched, _ := regexp.MatchString(`^\d+$`, parts[j+1]); matched {
								moduleIndex, _ = strconv.Atoi(parts[j+1])
								break
							}
						}
					}
					if moduleIndex >= 0 && declIndex >= 0 {
						if modules, ok := manifest["modules"].([]interface{}); ok && moduleIndex < len(modules) {
							if module, ok := modules[moduleIndex].(map[string]interface{}); ok {
								if declarations, ok := module["declarations"].([]interface{}); ok && declIndex < len(declarations) {
									if decl, ok := declarations[declIndex].(map[string]interface{}); ok {
										if arr, ok := decl[parentKey].([]interface{}); ok && index < len(arr) {
											issue.Index = index
											if item, ok := arr[index].(map[string]interface{}); ok {
												if name, ok := item["name"].(string); ok {
													issue.Property = fmt.Sprintf("%s %s", singularName, name)
												} else {
													issue.Property = fmt.Sprintf("%s[%d]", singularName, index)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	
	return issue
}

func deduplicateIssues(issues []ValidationIssue) []ValidationIssue {
	seen := make(map[string]bool)
	enumGroups := make(map[string][]string)
	var deduplicated []ValidationIssue
	
	for _, issue := range issues {
		// For enum validation errors, collect all valid values
		if strings.Contains(issue.Message, "value must be") {
			// Create a key that groups enum errors by context
			contextKey := fmt.Sprintf("%s::%s::%s::%s", 
				issue.Module, issue.Declaration, issue.Member, issue.Property)
			
			// Extract the valid value from the message
			if strings.Contains(issue.Message, "value must be \"") {
				start := strings.Index(issue.Message, "value must be \"") + len("value must be \"")
				end := strings.Index(issue.Message[start:], "\"")
				if end > 0 {
					validValue := issue.Message[start : start+end]
					enumGroups[contextKey] = append(enumGroups[contextKey], validValue)
				}
			}
			
			// Skip adding individual enum errors for now
			continue
		} else {
			// Clean up additionalProperties messages
			if strings.Contains(issue.Message, "additionalProperties") {
				// Extract the property names from the message
				if strings.Contains(issue.Message, "not allowed") {
					start := strings.Index(issue.Message, "'")
					end := strings.LastIndex(issue.Message, "'")
					if start >= 0 && end > start {
						props := issue.Message[start+1 : end]
						issue.Message = fmt.Sprintf("property '%s' not allowed", props)
					}
				}
			}
			
			// Create a key that represents the unique context and message
			key := fmt.Sprintf("%s::%s::%s::%s::%s", 
				issue.Module, issue.Declaration, issue.Member, issue.Property, issue.Message)
			
			if !seen[key] {
				seen[key] = true
				deduplicated = append(deduplicated, issue)
			}
		}
	}
	
	// Now add consolidated enum errors
	for contextKey, validValues := range enumGroups {
		parts := strings.Split(contextKey, "::")
		if len(parts) == 4 {
			// Deduplicate valid values using a map
			uniqueValues := make(map[string]bool)
			for _, val := range validValues {
				uniqueValues[val] = true
			}
			
			var deduped []string
			for val := range uniqueValues {
				deduped = append(deduped, val)
			}
			
			// Check if this is a "kind" field with unknown values
			memberPart := parts[2]
			isKindError := strings.Contains(contextKey, "kind")
			
			var message string
			if isKindError && (strings.Contains(memberPart, "invalid-") || strings.Contains(memberPart, "accessor")) {
				// For unknown kinds, just say it's invalid
				message = "invalid kind"
			} else {
				// For valid enum violations, show the options
				message = fmt.Sprintf("invalid value, must be one of: %s", strings.Join(deduped, ", "))
			}
			
			issue := ValidationIssue{
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

func groupIssuesByContext(issues []ValidationIssue) []GroupedIssues {
	contextMap := make(map[string][]ValidationIssue)
	
	for _, issue := range issues {
		key := fmt.Sprintf("%s::%s", issue.Module, issue.Declaration)
		contextMap[key] = append(contextMap[key], issue)
	}
	
	var grouped []GroupedIssues
	for key, issueList := range contextMap {
		parts := strings.Split(key, "::")
		group := GroupedIssues{
			Module:      parts[0],
			Declaration: parts[1],
			Issues:      issueList,
		}
		grouped = append(grouped, group)
	}
	
	return grouped
}

func detectWarnings(manifestData []byte) []Warning {
	var manifest map[string]interface{}
	json.Unmarshal(manifestData, &manifest)
	
	var warnings []Warning
	
	if modules, ok := manifest["modules"].([]interface{}); ok {
		for _, moduleInterface := range modules {
			if module, ok := moduleInterface.(map[string]interface{}); ok {
				modulePath := ""
				if path, ok := module["path"].(string); ok {
					modulePath = path
				}
				
				if declarations, ok := module["declarations"].([]interface{}); ok {
					for _, declInterface := range declarations {
						if decl, ok := declInterface.(map[string]interface{}); ok {
							declName := ""
							declKind := ""
							if name, ok := decl["name"].(string); ok {
								declName = name
							}
							if kind, ok := decl["kind"].(string); ok {
								declKind = kind
							}
							
							// Check superclass warnings
							if superclass, ok := decl["superclass"].(map[string]interface{}); ok {
								superclassWarnings := checkSuperclassWarnings(superclass, modulePath, declName, declKind)
								warnings = append(warnings, superclassWarnings...)
							}
							
							// Check members for warnings
							if members, ok := decl["members"].([]interface{}); ok {
								// Check if this extends LitElement
								isLitElement := false
								if superclass, ok := decl["superclass"].(map[string]interface{}); ok {
									if name, ok := superclass["name"].(string); ok {
										isLitElement = strings.Contains(name, "LitElement")
									}
								}
								
								for memberIndex, memberInterface := range members {
									if member, ok := memberInterface.(map[string]interface{}); ok {
										memberWarnings := checkMemberWarnings(member, modulePath, declName, declKind, memberIndex, isLitElement)
										warnings = append(warnings, memberWarnings...)
									}
								}
							}
							
							// Check CSS properties for large defaults
							if cssProps, ok := decl["cssProperties"].([]interface{}); ok {
								for propIndex, propInterface := range cssProps {
									if prop, ok := propInterface.(map[string]interface{}); ok {
										propWarnings := checkCSSPropertyWarnings(prop, modulePath, declName, propIndex)
										warnings = append(warnings, propWarnings...)
									}
								}
							}
							
							// Check for implementation detail fields that shouldn't be documented
							implWarnings := checkImplementationDetailWarnings(decl, modulePath, declName, declKind)
							warnings = append(warnings, implWarnings...)
						}
					}
				}
			}
		}
	}
	
	return warnings
}

func checkImplementationDetailWarnings(decl map[string]interface{}, modulePath, declName, declKind string) []Warning {
	var warnings []Warning
	
	// Check static fields and methods that are implementation details
	if members, ok := decl["members"].([]interface{}); ok {
		for _, memberInterface := range members {
			if member, ok := memberInterface.(map[string]interface{}); ok {
				if memberName, nameOk := member["name"].(string); nameOk {
					if memberKind, kindOk := member["kind"].(string); kindOk {
						// Check for static fields that are implementation details
						if memberKind == "field" {
							if isStatic, staticOk := member["static"].(bool); staticOk && isStatic {
								// List of static fields that shouldn't be documented
								implementationFields := map[string]string{
									"styles":             "static styles field is implementation detail, should not be documented in public API",
									"shadowRootOptions":  "shadowRootOptions is implementation detail, should not be documented in public API", 
									"formAssociated":     "formAssociated is implementation detail, should not be documented in public API",
									"observedAttributes": "observedAttributes is implementation detail, should not be documented in public API",
								}
								
								if message, isImplementationDetail := implementationFields[memberName]; isImplementationDetail {
									var id string
									switch memberName {
									case "styles":
										id = "implementation-static-styles"
									case "shadowRootOptions":
										id = "implementation-shadow-root-options"
									case "formAssociated":
										id = "implementation-form-associated"
									case "observedAttributes":
										id = "implementation-observed-attributes"
									}
									
									warnings = append(warnings, Warning{
										ID:          id,
										Module:      modulePath,
										Declaration: fmt.Sprintf("%s %s", declKind, declName),
										Member:      fmt.Sprintf("static field %s", memberName),
										Message:     message,
										Category:    "implementation",
									})
								}
							}
						}
						
						// Check for form*Callback methods (form-associated element lifecycle)
						if memberKind == "method" {
							formCallbacks := []string{
								"formAssociatedCallback", "formDisabledCallback", 
								"formResetCallback", "formStateRestoreCallback",
							}
							for _, callback := range formCallbacks {
								if memberName == callback {
									warnings = append(warnings, Warning{
										ID:          "lifecycle-form-callbacks",
										Module:      modulePath,
										Declaration: fmt.Sprintf("%s %s", declKind, declName),
										Member:      fmt.Sprintf("method %s", memberName),
										Message:     "form lifecycle callback should not be documented in public API",
										Category:    "lifecycle",
									})
									break
								}
							}
						}
					}
				}
			}
		}
	}
	
	return warnings
}

func filterWarningsByConfig(warnings []Warning) []Warning {
	var filtered []Warning
	
	for _, warning := range warnings {
		if !isWarningDisabled(warning) {
			filtered = append(filtered, warning)
		}
	}
	
	return filtered
}

func isWarningDisabled(warning Warning) bool {
	// Check global rule disabling
	disabledRules := viper.GetStringSlice("warnings.disable")
	
	// Check if this specific warning ID or its category is disabled
	for _, rule := range disabledRules {
		// Check for exact ID match first
		if rule == warning.ID {
			return true
		}
		
		// Check for category match
		if rule == warning.Category {
			return true
		}
	}
	
	return false
}

func checkSuperclassWarnings(superclass map[string]interface{}, modulePath, declName, declKind string) []Warning {
	var warnings []Warning
	
	name := ""
	if superclassName, ok := superclass["name"].(string); ok {
		name = superclassName
	}
	
	module := ""
	if superclassModule, ok := superclass["module"].(string); ok {
		module = superclassModule
	}
	
	// List of built-in types that should have "module": "global:"
	builtInTypes := []string{
		// DOM APIs
		"HTMLElement", "Element", "Node", "EventTarget", "Document", "Window",
		// Events
		"Event", "CustomEvent", "MouseEvent", "KeyboardEvent", "FocusEvent", "TouchEvent",
		// Other Web APIs
		"AbortController", "AbortSignal", "Blob", "File", "FormData", "Headers",
		"Request", "Response", "URL", "URLSearchParams", "WebSocket",
		// JavaScript built-ins
		"Object", "Array", "Map", "Set", "WeakMap", "WeakSet", "Promise",
		"Error", "TypeError", "ReferenceError", "SyntaxError",
	}
	
	for _, builtIn := range builtInTypes {
		if name == builtIn && module != "global:" {
			var moduleInfo string
			if module == "" {
				moduleInfo = "missing module field"
			} else {
				moduleInfo = fmt.Sprintf("module is %q", module)
			}
			
			warnings = append(warnings, Warning{
				ID:          "superclass-builtin-modules",
				Module:      modulePath,
				Declaration: fmt.Sprintf("%s %s", declKind, declName),
				Message:     fmt.Sprintf("superclass %s is a built-in type but %s, should be \"module\": \"global:\"", name, moduleInfo),
				Category:    "superclass",
			})
			break
		}
	}
	
	return warnings
}

func checkMemberWarnings(member map[string]interface{}, modulePath, declName, declKind string, memberIndex int, isLitElement bool) []Warning {
	var warnings []Warning
	
	memberName := ""
	if name, ok := member["name"].(string); ok {
		memberName = name
	}
	
	memberKind := ""
	if kind, ok := member["kind"].(string); ok {
		memberKind = kind
	}
	
	// Skip if not a method
	if memberKind != "method" {
		return warnings
	}
	
	// Always warn about these lifecycle methods
	alwaysPrivateLifecycleMethods := []string{
		// HTMLElement / Web Components
		"connectedCallback", "disconnectedCallback", "attributeChangedCallback", "adoptedCallback",
		// Lit Element lifecycle (excluding render which depends on context)
		"firstUpdated", "updated", "willUpdate", "getUpdateComplete", "performUpdate",
		"scheduleUpdate", "requestUpdate", "createRenderRoot",
		// Common framework methods that shouldn't be public
		"constructor",
	}
	
	for _, lifecycle := range alwaysPrivateLifecycleMethods {
		if memberName == lifecycle {
			var id string
			// Determine specific ID based on the method
			switch memberName {
			case "connectedCallback", "disconnectedCallback", "attributeChangedCallback", "adoptedCallback":
				id = "lifecycle-web-components"
			case "firstUpdated", "updated", "willUpdate", "getUpdateComplete", "performUpdate", "scheduleUpdate", "requestUpdate", "createRenderRoot":
				id = "lifecycle-lit-methods"
			case "constructor":
				id = "lifecycle-constructor"
			default:
				id = "lifecycle-other"
			}
			
			warnings = append(warnings, Warning{
				ID:          id,
				Module:      modulePath,
				Declaration: fmt.Sprintf("%s %s", declKind, declName),
				Member:      fmt.Sprintf("%s %s", memberKind, memberName),
				Message:     "lifecycle method should not be documented in public API",
				Category:    "lifecycle",
			})
			break
		}
	}
	
	// Special case: warn about render method only in Lit elements
	if memberName == "render" && isLitElement {
		warnings = append(warnings, Warning{
			ID:          "lifecycle-lit-render",
			Module:      modulePath,
			Declaration: fmt.Sprintf("%s %s", declKind, declName),
			Member:      fmt.Sprintf("%s %s", memberKind, memberName),
			Message:     "render method in Lit element should not be documented in public API",
			Category:    "lifecycle",
		})
	}
	
	// Check for private methods (starting with _ or #)
	if strings.HasPrefix(memberName, "_") || strings.HasPrefix(memberName, "#") {
		var id string
		if strings.HasPrefix(memberName, "_") {
			id = "private-underscore-methods"
		} else {
			id = "private-hash-methods"
		}
		
		warnings = append(warnings, Warning{
			ID:          id,
			Module:      modulePath,
			Declaration: fmt.Sprintf("%s %s", declKind, declName),
			Member:      fmt.Sprintf("%s %s", memberKind, memberName),
			Message:     "private method should not be documented in public API",
			Category:    "private",
		})
	}
	
	// Check for internal/debug methods
	internalMethods := []string{"init", "destroy", "dispose", "cleanup", "debug", "log"}
	for _, internal := range internalMethods {
		if memberName == internal {
			warnings = append(warnings, Warning{
				ID:          "internal-utility-methods",
				Module:      modulePath,
				Declaration: fmt.Sprintf("%s %s", declKind, declName),
				Member:      fmt.Sprintf("%s %s", memberKind, memberName),
				Message:     "internal/utility method may not belong in public API",
				Category:    "internal",
			})
			break
		}
	}
	
	return warnings
}

func checkCSSPropertyWarnings(prop map[string]interface{}, modulePath, declName string, propIndex int) []Warning {
	var warnings []Warning
	
	propName := ""
	if name, ok := prop["name"].(string); ok {
		propName = name
	}
	
	// Check for overly long default values (likely CSS template literals)
	if defaultVal, ok := prop["default"].(string); ok {
		if len(defaultVal) > 200 {
			warnings = append(warnings, Warning{
				ID:          "verbose-css-defaults",
				Module:      modulePath,
				Declaration: fmt.Sprintf("class %s", declName),
				Property:    fmt.Sprintf("CSS property %s", propName),
				Message:     "default value is very long, consider using external CSS file",
				Category:    "verbose",
			})
		}
	}
	
	return warnings
}

func printValidationWarnings(manifestPath string, warnings []Warning, verbose bool) {
	groupedWarnings := groupWarningsByContext(warnings)
	
	pterm.Warning.Printf("⚠ Manifest valid with %d warning%s (%s)\n", len(warnings), func() string {
		if len(warnings) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	pterm.Println()
	
	printGroupedWarnings(groupedWarnings)
	
	if verbose {
		pterm.Info.Println("Use --no-warnings to suppress warnings")
	}
}

func groupWarningsByContext(warnings []Warning) []GroupedWarnings {
	contextMap := make(map[string][]Warning)
	
	for _, warning := range warnings {
		key := fmt.Sprintf("%s::%s", warning.Module, warning.Declaration)
		contextMap[key] = append(contextMap[key], warning)
	}
	
	var grouped []GroupedWarnings
	for key, warningList := range contextMap {
		parts := strings.Split(key, "::")
		group := GroupedWarnings{
			Module:      parts[0],
			Declaration: parts[1],
			Warnings:    warningList,
		}
		grouped = append(grouped, group)
	}
	
	return grouped
}

func printGroupedWarnings(groupedWarnings []GroupedWarnings) {
	for _, group := range groupedWarnings {
		// Print group header
		if group.Module != "" && group.Declaration != "" {
			pterm.Printf("%s %s %s %s\n", 
				pterm.LightCyan("📁"), 
				pterm.FgLightBlue.Sprint(group.Module),
				pterm.FgGray.Sprint("→"),
				pterm.FgYellow.Sprint(group.Declaration))
		} else if group.Module != "" {
			pterm.Printf("%s %s\n", 
				pterm.LightCyan("📁"), 
				pterm.FgLightBlue.Sprint(group.Module))
		}
		
		// Print warnings in this group
		for _, warning := range group.Warnings {
			if warning.Member != "" && warning.Property != "" {
				pterm.Printf("  %s %s → %s: %s\n", 
					pterm.LightYellow("⚠"), 
					warning.Member,
					warning.Property,
					warning.Message)
			} else if warning.Member != "" {
				pterm.Printf("  %s %s: %s\n", 
					pterm.LightYellow("⚠"), 
					warning.Member,
					warning.Message)
			} else if warning.Property != "" {
				pterm.Printf("  %s %s: %s\n", 
					pterm.LightYellow("⚠"), 
					warning.Property,
					warning.Message)
			} else {
				pterm.Printf("  %s %s\n", 
					pterm.LightYellow("⚠"), 
					warning.Message)
			}
		}
		pterm.Println()
	}
}

func printGroupedIssues(groupedIssues []GroupedIssues) {
	for _, group := range groupedIssues {
		// Print group header
		if group.Module != "" && group.Declaration != "" {
			pterm.Printf("%s %s %s %s\n", 
				pterm.LightCyan("📁"), 
				pterm.FgLightBlue.Sprint(group.Module),
				pterm.FgGray.Sprint("→"),
				pterm.FgYellow.Sprint(group.Declaration))
		} else if group.Module != "" {
			pterm.Printf("%s %s\n", 
				pterm.LightCyan("📁"), 
				pterm.FgLightBlue.Sprint(group.Module))
		} else {
			pterm.Printf("%s %s\n", 
				pterm.LightRed("⚠"), 
				pterm.FgYellow.Sprint("Root level issues"))
		}
		
		// Print issues in this group
		for _, issue := range group.Issues {
			if issue.Member != "" && issue.Property != "" {
				pterm.Printf("  %s %s → %s: %s\n", 
					pterm.LightRed("●"), 
					issue.Member,
					issue.Property,
					issue.Message)
			} else if issue.Member != "" {
				pterm.Printf("  %s %s: %s\n", 
					pterm.LightRed("●"), 
					issue.Member,
					issue.Message)
			} else if issue.Property != "" {
				pterm.Printf("  %s %s: %s\n", 
					pterm.LightRed("●"), 
					issue.Property,
					issue.Message)
			} else {
				pterm.Printf("  %s %s\n", 
					pterm.LightRed("●"), 
					issue.Message)
			}
		}
		pterm.Println()
	}
}


func getSchema(version string) ([]byte, error) {
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
