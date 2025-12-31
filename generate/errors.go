/*
Copyright © 2025 Benny Powers

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
	"errors"
	"fmt"
)

// Error handling patterns and utilities for consistent error management

// Common error types for the generate package
var (
	// ErrModuleNotFound is returned when a requested module cannot be found
	ErrModuleNotFound = errors.New("module not found")

	// ErrInvalidManifest is returned when manifest data is invalid
	ErrInvalidManifest = errors.New("invalid manifest")

	// ErrProcessingCancelled is returned when processing is cancelled
	ErrProcessingCancelled = errors.New("processing cancelled")
)

// Standard error wrapping patterns - these follow Go conventions:
// - Error messages start with lowercase (unless proper noun)
// - Use %w verb for error wrapping
// - Be concise but informative about the context

// WrapPreprocessError wraps preprocessing errors with context
func WrapPreprocessError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("preprocess failed: %w", err)
}

// WrapProcessError wraps processing errors with context
func WrapProcessError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("process failed: %w", err)
}

// WrapPostprocessError wraps postprocessing errors with context
func WrapPostprocessError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("postprocess failed: %w", err)
}

// WrapIncrementalError wraps incremental processing errors with context
func WrapIncrementalError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("incremental %s failed: %w", operation, err)
}

// WrapFileError wraps file operation errors with context
func WrapFileError(operation, filepath string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s file %q: %w", operation, filepath, err)
}

// WrapModuleError wraps module-specific errors with context
func WrapModuleError(modulePath string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("module %q: %w", modulePath, err)
}

// WrapComponentError wraps component-specific errors (slots, parts, etc.)
func WrapComponentError(componentType, componentName string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s %q: %w", componentType, componentName, err)
}

// NewError creates a new error with a standardized message format
func NewError(message string) error {
	return errors.New(message)
}

// NewErrorf creates a new formatted error with standardized conventions
func NewErrorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}
