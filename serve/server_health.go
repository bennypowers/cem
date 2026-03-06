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

package serve

import (
	"bennypowers.dev/cem/health"
)

// HealthResult returns the cached health analysis result, computing it if needed.
func (s *Server) HealthResult() (*health.HealthResult, error) {
	// Check cache under read lock
	s.mu.RLock()
	cached := s.healthCache
	manifest := s.manifest
	s.mu.RUnlock()

	if cached != nil {
		return cached, nil
	}

	if len(manifest) == 0 {
		return &health.HealthResult{
			Modules:         []health.ModuleReport{},
			Recommendations: []string{},
		}, nil
	}

	// Compute health analysis without holding the lock
	result, err := health.AnalyzeBytes(manifest, health.Options{})
	if err != nil {
		return nil, err
	}

	// Cache the result
	s.mu.Lock()
	s.healthCache = result
	s.mu.Unlock()

	return result, nil
}

// invalidateHealthCache clears the cached health result.
// Must be called with s.mu write lock held.
func (s *Server) invalidateHealthCache() {
	s.healthCache = nil
}
