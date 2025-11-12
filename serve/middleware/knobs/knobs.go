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

package knobs

import "errors"

// GenerateKnobsHTML generates HTML for knobs UI from manifest and demo
//
// NOTE: Phase 5a (Basic Knobs) is complete! The working implementation is in
// serve/middleware/routes/knobs.go (GenerateKnobs + RenderKnobsHTML).
//
// This standalone function is planned for Phase 5b as a refactoring to extract
// knob generation logic from the routes package into a reusable API.
// See serve/PLANS/51-KNOBS-ADVANCED.md for Phase 5b specification.
func GenerateKnobsHTML(manifestJSON []byte, demoHTML []byte) (string, error) {
	return "", errors.New("not implemented - planned for Phase 5b refactoring")
}
