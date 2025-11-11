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

package serve_test

import (
	"testing"
)

// BenchmarkManifestRegeneration benchmarks manifest regeneration performance
func BenchmarkManifestRegeneration(b *testing.B) {
	// TODO: Implement once manifest generation is in place
	// This is a stub to establish the benchmark pattern
	b.ReportAllocs()

	for b.Loop() {
		// Simulate manifest regeneration
		_ = make([]byte, 1024)
	}
}

// BenchmarkTransformCache benchmarks transform cache operations
func BenchmarkTransformCache(b *testing.B) {
	// TODO: Implement once transform cache is in place
	// This is a stub to establish the benchmark pattern
	b.ReportAllocs()

	cache := make(map[string][]byte)

	b.Run("Set", func(b *testing.B) {
		for b.Loop() {
			cache["file.ts"] = []byte("content")
		}
	})

	b.Run("Get", func(b *testing.B) {
		cache["file.ts"] = []byte("content")
		b.ResetTimer()
		for b.Loop() {
			_ = cache["file.ts"]
		}
	})
}

// BenchmarkWebSocketBroadcast benchmarks WebSocket broadcast performance
func BenchmarkWebSocketBroadcast(b *testing.B) {
	// TODO: Implement once WebSocket broadcasting is in place
	// This is a stub to establish the benchmark pattern
	b.ReportAllocs()

	clients := make([]chan []byte, 10)
	for i := range clients {
		clients[i] = make(chan []byte, 1)
	}

	message := []byte(`{"type":"reload"}`)

	for b.Loop() {
		for _, client := range clients {
			select {
			case client <- message:
			default:
			}
		}
	}
}
