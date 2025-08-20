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
package platform

import (
	"time"
)

// TimeProvider provides an abstraction over time operations.
// This interface enables:
// - Controllable time in tests (no real delays)
// - Deterministic test execution
// - Fast test feedback cycles
type TimeProvider interface {
	// Sleep pauses execution for the given duration
	Sleep(d time.Duration)

	// Now returns the current time
	Now() time.Time

	// After returns a channel that delivers the current time after the duration
	After(d time.Duration) <-chan time.Time
}

// RealTimeProvider implements TimeProvider using the standard time package.
// This is the production implementation.
type RealTimeProvider struct{}

// NewRealTimeProvider creates a new time provider that uses the standard time package.
func NewRealTimeProvider() *RealTimeProvider {
	return &RealTimeProvider{}
}

func (t *RealTimeProvider) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (t *RealTimeProvider) Now() time.Time {
	return time.Now()
}

func (t *RealTimeProvider) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
