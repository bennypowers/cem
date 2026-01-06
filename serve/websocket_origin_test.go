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
	"net/http"
	"testing"
)

// TestIsLocalOrigin_AllowedOrigins verifies that localhost origins are allowed
func TestIsLocalOrigin_AllowedOrigins(t *testing.T) {
	tests := []struct {
		name   string
		origin string
		host   string // Request Host header
	}{
		{"localhost with port", "http://localhost:8000", "localhost:8000"},
		{"localhost without port", "http://localhost", "localhost"},
		{"localhost https", "https://localhost:8000", "localhost:8000"},
		{"127.0.0.1", "http://127.0.0.1:8000", "127.0.0.1:8000"},
		{"127.0.0.1 without port", "http://127.0.0.1", "127.0.0.1"},
		{"127.x.x.x range", "http://127.0.0.2:8000", "127.0.0.2:8000"},
		{"127.255.255.255", "http://127.255.255.255:8000", "127.255.255.255:8000"},
		{"IPv6 localhost brackets", "http://[::1]:8000", "[::1]:8000"},
		{"localhost subdomain", "http://app.localhost:8000", "app.localhost:8000"},
		{"localhost subdomain no port", "http://app.localhost", "app.localhost"},
		{"no origin header", "", "localhost:8000"}, // Empty origin should be allowed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Header: make(http.Header),
				Host:   tt.host,
			}
			if tt.origin != "" {
				req.Header.Set("Origin", tt.origin)
			}

			if !isLocalOrigin(req) {
				t.Errorf("Expected origin %q to be allowed", tt.origin)
			}
		})
	}
}

// TestIsLocalOrigin_SameOriginAllowed tests that same-origin requests are allowed
// This supports Cloudflare tunnels, reverse proxies, and containerized deployments
func TestIsLocalOrigin_SameOriginAllowed(t *testing.T) {
	tests := []struct {
		name   string
		origin string
		host   string
	}{
		{"cloudflare tunnel", "https://my-app.trycloudflare.com", "my-app.trycloudflare.com"},
		{"cloudflare tunnel with port", "https://my-app.trycloudflare.com:443", "my-app.trycloudflare.com:443"},
		{"reverse proxy", "https://dev.example.com", "dev.example.com"},
		{"custom domain", "https://my-project.dev", "my-project.dev"},
		{"ngrok tunnel", "https://abc123.ngrok.io", "abc123.ngrok.io"},
		{"host with port", "http://example.com:8080", "example.com:8080"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Header: make(http.Header),
				Host:   tt.host,
			}
			req.Header.Set("Origin", tt.origin)

			if !isLocalOrigin(req) {
				t.Errorf("Expected same-origin request to be allowed: origin=%q host=%q", tt.origin, tt.host)
			}
		})
	}
}

// TestIsLocalOrigin_RejectedOrigins verifies that non-localhost origins are rejected
func TestIsLocalOrigin_RejectedOrigins(t *testing.T) {
	tests := []struct {
		name   string
		origin string
		host   string
	}{
		{"external domain", "http://example.com", "localhost:8000"},
		{"external https", "https://example.com:8000", "localhost:8000"},
		{"fake localhost domain", "http://localhost.evil.com", "localhost:8000"},
		{"192.168.x.x", "http://192.168.1.1:8000", "localhost:8000"},
		{"10.x.x.x", "http://10.0.0.1:8000", "localhost:8000"},
		{"public IP", "http://8.8.8.8:8000", "localhost:8000"},
		{"malformed localhost", "http://local host:8000", "localhost:8000"}, // Space in hostname
		{"128.x.x.x (not 127)", "http://128.0.0.1:8000", "localhost:8000"},
		{"fake 127 domain", "http://127.0.0.1.evil.com", "localhost:8000"},
		{"127 with only 3 parts", "http://127.0.1:8000", "localhost:8000"}, // Not a valid 127.x.x.x
		{"starts with 127 but not IP", "http://127evil.com:8000", "localhost:8000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Header: make(http.Header),
				Host:   tt.host,
			}
			req.Header.Set("Origin", tt.origin)

			if isLocalOrigin(req) {
				t.Errorf("Expected origin %q to be rejected", tt.origin)
			}
		})
	}
}

// TestIsLocalOrigin_CrossOriginAttacks tests that cross-origin attacks are blocked
// even when the server is exposed via Cloudflare tunnel or reverse proxy
func TestIsLocalOrigin_CrossOriginAttacks(t *testing.T) {
	tests := []struct {
		name   string
		origin string // Attacker's origin
		host   string // Server's host
	}{
		{"evil.com attacking tunnel", "https://evil.com", "my-app.trycloudflare.com"},
		{"different tunnel attacking", "https://attacker.ngrok.io", "victim.ngrok.io"},
		{"subdomain attack", "https://evil.example.com", "app.example.com"},
		{"subdomain trying to attack parent", "https://sub.app.example.com", "app.example.com"},
		{"lookalike domain", "https://app.examp1e.com", "app.example.com"}, // Note the 1 vs l
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Header: make(http.Header),
				Host:   tt.host,
			}
			req.Header.Set("Origin", tt.origin)

			if isLocalOrigin(req) {
				t.Errorf("Expected cross-origin attack to be blocked: origin=%q host=%q", tt.origin, tt.host)
			}
		})
	}
}

// TestIsLocalOrigin_InvalidOrigins verifies that malformed origins are rejected
func TestIsLocalOrigin_InvalidOrigins(t *testing.T) {
	tests := []struct {
		name   string
		origin string
	}{
		{"invalid URL", "not-a-url"},
		{"missing scheme", "localhost:8000"},
		{"malformed", "http://[invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Header: make(http.Header),
				Host:   "localhost:8000",
			}
			req.Header.Set("Origin", tt.origin)

			if isLocalOrigin(req) {
				t.Errorf("Expected malformed origin %q to be rejected", tt.origin)
			}
		})
	}
}

// TestMaxWebSocketReadSize verifies the read size limit is set appropriately
func TestMaxWebSocketReadSize(t *testing.T) {
	// Verify the constant is defined and reasonable
	if maxWebSocketReadSize <= 0 {
		t.Error("maxWebSocketReadSize must be positive")
	}

	// 64KB should be sufficient for our use cases (reload messages, error messages)
	if maxWebSocketReadSize < 64*1024 {
		t.Error("maxWebSocketReadSize should be at least 64KB")
	}

	// But not too large (prevents DoS)
	if maxWebSocketReadSize > 1024*1024 {
		t.Error("maxWebSocketReadSize should not exceed 1MB")
	}
}
