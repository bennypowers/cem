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
	"regexp"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type ValidationError struct {
	ID          string `json:"id"` // Unique identifier for this error type
	Module      string `json:"module,omitempty"`
	Declaration string `json:"declaration,omitempty"`
	Member      string `json:"member,omitempty"`
	Property    string `json:"property,omitempty"`
	Message     string `json:"message"`
	Location    string `json:"location,omitempty"`
	Index       int    `json:"index,omitempty"` // For array items
}

// ErrorIDRegistry provides centralized error ID assignment
type ErrorIDRegistry struct {
	patterns map[*regexp.Regexp]string
}

// NewErrorIDRegistry creates a new error ID registry
func NewErrorIDRegistry() *ErrorIDRegistry {
	registry := &ErrorIDRegistry{
		patterns: make(map[*regexp.Regexp]string),
	}

	// Register error patterns
	registry.registerPattern(`required property`, "schema-required-property")
	registry.registerPattern(`additionalProperties`, "schema-additional-properties")
	registry.registerPattern(`value must be`, "schema-invalid-enum")
	registry.registerPattern(`type`, "schema-invalid-type")
	registry.registerPattern(`format`, "schema-invalid-format")
	registry.registerPattern(`pattern`, "schema-invalid-pattern")
	registry.registerPattern(`minimum`, "schema-value-too-small")
	registry.registerPattern(`maximum`, "schema-value-too-large")
	registry.registerPattern(`minLength`, "schema-string-too-short")
	registry.registerPattern(`maxLength`, "schema-string-too-long")
	registry.registerPattern(`minItems`, "schema-array-too-short")
	registry.registerPattern(`maxItems`, "schema-array-too-long")
	registry.registerPattern(`uniqueItems`, "schema-duplicate-items")

	return registry
}

func (r *ErrorIDRegistry) registerPattern(pattern, id string) {
	regex := regexp.MustCompile(pattern)
	r.patterns[regex] = id
}

// AssignID assigns an appropriate error ID based on the message and location
func (r *ErrorIDRegistry) AssignID(message, location string) string {
	// Special case for kind errors
	if strings.Contains(message, "value must be") && strings.Contains(location, "kind") {
		return "schema-invalid-kind"
	}

	// Check registered patterns
	for regex, id := range r.patterns {
		if regex.MatchString(message) {
			return id
		}
	}

	return "schema-validation-error"
}

// ErrorProcessor handles the parsing and processing of validation errors
type ErrorProcessor struct {
	navigator *ManifestNavigator
	registry  *ErrorIDRegistry
}

// NewErrorProcessor creates a new error processor
func NewErrorProcessor(navigator *ManifestNavigator) *ErrorProcessor {
	return &ErrorProcessor{
		navigator: navigator,
		registry:  NewErrorIDRegistry(),
	}
}

// ProcessValidationError processes a JSON schema validation error into structured errors
func (p *ErrorProcessor) ProcessValidationError(cause *jsonschema.ValidationError, location string) ValidationError {
	ctx := ParseContext(location)
	module, declaration, member, property := p.navigator.BuildContextualNames(ctx)

	error := ValidationError{
		ID:          p.registry.AssignID(cause.Message, location),
		Module:      module,
		Declaration: declaration,
		Member:      member,
		Property:    property,
		Message:     cause.Message,
		Location:    location,
		Index:       -1,
	}

	// Set index for array items
	switch ctx.Type {
	case "member":
		if ctx.MemberIndex >= 0 {
			error.Index = ctx.MemberIndex
		}
	case "property":
		if ctx.PropertyIndex >= 0 {
			error.Index = ctx.PropertyIndex
		}
	}

	return error
}
