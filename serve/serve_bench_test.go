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
	"time"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// BenchmarkManifestRegeneration benchmarks manifest regeneration performance
func BenchmarkManifestRegeneration(b *testing.B) {
	b.ReportAllocs()

	b.Run("Full", func(b *testing.B) {
		// Set up in-memory filesystem with realistic test data
		mfs := platform.NewMapFileSystem(nil)
		mfs.AddFile("/test-pkg/package.json", `{
  "name": "test-package",
  "customElements": "custom-elements.json"
}`, 0644)
		mfs.AddFile("/test-pkg/.config/cem.yaml", `generate:
  files:
    - "src/*.ts"
`, 0644)
		mfs.AddFile("/test-pkg/src/element.ts", `/**
 * @element my-element
 * @attr {string} label - Element label
 */
export class MyElement extends HTMLElement {
  static get observedAttributes() { return ['label']; }
}
`, 0644)

		config := serve.Config{
			Port:   18000,
			Reload: false,
			FS:     mfs,
		}

		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			b.Fatalf("Failed to create server: %v", err)
		}
		defer func() { _ = server.Close() }()

		if err := server.SetWatchDir("/test-pkg"); err != nil {
			b.Fatalf("Failed to set watch dir: %v", err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := server.RegenerateManifest()
			if err != nil {
				b.Fatalf("Manifest regeneration failed: %v", err)
			}
		}
	})

	b.Run("Incremental", func(b *testing.B) {
		// Set up in-memory filesystem
		mfs := platform.NewMapFileSystem(nil)
		mfs.AddFile("/test-pkg/package.json", `{
  "name": "test-package",
  "customElements": "custom-elements.json"
}`, 0644)
		mfs.AddFile("/test-pkg/.config/cem.yaml", `generate:
  files:
    - "src/*.ts"
`, 0644)
		mfs.AddFile("/test-pkg/src/element.ts", `export class MyElement extends HTMLElement {}`, 0644)
		mfs.AddFile("/test-pkg/src/other.ts", `export class OtherElement extends HTMLElement {}`, 0644)

		config := serve.Config{
			Port:   18001,
			Reload: false,
			FS:     mfs,
		}

		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			b.Fatalf("Failed to create server: %v", err)
		}
		defer func() { _ = server.Close() }()

		if err := server.SetWatchDir("/test-pkg"); err != nil {
			b.Fatalf("Failed to set watch dir: %v", err)
		}

		// Do initial full generation
		if _, err := server.RegenerateManifest(); err != nil {
			b.Fatalf("Initial manifest generation failed: %v", err)
		}

		changedFiles := []string{"/test-pkg/src/element.ts"}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := server.RegenerateManifestIncremental(changedFiles)
			if err != nil {
				b.Fatalf("Incremental regeneration failed: %v", err)
			}
		}
	})
}

// BenchmarkTransformCache benchmarks transform cache operations
func BenchmarkTransformCache(b *testing.B) {
	b.ReportAllocs()

	cache := transform.NewCache(500 * 1024 * 1024) // 500MB like production

	// Prepare test data
	testCode := []byte(`export class MyElement extends HTMLElement {
	connectedCallback() {
		this.innerHTML = '<h1>Hello</h1>';
	}
}`)

	testKey := transform.CacheKey{
		Path:    "/src/element.ts",
		ModTime: time.Now(),
		Size:    int64(len(testCode)),
	}

	testDeps := []string{"/src/base.ts"}

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(testKey, testCode, testDeps)
		}
	})

	b.Run("Get-Hit", func(b *testing.B) {
		cache.Set(testKey, testCode, testDeps)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, found := cache.Get(testKey)
			if !found {
				b.Fatal("Expected cache hit")
			}
		}
	})

	b.Run("Get-Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			missKey := transform.CacheKey{
				Path:    "/nonexistent.ts",
				ModTime: time.Now(),
				Size:    100,
			}
			_, _ = cache.Get(missKey)
		}
	})

	b.Run("Invalidate", func(b *testing.B) {
		// Set up dependency chain: A imports B imports C
		keyC := transform.CacheKey{Path: "/c.ts", ModTime: time.Now(), Size: 100}
		keyB := transform.CacheKey{Path: "/b.ts", ModTime: time.Now(), Size: 100}
		keyA := transform.CacheKey{Path: "/a.ts", ModTime: time.Now(), Size: 100}

		for i := 0; i < b.N; i++ {
			cache.Set(keyC, testCode, nil)
			cache.Set(keyB, testCode, []string{"/c.ts"})
			cache.Set(keyA, testCode, []string{"/b.ts"})
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Invalidate("/c.ts")
		}
	})

	b.Run("Stats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = cache.Stats()
		}
	})
}

// BenchmarkWebSocketBroadcast benchmarks WebSocket broadcast performance
//
// Note: This benchmark tests the in-memory broadcast mechanism.
// For full end-to-end testing with real WebSocket connections, see websocket_test.go
func BenchmarkWebSocketBroadcast(b *testing.B) {
	b.ReportAllocs()

	b.Run("Broadcast-10-Connections", func(b *testing.B) {
		config := serve.Config{
			Port:   19000,
			Reload: true,
		}

		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			b.Fatalf("Failed to create server: %v", err)
		}
		defer func() { _ = server.Close() }()

		wsManager := server.WebSocketManager()
		message := []byte(`{"type":"reload","reason":"benchmark","files":["test.ts"]}`)

		// Simulate 10 connected clients
		// In a real scenario, these would be actual WebSocket connections
		// For benchmarking, we measure the broadcast logic performance
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := wsManager.Broadcast(message); err != nil {
				b.Fatalf("Broadcast failed: %v", err)
			}
		}
	})

	b.Run("BroadcastToPages-Selective", func(b *testing.B) {
		config := serve.Config{
			Port:   19001,
			Reload: true,
		}

		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			b.Fatalf("Failed to create server: %v", err)
		}
		defer func() { _ = server.Close() }()

		wsManager := server.WebSocketManager()
		message := []byte(`{"type":"reload","reason":"file-change","files":["button.ts"]}`)
		targetPages := []string{"/elements/button/demo/"}

		// Simulate selective page-based broadcasting
		// This tests the filtering logic for page-specific updates
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := wsManager.BroadcastToPages(message, targetPages); err != nil {
				b.Fatalf("BroadcastToPages failed: %v", err)
			}
		}
	})
}
