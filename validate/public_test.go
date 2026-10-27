/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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
package validate_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/validate"
)

// The public alias must be nameable from an external package.
var _ validate.FileSystem = validate.NewOSFileSystem()

// TestPublicValidateAPI exercises the documented external usage of the
// package. It deliberately imports only the public validate package and the
// standard library, so it fails to compile if an internal type leaks into the
// public example.
func TestPublicValidateAPI(t *testing.T) {
	fsys := validate.NewOSFileSystem()
	manifestPath := filepath.Join(t.TempDir(), "custom-elements.json")
	manifest := []byte(`{"schemaVersion":"2.1.0","modules":[]}`)
	if err := fsys.WriteFile(manifestPath, manifest, 0o644); err != nil {
		t.Fatalf("writing manifest: %v", err)
	}

	result, err := validate.Validate(fsys, manifestPath, validate.ValidationOptions{
		IncludeWarnings: true,
	})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate reported invalid manifest: %+v", result.Errors)
	}

	result, err = validate.ValidateFile(manifestPath, validate.ValidationOptions{})
	if err != nil {
		t.Fatalf("ValidateFile: %v", err)
	}
	if !result.IsValid {
		t.Errorf("ValidateFile reported invalid manifest: %+v", result.Errors)
	}
}
