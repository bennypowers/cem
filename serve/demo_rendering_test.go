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
	"testing"
)

// TestDemoRenderingMode_Default verifies light mode is default when not configured
func TestDemoRenderingMode_Default(t *testing.T) {
	// Create server with no Demos config
	server, err := NewServerWithConfig(Config{
		Port: 8123,
		Demos: DemosConfig{
			Rendering: "", // Empty should default to light
		},
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if server.IsRunning() {
			_ = server.Close()
		}
	}()

	mode := server.DemoRenderingMode()
	if mode != "light" {
		t.Errorf("Expected default rendering mode 'light', got '%s'", mode)
	}
}

// TestDemoRenderingMode_Light verifies explicit light mode config
func TestDemoRenderingMode_Light(t *testing.T) {
	server, err := NewServerWithConfig(Config{
		Port: 8124,
		Demos: DemosConfig{
			Rendering: "light",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if server.IsRunning() {
			_ = server.Close()
		}
	}()

	mode := server.DemoRenderingMode()
	if mode != "light" {
		t.Errorf("Expected rendering mode 'light', got '%s'", mode)
	}
}

// TestDemoRenderingMode_Shadow verifies shadow mode config
func TestDemoRenderingMode_Shadow(t *testing.T) {
	server, err := NewServerWithConfig(Config{
		Port: 8125,
		Demos: DemosConfig{
			Rendering: "shadow",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if server.IsRunning() {
			_ = server.Close()
		}
	}()

	mode := server.DemoRenderingMode()
	if mode != "shadow" {
		t.Errorf("Expected rendering mode 'shadow', got '%s'", mode)
	}
}

// TestDemoRenderingMode_ThreadSafety verifies method is thread-safe
func TestDemoRenderingMode_ThreadSafety(t *testing.T) {
	server, err := NewServerWithConfig(Config{
		Port: 8126,
		Demos: DemosConfig{
			Rendering: "shadow",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if server.IsRunning() {
			_ = server.Close()
		}
	}()

	// Call from multiple goroutines concurrently
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			mode := server.DemoRenderingMode()
			if mode != "shadow" {
				t.Errorf("Expected 'shadow', got '%s'", mode)
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
}
